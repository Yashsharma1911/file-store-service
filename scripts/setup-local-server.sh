#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Define environment variables
# Keys and password are public, provided by minio for testing
# File is no safe with this environment, use only for testing
# For secure file save, use docker deployment
export MINIO_ROOT_USER="Q3AM3UQ867SPQQA43P2F"
export MINIO_ROOT_PASSWORD="zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
export MINIO_ENDPOINT="play.min.io"
export MINIO_BUCKET="testbucket"
export MINIO_USE_SSL=true
export FILE_STORE_SERVER_URL=""
export FILE_STORE_SERVER_URL="http://localhost:30000"
echo "Server starting at : $FILE_STORE_SERVER_URL"
# Build the echo-server binary
echo "Building echo-server..."
go build -o echo-server .
echo "echo-server built successfully."

# Build the store binary
echo "Building store..."
go build -o store ./cmd/store
echo "store built successfully."

# Install the store command globally
echo "Installing store command..."
go install ./cmd/store
echo "store command installed successfully."

echo ""
echo "Yay 🎉 you can now start using "store" command anywhere in you system"
echo "example run: store ls"
echo ""
# Run echo-server
echo "Starting echo-server..."
./echo-server