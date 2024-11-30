# File Store System CLI

`store` is a command-line tool for interacting with your store service. This README will guide you through installation, setup, and usage of the tool.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)

## Installation

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
### Step 2: Run setup file

Run setup file for quick start, it will start server locally and will connect to public minio server.

```bash
./setup.sh
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
