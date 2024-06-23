# Todo-Go

## Overview

- This is a web server for creating, listing, updating, and deleting tasks.

- Data storage uses an in-memory mechanism, meaning it will only exist during runtime.

- For the API documentation, please refer to [Swagger](./cmd/todo/docs/swagger.yaml)

## Project Structure

- Follow [Golang Project Layout](https://github.com/golang-standards/project-layout)

- `mock` directories store generated files by [gomock](https://github.com/uber-go/mock) for unit tests.

```sh
.
├── cmd                   # Applications
│   └── todo
│       ├── config        # Config files for the application
│       └── docs          # Documents for the application
├── internal              # Private application and library code. only use for specific cases
│   ├── controller        # Put business log here
│   │   └── mock
│   ├── db                # Implement in-memory storage mechanism
│   │   └── mock
│   ├── https
│   │   ├── server        # Server instance and middleware
│   │   │   └── handler   # Handle converting data from request
│   │   └── test          # HTTP Server integration tests
│   ├── models            # Data types
│   └── store             # Handle converting data from db
│       └── mock
└── pkg                   # Library code that's ok to use by external applications
    ├── config            # Reading config from env, file and default values.
    ├── http
    │   └── server        # HTTP Server Middleware
    ├── logger
    └── validator
```

## How to run

### Pull image from DockerHub

```sh
docker run -p 8080 dragon0huang/todo-go:latest
```

### Docker Compose

```sh
docker compose up
```

### Local run

```
make todo
```

## How to build

```sh
make build
```

## How to Release

- Create a git tag with `${component}-v${version}` and push. Example:

```sh
git tag todo-v1.0.0
git push origin todo-v1.0.0
```

## Environments

1. Set Env

```sh
export HTTP_SERVER__ADDR_PORT=127.0.0.1:8080
```

2. Using [toml config file](./cmd/todo/config/config.toml)

#### Note

- Env will override variables from the config file
- Set `OPERATION__LOG_LEVEL` to `debug` will also set logger to debug mode.

## GitHub Actions (CI)

- When you push to specified branches or create a pull request, two workflows are triggered:
  1. Validate the code by running linter checks, tests, and builds.
  2. Build two types of Docker images: distroless and alpine for debugging.

## Contact

For any questions or suggestions, feel free to contact me at j0918023423@gmail.com.
