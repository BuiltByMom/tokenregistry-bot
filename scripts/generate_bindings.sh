#!/bin/bash

# Create bindings directory if it doesn't exist
mkdir -p bindings

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

# Generate bindings for TokenRegistry with a specific package name
$ABIGEN_PATH --abi tmp/TokenRegistry.abi.json \
       --pkg tokenregistry \
       --out bindings/tokenregistry/tokenregistry.go \
       --type TokenRegistry

# Generate bindings for TokenEdits with a different package name
$ABIGEN_PATH --abi tmp/TokenEdits.abi.json \
       --pkg tokenedits \
       --out bindings/tokenedits/tokenedits.go \
       --type TokenEdits

# Clean up temporary files
rm -rf tmp 