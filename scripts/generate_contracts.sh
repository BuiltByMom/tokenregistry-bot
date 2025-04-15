#!/bin/bash

# Create contracts directory if it doesn't exist
mkdir -p pkg/contracts/tokenedits pkg/contracts/tokenregistry

# Get the path to abigen
ABIGEN_PATH="$HOME/go/bin/abigen"

# Move to the TokenRegistry directory and compile contracts
cd ../TokenRegistry
forge build

# Move back to the bot directory
cd ../TokenRegistry-bot

# Create temporary directory for ABIs
mkdir -p tmp

# Extract ABI arrays from Forge output
jq .abi ../TokenRegistry/out/TokenRegistry.sol/TokenRegistry.json > tmp/TokenRegistry.abi.json
jq .abi ../TokenRegistry/out/TokenEdits.sol/TokenEdits.json > tmp/TokenEdits.abi.json

# Generate contracts for TokenRegistry with a specific package name
$ABIGEN_PATH --abi tmp/TokenRegistry.abi.json \
       --pkg tokenregistry \
       --out pkg/contracts/tokenregistry/tokenregistry.go \
       --type TokenRegistry

# Generate contracts for TokenEdits with a different package name
$ABIGEN_PATH --abi tmp/TokenEdits.abi.json \
       --pkg tokenedits \
       --out pkg/contracts/tokenedits/tokenedits.go \
       --type TokenEdits

# Clean up temporary files
rm -rf tmp 