# Jokenpo with AI

A simple Rock-Paper-Scissors (Jokenpo) game implemented in Go using Hexagonal Architecture.

## Features
- Play against a random AI.
- Clean separation of concerns using Hexagonal Architecture.
- CLI interface for user interaction.

## Architecture
The project follows the Hexagonal Architecture (Ports and Adapters) pattern:

- **Core (Domain)**: Business logic and entities (Choices, Results, Game Rules).
- **Core (Ports)**: Interfaces defining inbound and outbound communication.
- **Core (Services)**: Implementation of the game logic.
- **Adapters (Inbound)**: CLI handler for user interaction.
- **Adapters (Outbound)**: AI provider (random strategy).

## Project Structure
```text
.
├── internal
│   ├── adapters
│   │   ├── ai          # Driven adapters (AI logic)
│   │   └── handlers    # Driver adapters (CLI)
│   └── core
│       ├── domain      # Domain entities and logic
│       ├── ports       # Port interfaces
│       └── services    # Application services
├── main.go             # Application entry point and dependency wiring
└── go.mod              # Go module definition
```

## How to Run
Ensure you have Go installed on your machine.

1. Clone the repository (if applicable).
2. Run the application:
   ```bash
   go run main.go
   ```

## How to Test
Run the unit tests for the domain logic:
```bash
go test ./internal/core/domain/...
```
