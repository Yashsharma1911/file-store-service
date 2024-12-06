# Problem Statement 1: File Store System CLI

`store` is a command-line tool to manage your files through server.

I've used MinIO to mock the real life implementation of bucket to store files in server instead of storing them in file store path, MinIO also has support for S3 AWS bucket which shows extensibility of this implementation. However, if you use docker to run server it will store uploaded files in current file store path in `data/` dir.

### Example
In this video I used kubernetes to run application however it is optional, checkout Installation part to run application locally quickly
<br>
![Description of GIF](assets/example-recording.gif)

### Prerequisites

Before you start, ensure that you have the following tools installed on your system:

- [Git](https://git-scm.com/)
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop) (Optional, if you're using Docker to run MinIO)

# Installation
## Quick start
This will store your files in public MinIO bucket instead of file store itself, use docker (Installation is mentioned in Readme below) to run the server if you want to store files in current file store path.

### Unix systems (Linux, Mac)
1. Run below command, it will download and start the server and also install `store` CLI to your system:
```bash
git clone https://github.com/Yashsharma1911/file-store-service
cd file-store-service
chmod +x ./scripts/setup-local-server.sh
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc && source ~/.bashrc
./scripts/setup-local-server.sh
```

**Note:** It is suggested to use default (bash) terminal of your system. In case you're using any customized shell ensure to restart it, if you're using Z shell in Mac run `echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && source ~/.zshrc` to take effect or restart the terminal.

2. **Now open a new terminal** and test CLI by adding a file to server.
```bash
store add [file path]

# If you are in root dir of project you can run the below command also
store add examples/newText.txt

# Use "" in file path if you want to upload two files or if there is space or a special character in your path
store add "examples/newText.txt" "examples/test.txt"
```

Use below command to see list of stored files.
```bash
store ls
```



### Windows
1. Run below command, it will download and start the server and also install `store` CLI to your system:
```bash
git clone https://github.com/Yashsharma1911/file-store-service
cd file-store-service
./scripts/setup-local-server.sh
```

3. **Now open a new terminal** and test CLI by adding a file to server.
```bash
store add [file path]

# If you are in root dir of project you can run the below command also
store add examples/newText.txt

# Use "" in file path if you want to upload two files or if there is space or a special character in your path
store add "examples/newText.txt" "examples/test.txt"
```

Use below command to see list of stored files.
```bash
store ls
```

## Deploy file store in Kubernetes (Optional)

You can deploy file store to your Kubernetes cluster, run below make command which will auto-deploy Kubernetes resources to run your application.

**Note:** After installation gets successful do run `kubectl get pods` to ensure deployments are up and running.

```bash
make kubernetes-deployment
```

**[Alternative]** : if `make` is not installed run below command in root dir of the project
```bash
./scripts/setup-kubernetes.sh
```

Use below command to uninstall resources.

```bash
make uninstall-deployment
```

## Deploy file store in Docker (Optional)

Run below command to deploy application containers, it will start Minio container exposed at `:9000` port and will start server container at `:30000` port. After you get server started open new terminal and run `store add [file path]` command to check you're able to add files. Run `store ls` to check uploaded files or checkout `./data/testbucket` dir in current file store path to see uploaded files .

```bash
make docker-up
```
**[Alternative]** : if `make` is not installed.
```bash
chmod +x ./scripts/setup-docker.sh
./scripts/setup-docker.sh
```
#### Additional (Optional)
Checkout File Store CLI [docker image](https://hub.docker.com/repository/docker/yashsharma1911/file-store/general) for any additional updates.

# Suported Commands
`store add [file path] [file2 path]` - Store file to server, if file already exist on server it will return error for that file<br>
`store ls` - List of all stored files<br>
`store update [file path]` - Update existing file content, if file not exist it will create a new file<br>
`store rm [file name]` - Remove file from store.<br>
`store wc` - Get total number of words present in all files. <br>
`store freq-words [--limit|-n 10] [--order=dsc|asc]` - List of least or most frequent words, use `--order` to change order of least or most and use `--limit | -n` flag for the number of words you want to check, default is 10 <br>

# System Architecture
<img src="https://github.com/user-attachments/assets/6ca653c2-95fd-4a9a-90e1-d86c01b851ab" alt="Alt text" width="700"/>
