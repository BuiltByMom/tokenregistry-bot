package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/builtbymom/TokenRegistry-bot/pkg/config"
	"github.com/builtbymom/TokenRegistry-bot/pkg/monitor"
	"github.com/builtbymom/TokenRegistry-bot/pkg/telegram"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Initialize Telegram bot
	bot, err := telegram.New(cfg.TelegramBotToken, cfg.TelegramChannel)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Get chain names for startup message
	chainNames := make([]string, len(cfg.Chains))
	for i, chain := range cfg.Chains {
		chainNames[i] = chain.Name
	}

	// Send startup message
	if err := bot.SendStartupMessage(chainNames); err != nil {
		log.Printf("Failed to send startup message: %v", err)
	}
	log.Println("Startup message sent to channel")

	// Create a channel for OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start event monitoring for each chain
	for _, chain := range cfg.Chains {
		client, err := ethclient.Dial(chain.RPCURL)
		if err != nil {
			log.Printf("Failed to connect to %s node: %v", chain.Name, err)
			continue
		}

		mon := monitor.New(client, chain, bot, cfg)
		go func() {
			if err := mon.Start(ctx); err != nil {
				log.Printf("Error in monitor for %s: %v", chain.Name, err)
			}
		}()
	}

	// Wait for interrupt signal
	<-sigChan
	log.Println("Shutting down...")
} 