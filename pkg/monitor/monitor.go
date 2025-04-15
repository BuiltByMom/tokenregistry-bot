package monitor

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/builtbymom/TokenRegistry-bot/pkg/config"
	"github.com/builtbymom/TokenRegistry-bot/pkg/contracts/tokenedits"
	"github.com/builtbymom/TokenRegistry-bot/pkg/contracts/tokenregistry"
	"github.com/builtbymom/TokenRegistry-bot/pkg/telegram"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EventEmojis maps event names to their corresponding emojis
var EventEmojis = map[string]string{
	"TokenAdded":        "➕",
	"TokenApproved":     "✅",
	"TokenRejected":     "❌",
	"TokentrollerUpdated": "🔄",
	"EditProposed":      "📝",
	"EditAccepted":      "✅",
	"EditRejected":      "❌",
}

// ERC20ABI contains the minimum ABI needed to fetch token name and symbol
const ERC20ABI = `[
	{
		"constant": true,
		"inputs": [],
		"name": "name",
		"outputs": [{"name": "", "type": "string"}],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "symbol",
		"outputs": [{"name": "", "type": "string"}],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]`

// TokenInfo holds basic token information
type TokenInfo struct {
	Symbol string
	Name   string
}

// Monitor handles event monitoring for a specific chain
type Monitor struct {
	client  *ethclient.Client
	chain   config.ChainConfig
	bot     *telegram.Bot
	config  *config.Config
}

// New creates a new Monitor instance
func New(client *ethclient.Client, chain config.ChainConfig, bot *telegram.Bot, config *config.Config) *Monitor {
	return &Monitor{
		client: client,
		chain:  chain,
		bot:    bot,
		config: config,
	}
}

// Start begins monitoring events for the chain
func (m *Monitor) Start(ctx context.Context) error {
	// Create contract instances
	registry, err := tokenregistry.NewTokenRegistry(m.chain.RegistryAddress, m.client)
	if err != nil {
		return fmt.Errorf("failed to create TokenRegistry instance: %v", err)
	}

	edits, err := tokenedits.NewTokenEdits(m.chain.EditsAddress, m.client)
	if err != nil {
		return fmt.Errorf("failed to create TokenEdits instance: %v", err)
	}

	// Create a query for all events from both contracts
	query := ethereum.FilterQuery{
		Addresses: []common.Address{m.chain.RegistryAddress, m.chain.EditsAddress},
	}

	// Subscribe to logs
	logs := make(chan types.Log)
	sub, err := m.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %v", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			log.Printf("Subscription error on %s: %v", m.chain.Name, err)
		case evtLog := <-logs:
			if err := m.handleEvent(ctx, registry, edits, evtLog); err != nil {
				log.Printf("Error handling event: %v", err)
			}
		}
	}
}

func (m *Monitor) handleEvent(ctx context.Context, registry *tokenregistry.TokenRegistry, edits *tokenedits.TokenEdits, evtLog types.Log) error {
	// Get transaction details
	tx, _, err := m.client.TransactionByHash(ctx, evtLog.TxHash)
	if err != nil {
		return fmt.Errorf("error getting transaction: %v", err)
	}

	// Get sender address from the transaction
	chainID, err := m.client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("error getting chain ID: %v", err)
	}

	signer := types.LatestSignerForChainID(chainID)
	sender, err := types.Sender(signer, tx)
	if err != nil {
		return fmt.Errorf("error getting transaction sender: %v", err)
	}

	// Try to parse the event
	var eventName string
	var tokenAddr common.Address
	var editID *big.Int
	var reason string

	switch evtLog.Address {
	case m.chain.RegistryAddress:
		if tokenAdded, err := registry.ParseTokenAdded(evtLog); err == nil {
			eventName = "TokenAdded"
			tokenAddr = tokenAdded.ContractAddress
		} else if tokenApproved, err := registry.ParseTokenApproved(evtLog); err == nil {
			eventName = "TokenApproved"
			tokenAddr = tokenApproved.ContractAddress
		} else if tokenRejected, err := registry.ParseTokenRejected(evtLog); err == nil {
			eventName = "TokenRejected"
			tokenAddr = tokenRejected.ContractAddress
			reason = tokenRejected.Reason
		} else if tokentrollerUpdated, err := registry.ParseTokentrollerUpdated(evtLog); err == nil {
			eventName = "TokentrollerUpdated"
			tokenAddr = tokentrollerUpdated.NewTokentroller
		}
	case m.chain.EditsAddress:
		if editProposed, err := edits.ParseEditProposed(evtLog); err == nil {
			eventName = "EditProposed"
			tokenAddr = editProposed.ContractAddress
			editID = editProposed.EditId
		} else if editAccepted, err := edits.ParseEditAccepted(evtLog); err == nil {
			eventName = "EditAccepted"
			tokenAddr = editAccepted.ContractAddress
			editID = editAccepted.EditId
		} else if editRejected, err := edits.ParseEditRejected(evtLog); err == nil {
			eventName = "EditRejected"
			tokenAddr = editRejected.ContractAddress
			editID = editRejected.EditId
			reason = editRejected.Reason
		}
	}

	if eventName == "" {
		return fmt.Errorf("unknown event from address %s", evtLog.Address.Hex())
	}

	// Get token info
	tokenInfo := m.getTokenInfo(ctx, tokenAddr)

	// Format the message
	emoji := EventEmojis[eventName]
	
	// Format URL based on event type
	var url string
	if editID != nil {
		url = fmt.Sprintf("%s/%s?token=%s-%s", m.config.UIBaseURL, strings.ToLower(m.chain.Name), tokenAddr.Hex(), editID.String())
	} else {
		url = fmt.Sprintf("%s/%s?token=%s", m.config.UIBaseURL, strings.ToLower(m.chain.Name), tokenAddr.Hex())
	}

	message := fmt.Sprintf("%s New %s event detected\n\nChain: %s (ID: %d)\nToken: %s (%s)\nContract: [%s](%s/token/%s)\nSubmitter: [%s](%s/address/%s)\nTransaction: [%s](%s/tx/%s)",
		emoji,
		eventName,
		m.chain.Name,
		chainID,
		tokenInfo.Name,
		tokenInfo.Symbol,
		tokenAddr.Hex(),
		m.chain.ExplorerURL,
		tokenAddr.Hex(),
		sender.Hex(),
		m.chain.ExplorerURL,
		sender.Hex(),
		evtLog.TxHash.Hex()[:8] + "..." + evtLog.TxHash.Hex()[len(evtLog.TxHash.Hex())-6:],
		m.chain.ExplorerURL,
		evtLog.TxHash.Hex(),
	)

	if reason != "" {
		message += fmt.Sprintf("\nReason: %s", reason)
	}

	message += fmt.Sprintf("\n[View in UI](%s)", url)

	return m.bot.SendMessage(message)
}

func (m *Monitor) getTokenInfo(ctx context.Context, tokenAddr common.Address) TokenInfo {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		log.Printf("Error parsing ABI: %v", err)
		return TokenInfo{Symbol: "UNKNOWN", Name: "Unknown Token"}
	}

	// Pack the name function call
	nameData, err := parsedABI.Pack("name")
	if err != nil {
		log.Printf("Error packing name call: %v", err)
		return TokenInfo{Symbol: "UNKNOWN", Name: "Unknown Token"}
	}

	// Call the name function
	nameResult, err := m.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddr,
		Data: nameData,
	}, nil)
	if err != nil {
		log.Printf("Error calling name: %v", err)
		return TokenInfo{Symbol: "UNKNOWN", Name: "Unknown Token"}
	}

	// Unpack the name result
	var name string
	err = parsedABI.UnpackIntoInterface(&name, "name", nameResult)
	if err != nil {
		log.Printf("Error unpacking name: %v", err)
		name = "Unknown Token"
	}

	// Pack the symbol function call
	symbolData, err := parsedABI.Pack("symbol")
	if err != nil {
		log.Printf("Error packing symbol call: %v", err)
		return TokenInfo{Symbol: "UNKNOWN", Name: name}
	}

	// Call the symbol function
	symbolResult, err := m.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddr,
		Data: symbolData,
	}, nil)
	if err != nil {
		log.Printf("Error calling symbol: %v", err)
		return TokenInfo{Symbol: "UNKNOWN", Name: name}
	}

	// Unpack the symbol result
	var symbol string
	err = parsedABI.UnpackIntoInterface(&symbol, "symbol", symbolResult)
	if err != nil {
		log.Printf("Error unpacking symbol: %v", err)
		symbol = "UNKNOWN"
	}

	return TokenInfo{
		Symbol: symbol,
		Name:   name,
	}
} 