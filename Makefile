SHELL := /bin/bash

KUBERNETES_DEPLOYMENT_SCRIPT := ./scripts/setup-kubernetes.sh
PLAYGROUND_SERVER_SCRIPT := ./scripts/setup-local-server.sh
UNINSTALL_SCRIPT := ./scripts/uninstall-kubernetes.sh

# Targets
.PHONY: kubernetes-deployment playground-server

# Target to run the Kubernetes deployment script
kubernetes-deployment:
	@bash $(KUBERNETES_DEPLOYMENT_SCRIPT)

# Target to run the local server script
# It is used to quick run of application without overhead of docker and kubernetes deployment
local-server:
	@bash $(PLAYGROUND_SERVER_SCRIPT)

# Delete all the file store resources deployed in k8s
uninstall-kubernetes:
	@bash $(UNINSTALL_SCRIPT)
