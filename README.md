
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
