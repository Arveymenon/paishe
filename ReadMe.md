# Backend Structure

### Project Structure
```
/paishe
├── cmd/                # Entry points for your application (main package)
│   └── paishe
│       └── main.go     # The main function, starting point of the API
├── internal/           # Internal packages (for app logic that shouldn't be exposed)
│   ├── config/         # Configuration management (env, config files)
│   ├── domain/         # Domain models (entities, structs, core logic)
│   ├── repository/     # Data persistence (DB interactions, third-party services)
│   ├── service/        # Business logic (core logic not tied to transport layer)
│   └── transport/      # Transport layer (HTTP API, gRPC, etc.)
│       ├── http/       # HTTP-specific routes, handlers, middlewares
│       └── middleware/ # Middleware for logging, auth, etc.
├── pkg/                # External packages (can be used by other applications)
├── api/                # API specifications (OpenAPI/Swagger files)
├── go.mod              # Go module file
└── go.sum              # Dependency versions
```

### Packages to be installed:
github.com/gin-gonic/gin v1.10.0  
github.com/spf13/viper v1.19.0  
github.com/stretchr/testify v1.9.0  
gorm.io/gorm v1.25.12  


Run with:  
go run ./cmd/paishe


# DB Setup