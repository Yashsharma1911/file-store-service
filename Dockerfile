FROM golang:1.22.5

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the app
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Expose the port
EXPOSE 30000

CMD ["./main"]
