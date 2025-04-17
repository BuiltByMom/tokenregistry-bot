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

func New(outputDir string) *Generator {
	return &Generator{
		outputDir: outputDir,
	}
}

func (g *Generator) GenerateForChain(chain config.ChainConfig) error {
	// Connect to chain's RPC
	client, err := ethclient.Dial(chain.RPCURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC for %s: %v", chain.Name, err)
	}
	defer client.Close()

	// Initialize registry contract
	registry, err := tokenregistry.NewTokenRegistry(chain.RegistryAddress, client)
	if err != nil {
		return fmt.Errorf("failed to create registry instance for %s: %v", chain.Name, err)
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
		Keywords: []string{"ajna", "token registry", chain.Name},
		Tokens:   []Token{},
	}

	// Fetch tokens with pagination
	const pageSize = 100
	offset := big.NewInt(0)
	totalTokens := big.NewInt(0)

	for {
		// Get tokens for current page
		result, err := registry.ListTokens(&bind.CallOpts{}, offset, big.NewInt(pageSize), 1, false)
		if err != nil {
			return fmt.Errorf("failed to list tokens for %s: %v", chain.Name, err)
		}

		// Update total tokens if this is the first page
		if offset.Cmp(big.NewInt(0)) == 0 {
			totalTokens = result.Total
		}

		// Convert contract tokens to tokenlist format
		for _, t := range result.Tokens {
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

	// Create chain-specific output directory
	chainDir := filepath.Join(g.outputDir, fmt.Sprintf("%d", chain.ChainID))
	if err := os.MkdirAll(chainDir, 0755); err != nil {
		return fmt.Errorf("failed to create chain directory: %v", err)
	}

	// Write token list to file
	outputPath := filepath.Join(chainDir, "tokens.json")
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal token list: %v", err)
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write token list: %v", err)
	}

	return nil
} 