#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Build the store binary
echo "Building store CLI..."
go build -o store ./cmd/store
echo "✅ store built successfully."
echo ""
# Install the store command globally
echo "Installing store command..."
go install ./cmd/store
echo "✅ store command installed successfully."
echo ""

echo ""
docker compose up
echo ""