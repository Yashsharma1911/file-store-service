# file-store-service

Go-Service/
├── cmd/
│   ├── api/
│   │   └── main.go           # Entry point of the application
├── server/
│   ├── database/             # Database initialization and migrations
│   │   ├── db.go
│   │   └── migrations/
│   ├── dataAccess/           # Data access logic (factory pattern)
│   │   ├── file.go
│   │   └── dataAccess.go
│   ├── handlers/             # HTTP request handlers
│   │   ├── fileHandler.go
│   │   └── handlers.go
│   ├── router.go             # Route definitions
│   └── config/               # Configuration files (DB, app config)
│       └── config.go
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
└── README.md