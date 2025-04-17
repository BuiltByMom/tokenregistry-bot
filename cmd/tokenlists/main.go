package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/builtbymom/TokenRegistry-bot/pkg/config"
	"github.com/builtbymom/TokenRegistry-bot/pkg/tokenlist"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize generator
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	outputDir := filepath.Join(dir, "data", "tokenlists")
	generator := tokenlist.New(outputDir)

	// Generate token lists for all chains
	for _, chain := range cfg.Chains {
		if err := generator.GenerateForChain(chain); err != nil {
			log.Printf("Error generating token list for %s: %v", chain.Name, err)
			continue
		}
		log.Printf("Generated token list for %s", chain.Name)
	}
}