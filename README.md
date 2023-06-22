# damasdev/fiber

Go Fiber with Hexagonal (Ports & Adapters) Architecture

## Project Structure

```
├── cmd
├── internal
│   ├── core
│   │   ├── domain
│   │   ├── port
│   │   └── service
│   ├── infrastructure
│   └── interface
├── pkg
└── test
    └── mocks
```

### internal/core

All the core components or business logic (service, domain and port).

### internal/infrastructure

All the driven adapters (external dependencies and implementation details).

### internal/interfaces

All the driver adapters for communicating with external systems.