package config

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
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

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
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
	}, nil
} 