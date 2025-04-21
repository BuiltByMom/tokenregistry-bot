package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

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

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	// Create a channel to collect errors
	errChan := make(chan error, len(cfg.Chains))
	// Create a channel to collect success messages
	successChan := make(chan string, len(cfg.Chains))
	
	// Generate token lists for all chains in parallel
	for _, chain := range cfg.Chains {
		wg.Add(1)
		go func(chain config.ChainConfig) {
			defer wg.Done()
			
			if err := generator.GenerateForChain(chain); err != nil {
				errChan <- fmt.Errorf("error generating token list for %s: %v", chain.Name, err)
				return
			}
			successChan <- fmt.Sprintf("Generated token list for %s", chain.Name)
		}(chain)
	}

	// Start a goroutine to close channels after all work is done
	go func() {
		wg.Wait()
		close(errChan)
		close(successChan)
	}()

	// Process results as they come in
	var hasErrors bool
	for i := 0; i < len(cfg.Chains); i++ {
		select {
		case err := <-errChan:
			hasErrors = true
			log.Printf("%v", err)
		case msg := <-successChan:
			log.Printf("%s", msg)
		}
	}

	if hasErrors {
		os.Exit(1)
	}
}