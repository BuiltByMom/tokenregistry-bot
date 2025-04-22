package main

import (
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
	// Create channels to collect results
	resultChan := make(chan struct {
		chain   string
		err     error
		changes bool
	}, len(cfg.Chains))
	
	// Generate token lists for all chains in parallel
	for _, chain := range cfg.Chains {
		wg.Add(1)
		go func(chain config.ChainConfig) {
			defer wg.Done()
			
			result := generator.GenerateForChain(chain)
			resultChan <- struct {
				chain   string
				err     error
				changes bool
			}{
				chain:   chain.Name,
				err:     result.Error,
				changes: result.HasChanges,
			}
		}(chain)
	}

	// Start a goroutine to close channels after all work is done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Process results as they come in
	var hasErrors bool
	var hasChanges bool
	
	for result := range resultChan {
		if result.err != nil {
			hasErrors = true
			log.Printf("Error generating token list for %s: %v", result.chain, result.err)
		} else if result.changes {
			hasChanges = true
			log.Printf("Success: Generated token list for %s (changes detected)", result.chain)
		} else {
			log.Printf("Success: Generated token list for %s (no changes)", result.chain)
		}
	}

	if hasErrors {
		os.Exit(1)
	}

	// Just log the final status
	if hasChanges {
		log.Println("Changes detected in token lists")
	} else {
		log.Println("No changes detected in token lists")
	}

	// Exit with status code 0 if there are changes, 2 if no changes
	// This allows the workflow to distinguish between changes and no changes
	if !hasChanges {
		os.Exit(2)
	}
}