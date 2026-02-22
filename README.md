# gunginx
My hobby project to create a load balancer from scratch


# Project Structure
```
gunginx/ 
├── go.mod
├── cmd/
│   ├── gunginx-server/
│   │   ├── main.go       
│   │   └── plugins.go    <-- NEW: The file where users "enable" plugins
│   └── gunginx/          
│       └── main.go       
├── internal/
│   └── engine/           <-- Core load balancing logic (Round Robin, Mutexes)
├── pkg/
│   ├── api/              
│   └── sdk/              <-- NEW: The Plugin Interface (The Contract)
│       └── plugin.go     <-- Defines `type Middleware interface { ... }`
└── plugins/              <-- NEW: The actual out-of-the-box implementations
    ├── logger/
    │   └── logger.go     <-- Implements the SDK interfaces
    └── notify-discord/
        └── discord.go    <-- Implements the SDK interfaces
```