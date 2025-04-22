package tokenlist

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/builtbymom/TokenRegistry-bot/pkg/config"
	"github.com/builtbymom/TokenRegistry-bot/pkg/contracts/tokenregistry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GeneratorResult contains the result of token list generation
type GeneratorResult struct {
	HasChanges bool
	Error      error
}

// hasTokenChanges compares two token lists and returns true if there are any differences
func hasTokenChanges(existing, new TokenList) bool {
	if len(existing.Tokens) != len(new.Tokens) {
		return true
	}

	// Create maps for faster lookup
	existingTokens := make(map[string]Token)
	for _, token := range existing.Tokens {
		existingTokens[token.Address] = token
	}

	// Compare each new token with existing
	for _, newToken := range new.Tokens {
		existingToken, exists := existingTokens[newToken.Address]
		if !exists {
			return true
		}
		
		// Compare all fields except timestamp
		if newToken.Name != existingToken.Name ||
			newToken.Symbol != existingToken.Symbol ||
			newToken.LogoURI != existingToken.LogoURI ||
			newToken.ChainId != existingToken.ChainId ||
			newToken.Decimals != existingToken.Decimals {
			return true
		}
	}

	return false
}

func New(outputDir string) *Generator {
	return &Generator{
		outputDir: outputDir,
	}
}

// GenerateForChain generates the token list for a specific chain and returns whether changes were detected
func (g *Generator) GenerateForChain(chain config.ChainConfig) GeneratorResult {
	result := GeneratorResult{HasChanges: false}

	// Connect to chain's RPC
	client, err := ethclient.Dial(chain.RPCURL)
	if err != nil {
		result.Error = fmt.Errorf("failed to connect to RPC for %s: %v", chain.Name, err)
		return result
	}
	defer client.Close()

	// Initialize registry contract
	registry, err := tokenregistry.NewTokenRegistry(chain.RegistryAddress, client)
	if err != nil {
		result.Error = fmt.Errorf("failed to create registry instance for %s: %v", chain.Name, err)
		return result
	}

	// Create token list
	list := TokenList{
		Name:        fmt.Sprintf("Token Registry - %s", chain.Name),
		Description: fmt.Sprintf("Official Token Registry list for %s", chain.Name),
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		Version: Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
		Keywords: []string{"token registry", chain.Name},
		Tokens:   []Token{},
	}

	// Fetch tokens with pagination
	const pageSize = 100
	offset := big.NewInt(0)
	totalTokens := big.NewInt(0)
	var listErr error

	for {
		// Get tokens for current page
		fetchResult, err := registry.ListTokens(&bind.CallOpts{}, offset, big.NewInt(pageSize), 1, false)
		if err != nil {
			listErr = err
			break
		}

		// Update total tokens if this is the first page
		if offset.Cmp(big.NewInt(0)) == 0 {
			totalTokens = fetchResult.Total
		}

		// Convert contract tokens to tokenlist format
		for _, t := range fetchResult.Tokens {
			list.Tokens = append(list.Tokens, Token{
				Address:  t.ContractAddress.Hex(),
				Name:     t.Name,
				Symbol:   t.Symbol,
				LogoURI:  t.LogoURI,
				ChainId:  chain.ChainID,
				Decimals: int(t.Decimals),
			})
		}

		// Check if we've fetched all tokens
		offset.Add(offset, big.NewInt(pageSize))
		if offset.Cmp(totalTokens) >= 0 {
			break
		}
	}

	if listErr != nil {
		result.Error = fmt.Errorf("failed to list tokens for %s: %v", chain.Name, listErr)
		return result
	}

	// Create chain-specific output directory
	chainDir := filepath.Join(g.outputDir, fmt.Sprintf("%d", chain.ChainID))
	if err := os.MkdirAll(chainDir, 0755); err != nil {
		result.Error = fmt.Errorf("failed to create chain directory: %v", err)
		return result
	}

	// Check for existing token list
	outputPath := filepath.Join(chainDir, "tokens.json")
	var existingList TokenList
	existingData, err := os.ReadFile(outputPath)
	if err == nil {
		if err := json.Unmarshal(existingData, &existingList); err != nil {
			// If we can't unmarshal, treat it as a new file
			result.HasChanges = true
		} else {
			// Compare with new list
			if !hasTokenChanges(existingList, list) {
				// No changes detected
				return result
			}
			// Changes detected, increment patch version
			list.Version = existingList.Version
			list.Version.Patch++
			result.HasChanges = true
		}
	} else if os.IsNotExist(err) {
		// No existing file means we have changes
		result.HasChanges = true
	} else {
		// Other file system errors should be reported
		result.Error = fmt.Errorf("failed to read existing token list: %v", err)
		return result
	}

	// Write token list to file only if there are changes or no existing file
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		result.Error = fmt.Errorf("failed to marshal token list: %v", err)
		return result
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		result.Error = fmt.Errorf("failed to write token list: %v", err)
		return result
	}

	return result
} 