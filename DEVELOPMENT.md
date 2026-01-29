# Development Guide

## Quick Start

### First Time Setup

```bash
# Clone the repository
git clone <repository-url>
cd ns-tui

# Install dependencies
go mod download

# Build
task build

# Run
task run
```

## Building

### Using Task (Recommended)

```bash
task build           # Standard build
task release         # Optimized release build
task clean           # Clean build artifacts
```

### Using Go Directly

```bash
go build -o bin/ns-tui ./cmd/ns-tui
```

### Using Nix

```bash
# First time - update vendorHash
nix build 2>&1 | grep "got:" | awk '{print $2}'
# Copy the hash and update vendorHash in flake.nix

# Then build
nix build

# Or run directly
nix run
```

## Project Structure

```
├── cmd/ns-tui/          # Main entry point
├── internal/            # Internal packages
│   ├── api/            # API client for NixOS search
│   ├── models/         # Data models
│   ├── styles/         # UI styles (Lipgloss)
│   └── ui/             # UI components (Bubbletea)
├── bin/                # Build output
├── Taskfile.yml        # Task automation
└── flake.nix          # Nix build config
```

## Development Workflow

### Making Changes

1. **Edit code** in `internal/` or `cmd/`
2. **Format**: `task fmt`
3. **Lint**: `task lint`
4. **Build**: `task build`
5. **Test**: `task run`

### Adding Features

#### New UI Component
1. Add styles to `internal/styles/styles.go`
2. Add rendering to `internal/ui/view.go`
3. Add interactions to `internal/ui/update.go`

#### New API Endpoint
1. Add method to `internal/api/client.go`
2. Add models to `internal/models/` if needed
3. Wire up in UI

#### New Mode
1. Add to `internal/models/mode.go`
2. Update `internal/ui/update.go` for keybindings
3. Update `internal/ui/view.go` for rendering

## Code Style

### Formatting
- Use `gofmt` (run `task fmt`)
- Follow standard Go conventions
- Keep functions small and focused

### Naming
- Use descriptive names
- Follow Go naming conventions
- Exported types/functions start with uppercase
- Internal use lowercase

### Comments
- Document exported functions
- Use full sentences
- Explain "why", not "what"

## Testing

```bash
task test     # Run all tests
go test ./... # Direct go test
```

### Writing Tests
- Place tests next to code (`_test.go`)
- Use table-driven tests
- Mock external dependencies

## Debugging

### Print Debugging
```go
// Don't use fmt.Print in TUI!
// Use a log file instead
f, _ := os.Create("/tmp/ns-tui.log")
log.SetOutput(f)
log.Printf("Debug: %v", value)
```

### Delve Debugger
```bash
dlv debug ./cmd/ns-tui
```

## Common Tasks

### Update Dependencies
```bash
go get -u ./...
task tidy
```

### Add New Dependency
```bash
go get github.com/example/package
task tidy
```

### Update Nix Flake
```bash
nix flake update
# Update vendorHash in flake.nix
```

## Performance

### Profiling
```bash
go build -gcflags="-m" ./cmd/ns-tui  # Escape analysis
go test -cpuprofile cpu.prof         # CPU profiling
go test -memprofile mem.prof         # Memory profiling
```

### Optimization Tips
- Minimize allocations in hot paths
- Use string builders for concatenation
- Cache expensive operations
- Profile before optimizing

## Release Process

1. Update version in relevant files
2. Run full test suite: `task test`
3. Build release binary: `task release`
4. Test release binary
5. Create git tag
6. Push tag to trigger CI/CD

## Troubleshooting

### Build Fails
```bash
task clean
go mod tidy
task build
```

### Import Errors
```bash
go mod tidy
go mod download
```

### Nix Build Fails
- Update `vendorHash` in `flake.nix`
- Run `nix flake check`

## Resources

- [Bubbletea Docs](https://github.com/charmbracelet/bubbletea)
- [Lipgloss Docs](https://github.com/charmbracelet/lipgloss)
- [Task Documentation](https://taskfile.dev/)
- [Go Style Guide](https://go.dev/doc/effective_go)
