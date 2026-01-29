# Internal Package Architecture

This directory contains the internal packages for ns-tui, organized by responsibility.

## Package Overview

### `models/`
Data models and types used throughout the application.

- **`package.go`**: Defines the `Package` struct representing a NixOS package with all metadata (name, version, description, license, etc.)
- **`mode.go`**: Defines the `Mode` enum for vim-style modes (Insert, Normal, Detail)

### `api/`
Handles communication with the NixOS search backend.

- **`client.go`**:
  - Elasticsearch client for the NixOS package search API
  - Query builder for package searches
  - Response parsing and data extraction
  - Uses read-only public credentials for `search.nixos.org`

### `styles/`
Lipgloss styling definitions for the UI.

- **`styles.go`**:
  - Catppuccin-inspired color palette
  - Reusable style definitions for all UI components
  - Separate styles for different modes and states
  - Consistent theming across the application

### `ui/`
Bubbletea UI components implementing the Elm Architecture.

- **`model.go`**:
  - Application state (`Model` struct)
  - Model initialization
  - Search command helpers

- **`update.go`**:
  - Message handling and state updates
  - Keyboard input processing
  - Mode switching logic (Insert ↔ Normal ↔ Detail)
  - Vim-style keybinding implementation

- **`view.go`**:
  - UI rendering logic
  - Search view with results
  - Detailed package view
  - Scroll management and pagination

- **`messages.go`**:
  - Custom Bubbletea messages
  - Currently defines `searchResultMsg` for async search results

## Data Flow

```
User Input
    ↓
Update (ui/update.go)
    ↓
Model State Change
    ↓
Search Trigger
    ↓
API Client (api/client.go)
    ↓
Elasticsearch Backend
    ↓
Parse Response
    ↓
Update Model with Results
    ↓
View (ui/view.go)
    ↓
Render UI
```

## Key Design Patterns

### Separation of Concerns
- **Models**: Pure data structures, no logic
- **API**: External communication, isolated from UI
- **Styles**: Visual presentation, separate from logic
- **UI**: Bubbletea architecture, split into model/update/view

### Vim-Style Modal Interface
The application implements three distinct modes:
1. **Insert Mode**: For typing and searching
2. **Normal Mode**: For navigating results with vim keys
3. **Detail Mode**: For viewing package information

### Async Operations
- Search requests are non-blocking
- Uses Bubbletea's command pattern for async operations
- Results delivered via messages

## Adding New Features

### Adding a New UI Component
1. Define styles in `styles/styles.go`
2. Add rendering logic to `ui/view.go`
3. Add interaction handling to `ui/update.go`

### Adding a New API Endpoint
1. Add method to `api/Client` in `client.go`
2. Define response parsing logic
3. Update `models/` if new data structures are needed

### Adding a New Mode
1. Add mode constant to `models/mode.go`
2. Update mode switch logic in `ui/update.go`
3. Add view rendering in `ui/view.go`

## Testing Guidelines

- Models should have no dependencies and be easily testable
- API client should mock HTTP responses for tests
- UI components should test state transitions
- Integration tests should verify full workflows

## Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Terminal styling
- `github.com/charmbracelet/bubbles` - TUI components (textinput)
- Standard library for HTTP, JSON, etc.
