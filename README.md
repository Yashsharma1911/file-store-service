# File Store System CLI

The store CLI is a command-line tool for managing files on a server.

This implementation uses MinIO as a mock for a real-world bucket storage system, demonstrating extensibility with services like AWS S3. By default, files are stored in a data/ directory in the file store path. However, you can update location of stored files or use Kubernetes deployment for production deployments as it uses persistent volume by default.

## Example
Below is an example of using the store CLI to manage files. In this demo, Kubernetes is used to run the application, but it's optional. Refer to the Installation section for quick local setup instructions.

![Description of GIF](assets/example-recording.gif)

### Prerequisites

Before you start, ensure that you have the following tools installed on your system:

- [Git](https://git-scm.com/)
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop) (Optional, if you're using Docker to run MinIO)

# Installation
### Clone the Repository
First, clone the repository and navigate to the project directory.
*For Windows users, please use a `bash` terminal.*
```bash
git clone https://github.com/Yashsharma1911/file-store-service
cd file-store-service
```

## Deploy File Store with Docker

Run the following command to deploy the application using Docker:
```bash
make docker-up
```
If `make` is unavailable, use the script directly:
```bash
chmod +x ./scripts/setup-docker.sh
./scripts/setup-docker.sh
```
For Z Shell (~zsh) users, update the path using below command to reflect changes or restart the terminal:
```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && source ~/.zshrc
```
After the server is running, open new terminal and use the CLI commands to interact with the file store:
```bash
store add examples/newText.txt
store add "examples/newText.txt" "examples/test.txt" # For multiple files or paths with spaces

# List stored files
store ls
```

## Deploy file store in Kubernetes (Optional)

You can deploy file store to your Kubernetes cluster, run below make command which will auto-deploy Kubernetes resources to run your application.

**Note:** After installation gets successful do run `kubectl get pods` to ensure deployments are up and running.

```bash
make kubernetes-deployment
```

*Alternatively:*
```bash
chmod +x ./scripts/setup-kubernetes.sh
./scripts/setup-kubernetes.sh
```
*To uninstall resources:*
```bash
make uninstall-deployment
```

# Suported Commands
* `store add [file path] [file2 path]` - Store file to server, if file already exist on server it will return error for that file<br>
* `store ls` - List of all stored files<br>
* `store update [file path]` - Update existing file content, if file not exist it will create a new file<br>
* `store rm [file name]` - Remove file from store.<br>
* `store wc` - Get total number of words present in all files. <br>
* `store freq-words [--limit|-n 10] [--order=dsc|asc]` - List of least or most frequent words, use `--order` to change order of least or most and use `--limit | -n` flag for the number of words you want to check, default is 10 <br>

# System Architecture
<img src="https://github.com/user-attachments/assets/6ca653c2-95fd-4a9a-90e1-d86c01b851ab" alt="Alt text" width="700"/>

# Contribute
The quickest way to get started is by running the server and CLI locally. By default, this stores files in a public MinIO bucket provided by MinIO.

### Unix systems (Linux, Mac)
1. Build and start the server, and install the store CLI:
```bash
chmod +x ./scripts/setup-local-server.sh
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc && source ~/.bashrc
./scripts/setup-local-server.sh
```
For macOS Z Shell users, update the path using below command to reflect changes or restart the terminal:
```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && source ~/.zshrc
```
2. Open a new terminal and test the CLI:
```bash
store add examples/newText.txt
store add "examples/newText.txt" "examples/test.txt" # For multiple files or paths with spaces

# List stored files
store ls
```

### Windows
1. Build and start the server, and install the store CLI:
```bash
./scripts/setup-local-server.sh
```

2. Open a new terminal and test the CLI:
```bash
store add examples/newText.txt
store add "examples/newText.txt" "examples/test.txt" # For multiple files or paths with spaces

# List stored files
store ls
```
