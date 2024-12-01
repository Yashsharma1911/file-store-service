# File Store System CLI

`store` is a command-line tool for interacting with your store service. This README will guide you through installation, setup, and usage of the tool.



### Prerequisites

Before you start, ensure that you have the following tools installed on your system:

- [Git](https://git-scm.com/)
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop) (Optional, if you're using Docker to run MinIO)

# Installation
## Quick start
You can run CLI application and use public minio server to host file, this is effect to run and test CLI application quickly

Clone the repository to your local machine using the following command:
```bash
git clone https://github.com/Yashsharma1911/file-store-service
cd file-store-service
```
After clone get successful, run the below command (Note: `make` command should be installed in your system if not run next command of it)

```bash
make local-server
```

**[Alternative]** : If `make` is not installed in your system, copy this and run in your root dir of repo if `make` is not installed in your system
```bash
./scripts/setup.sh
```

Now open a new terminal and test CLI by adding a file to server
```bash
store add [file path]
```

Use below command to see list of stored files
```bash
store ls
```

## Deploy file store in Kubernetes (Optional)

You can deploy file store to your kubernetes cluster, run below make command which will auto deploy kubernetes resources to run your application

```bash
make kubernetes-deployment
```

Use below command to uninstall resources

```bash
make uninstall-deployment
```

**Note**: Sometimes there can be issue with pod deployment, if it doesn't auto started, try to uninstall resources and install again

## Deploy file store in Docker (Optional)

Run below command to deploy application constainers, it will start Minio container exposed at `:9000` port and will start server container at `:30000` port

```bash
make docker-up
```
**[Alternative]** : if `make` is not installed
```bash
docker compose up
```

# System Achitecture
<img src="https://github.com/user-attachments/assets/6ca653c2-95fd-4a9a-90e1-d86c01b851ab" alt="Alt text" width="700"/>
