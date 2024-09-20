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
```
$ go get <package>
# Sample:
$ go get github.com/gin-gonic/gin
```
Package list
github.com/gin-gonic/gin v1.10.0
github.com/spf13/viper v1.19.0  
github.com/stretchr/testify v1.9.0
gorm.io/gorm v1.25.12

## .env
Copy `.env.example` to a new file named `.env`

Run with:  
go run ./cmd/paishe

By default the project should be up and running at: localhost:8080
```
Sample
http://localhost:8080/users/1
```

# DB Setup

## MAC Setup
brew install postgresql@15  

Test if postgresql is installed
```
$ psql
```

Likely there is no user at this point.  
For first time set up
```
$ sudo -i -u postgres psql
$ Password: <Your system password>
$ postgres=#ALTER USER postgres WITH PASSWORD '<your_strong_password>';
```
Can even create a new user. By default we're using postgres as the the default user here

_Suggestion: Download and Setup PGAdmin: https://www.pgadmin.org/download_

For starts if your setup is completed  
```
POST: localhost:8080/users/create
Req payload raw: json
{
    "name": "any name",
    "email": "something@gmail",
    "phone": "87098709"
}

GET: localhost:8080/users/1
```
