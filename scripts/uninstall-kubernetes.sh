#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "Starting uninstallation process..."

# Delete Kubernetes resources
echo "Deleting file store Kubernetes resources..."
kubectl delete -f configs/kubernetes/file-store-service.yaml || echo "file-store-service.yaml not found or already deleted"
kubectl delete -f configs/kubernetes/file-store-deployment.yaml || echo "file-store-deployment.yaml not found or already deleted"
kubectl delete -f configs/kubernetes/minio-service.yaml || echo "minio-service.yaml not found or already deleted"
kubectl delete -f configs/kubernetes/minio-deployment.yaml || echo "minio-deployment.yaml not found or already deleted"
kubectl delete -f configs/kubernetes/minio-pvc.yaml || echo "minio-pvc.yaml not found or already deleted"
kubectl delete -f configs/kubernetes/minio-pv.yaml || echo "minio-pv.yaml not found or already deleted"
kubectl delete -f configs/kubernetes/minio-secret.yaml || echo "minio-secret.yaml not found or already deleted"

echo "Kubernetes resources deleted."
