package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/builtbymom/TokenRegistry-bot/bindings/tokenedits"
	"github.com/builtbymom/TokenRegistry-bot/bindings/tokenregistry"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// ChainConfig holds configuration for a specific chain
type ChainConfig struct {
	Name           string
	RPCURL         string
	RegistryAddress common.Address
	EditsAddress   common.Address
	ExplorerURL    string
}

// Config holds the configuration for the bot
type Config struct {
	TelegramBotToken string
	TelegramChannel  string
	UIBaseURL        string
	Chains           []ChainConfig
}

// TokenInfo holds basic token information
type TokenInfo struct {
	Symbol string
	Name   string
}

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

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Load configuration
	config := Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChannel:  os.Getenv("TELEGRAM_CHANNEL"),
		UIBaseURL:        os.Getenv("UI_BASE_URL"),
		Chains: []ChainConfig{
			{
				Name:           "Optimism",
				RPCURL:         os.Getenv("OPTIMISM_RPC_URL"),
				RegistryAddress: common.HexToAddress(os.Getenv("OPTIMISM_REGISTRY_ADDRESS")),
				EditsAddress:   common.HexToAddress(os.Getenv("OPTIMISM_EDITS_ADDRESS")),
				ExplorerURL:    "https://optimistic.etherscan.io",
			},
			{
				Name:           "Base",
				RPCURL:         os.Getenv("BASE_RPC_URL"),
				RegistryAddress: common.HexToAddress(os.Getenv("BASE_REGISTRY_ADDRESS")),
				EditsAddress:   common.HexToAddress(os.Getenv("BASE_EDITS_ADDRESS")),
				ExplorerURL:    "https://basescan.org",
			},
		},
	}

	// Initialize Telegram bot
	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Send startup message
	startupMsg := "🤖 TokenRegistry Monitor Bot is online!\n\n📊 Monitoring on chains:\n"
	for _, chain := range config.Chains {
		startupMsg += fmt.Sprintf("- %s\n", chain.Name)
	}
	sendTelegramMessage(bot, config.TelegramChannel, startupMsg)
	log.Println("Startup message sent to channel")

	// Create a channel for OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start event monitoring for each chain
	for _, chain := range config.Chains {
		client, err := ethclient.Dial(chain.RPCURL)
		if err != nil {
			log.Printf("Failed to connect to %s node: %v", chain.Name, err)
			continue
		}
		go monitorEvents(ctx, client, chain, bot, config.TelegramChannel, config)
	}

	// Wait for interrupt signal
	<-sigChan
	log.Println("Shutting down...")
}

func monitorEvents(ctx context.Context, client *ethclient.Client, chain ChainConfig, bot *tgbotapi.BotAPI, channelID string, config Config) {
	// Create contract instances
	registry, err := tokenregistry.NewTokenRegistry(chain.RegistryAddress, client)
	if err != nil {
		log.Printf("Failed to create TokenRegistry instance: %v", err)
		return
	}

	edits, err := tokenedits.NewTokenEdits(chain.EditsAddress, client)
	if err != nil {
		log.Printf("Failed to create TokenEdits instance: %v", err)
		return
	}

	// Create a query for all events from both contracts
	query := ethereum.FilterQuery{
		Addresses: []common.Address{chain.RegistryAddress, chain.EditsAddress},
	}

	// Subscribe to logs
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Printf("Failed to subscribe to logs on %s: %v", chain.Name, err)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return
		case err := <-sub.Err():
			log.Printf("Subscription error on %s: %v", chain.Name, err)
		case evtLog := <-logs:
			// Get transaction details
			tx, _, err := client.TransactionByHash(ctx, evtLog.TxHash)
			if err != nil {
				log.Printf("Error getting transaction on %s: %v", chain.Name, err)
				continue
			}

			// Get sender address from the transaction
			chainID, err := client.ChainID(ctx)
			if err != nil {
				log.Printf("Error getting chain ID on %s: %v", chain.Name, err)
				continue
			}

			signer := types.LatestSignerForChainID(chainID)
			sender, err := types.Sender(signer, tx)
			if err != nil {
				log.Printf("Error getting transaction sender on %s: %v", chain.Name, err)
				continue
			}

			// Try to parse the event using the generated bindings
			var eventName string
			var tokenAddr common.Address
			var editID *big.Int
			var reason string

			switch evtLog.Address {
			case chain.RegistryAddress:
				// Try to parse TokenRegistry events
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
			case chain.EditsAddress:
				// Try to parse TokenEdits events
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
				log.Printf("Unknown event from address %s", evtLog.Address.Hex())
				continue
			}

			// Get token info
			tokenInfo := getTokenInfo(client, tokenAddr)

			// Format the message
			emoji := EventEmojis[eventName]
			
			// Format URL based on event type
			var url string
			if editID != nil {
				url = fmt.Sprintf("%s/%s?token=%s-%s", config.UIBaseURL, strings.ToLower(chain.Name), tokenAddr.Hex(), editID.String())
			} else {
				url = fmt.Sprintf("%s/%s?token=%s", config.UIBaseURL, strings.ToLower(chain.Name), tokenAddr.Hex())
			}

			
			
			// Build message with optional reason
			message := fmt.Sprintf("%s New %s event detected\n\nChain: %s (ID: %d)\nToken: %s (%s)\nContract: [%s](%s/token/%s)\nSubmitter: [%s](%s/address/%s)\nTransaction: [%s](%s/tx/%s)",
				emoji,
				eventName,
				chain.Name,
				chainID,
				tokenInfo.Name,
				tokenInfo.Symbol,
				tokenAddr.Hex(),
				chain.ExplorerURL,
				tokenAddr.Hex(),
				sender.Hex(),
				chain.ExplorerURL,
				sender.Hex(),
				evtLog.TxHash.Hex()[:8] + "..." + evtLog.TxHash.Hex()[len(evtLog.TxHash.Hex())-6:],
				chain.ExplorerURL,
				evtLog.TxHash.Hex(),
			)

			// Add reason if present
			if reason != "" {
				message += fmt.Sprintf("\nReason: %s", reason)
			}

			// Add URL
			message += fmt.Sprintf("\n[View in UI](%s)", url)

			// Send message with Markdown formatting
			telegramMsg := tgbotapi.NewMessageToChannel(channelID, message)
			telegramMsg.ParseMode = "Markdown"
			_, err = bot.Send(telegramMsg)
			if err != nil {
				log.Printf("Failed to send Telegram message: %v", err)
			}
		}
	}
}

// getTokenInfo retrieves token information from the contract
func getTokenInfo(client *ethclient.Client, tokenAddr common.Address) TokenInfo {
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
	nameResult, err := client.CallContract(context.Background(), ethereum.CallMsg{
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
	symbolResult, err := client.CallContract(context.Background(), ethereum.CallMsg{
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

func sendTelegramMessage(bot *tgbotapi.BotAPI, channelID string, message string) {
	msg := tgbotapi.NewMessageToChannel(channelID, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send Telegram message: %v", err)
	}
} 