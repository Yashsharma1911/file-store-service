# File Store System CLI

`store` is a command-line tool for interacting with your store service. This README will guide you through installation, setup, and usage of the tool.

<img src="https://github.com/user-attachments/assets/6ca653c2-95fd-4a9a-90e1-d86c01b851ab" alt="Alt text" width="700"/>


### Prerequisites

Before you start, ensure that you have the following tools installed on your system:

- [Git](https://git-scm.com/)
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop) (Optional, if you're using Docker to run MinIO)

### Step 1: Clone the Repository

Clone the repository to your local machine using the following command:

```bash
git clone https://github.com/Yashsharma1911/file-store-service
cd file-store-service
```
### Step 2: Run playground server

Run `make local-server`, it will start server locally and will connect to public minio server.

Run make command, ensure you gave make command installed
```bash
make local-server
```

If not run below script through your terminal
```bash
./scripts/setup.sh
```

### Step 2: Test

Run a test command to see it `store` is working

```bash
store ls
```

Run `add` command to store file, give a path of file, make sure to use "" if there is space between your file path

```bash
store add [file path]
```

## Deploy file store in Kubernetes (Optional)

You can run `make kubernetes-deployment` to do a auto resource deployment

```bash
make kubernetes-deployment
```

Use below command to uninstall resources

```bash
make uninstall-deployment
```

*Note*: Sometimes there can be issue with pod deployment, if it doesn't auto started, try to uninstall resources and install again

## Deploy file store in Docker (Optional)

You can run `docker-compose up`

```bash
docker-compose up
```
