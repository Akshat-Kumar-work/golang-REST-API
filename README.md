
# To initialize go project
```
go mod init project-name
```

# In go community the project-name convention is
```
github.com/github-user-name/repository-name
```

# For config management use and mapping
```
viper library
install: go get github.com/spf13/viper
```

# For maintaining ,loading env from .env
```
Joho/Godotenv
install: go get github.com/joho/godotenv

variable priority: .env=>yaml=>go struct(zero value if missing)

```

# For structured logging use zap
```
go get go.uber.org/zap
```
# If log log format !=json
```
2026-01-01T00:04:28.162+0530    INFO    logger/logger.go:82     Starting server {"service": "students-api", "env": "local", "storage_path": "storage/storage.db", "address": "localhost:8082"}
```
# If log format == json
```
{"level":"info","ts":1767206134.6356711,"caller":"logger/logger.go:82","msg":"Starting server","service":"students-api","env":"staging","storage_path":"staging/storage/storage.db","address":"staging/localhost:8082"}
```

# If using custom errors
```
Use fmt.Errorf() for operational errors like creation etc to buit a stack
For known err like validations or not found return the err directly coming from repo to server to handler 
```

# Structure 
```
yourproject/
├── cmd/
│   ├── yourapp/
│   │   └── main.go          # Entry point for your app
│   └── worker/
│       └── main.go          # Another entry point (if you have multiple binaries)
│
├── internal/
│   ├── config/              # Configuration loading and management
│   ├── db/                  # Database layer (connection, migrations)
│   ├── service/             # Core business logic
│   ├── api/                 # HTTP handlers, routes
│   └── utils/               # Internal helpers (not for export)
│
├── pkg/
│   ├── logger/              # Public packages that can be reused
│   └── middleware/          # Shared middleware, libraries
│
├── api/
│   ├── openapi.yaml         # API specification files
│   └── proto/               # Protobuf files (if using gRPC)
│
├── scripts/                 # Build, CI/CD, or local dev scripts
│
├── configs/                 # Example config files (YAML, JSON, ENV)
│
├── web/                     # Frontend assets (if applicable)
│
├── test/                    # Integration / end-to-end tests
│
├── go.mod
├── go.sum
└── README.md
```
