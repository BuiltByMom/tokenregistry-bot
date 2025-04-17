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
	ChainID        int64
}

// Config holds the configuration for the bot
type Config struct {
	TelegramBotToken string
	TelegramChannel  string
	UIBaseURL        string
	Chains           []ChainConfig
}

// LoadConfig loads configuration from .env file if present, otherwise from environment variables
func LoadConfig() (*Config, error) {
	// Try to load from .env file, but don't fail if it doesn't exist
	_ = godotenv.Load()

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
				ChainID:        10,
			},
			{
				Name:           "Base",
				RPCURL:         os.Getenv("BASE_RPC_URL"),
				RegistryAddress: common.HexToAddress(os.Getenv("BASE_REGISTRY_ADDRESS")),
				EditsAddress:   common.HexToAddress(os.Getenv("BASE_EDITS_ADDRESS")),
				ExplorerURL:    "https://basescan.org",
				ChainID:        8453,
			},
		},
	}, nil
} 