#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Define environment variables
# Keys and password are public, provided by minio for testing
# File is no safe with this environment, use only for testing
# For secure file save, use docker deployment
export FILE_STORE_SERVER_URL="http://localhost:30000"
echo "Server starting at : $FILE_STORE_SERVER_URL"
echo "Applying Kubernetes resources..."
kubectl apply -f configs/kubernetes/minio-secret.yaml
kubectl apply -f configs/kubernetes/minio-pv.yaml
kubectl apply -f configs/kubernetes/minio-pvc.yaml
echo ""
kubectl apply -f configs/kubernetes/minio-deployment.yaml
kubectl apply -f configs/kubernetes/minio-service.yaml
echo ""
echo "Waiting for MinIO to be fully deployed..."
sleep 20  # Sleep for 10 seconds to allow MinIO to be ready
# Apply deployment and service YAML files
kubectl apply -f configs/kubernetes/file-store-deployment.yaml
kubectl apply -f configs/kubernetes/file-store-service.yaml
echo ""
echo "Waiting for server to start..."
sleep 20  # 
# Build the store binary
echo ""
echo "Building store..."
go build -o store ./cmd/store
echo "store built successfully."

# Install the store command globally
echo ""
echo "Installing store command..."
go install ./cmd/store
echo "store command installed successfully."

echo ""
echo "Yay 🎉 you can now start using "store" command anywhere in you system"
echo "example run: store ls"
echo ""