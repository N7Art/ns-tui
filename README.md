# ns-tui ⚡

A beautiful Terminal User Interface (TUI) for searching NixOS packages in real-time, built with Go and [Bubbletea](https://github.com/charmbracelet/bubbletea).

<div align="center">

**Fast • Beautiful • Real-time**

</div>

## Features

- 🔍 **Real-time search** - Results appear as you type with 300ms debouncing
- 🎨 **Beautiful UI** - Catppuccin-inspired color scheme with centered layout
- ⚡ **Fast** - Powered by the official NixOS Elasticsearch backend
- ⌨️ **Vim keybindings** - Full vim-style navigation with Insert/Normal modes
- 📊 **Detailed package view** - Press Enter to see full package information including:
  - Description, license, homepage
  - Available programs and ALL supported platforms
  - Maintainers, teams, and metadata
  - **4 installation methods**: nix-env, nix profile, NixOS config, nix-shell
- 🔢 **Smart navigation** - j/k in Normal mode, or arrow keys in Insert mode
- 📦 **Comprehensive search** - Searches package names, programs, and descriptions
- 💻 **Responsive** - Adapts to your terminal size
- 🎯 **Accurate** - Uses the same backend as search.nixos.org

## Installation

### Prerequisites

- Go 1.25 or later
- [Task](https://taskfile.dev/) (optional, for task automation)
- Nix (optional, for Nix builds)

### Build from source

#### Using Task (recommended)

```bash
git clone <repository-url>
cd ns-tui
task build
```

#### Using Go directly

```bash
go build -o bin/ns-tui ./cmd/ns-tui
```

#### Using Nix

```bash
nix build
# or run directly
nix run
```

## Usage

Run the application:

```bash
./bin/ns-tui
# or with Task
task run
```

### Quick Start

1. **Search** - Start typing to search (you're in Insert mode by default)
2. **Navigate** - Press `Esc` to enter Normal mode, then use `j`/`k` to navigate
3. **View Details** - Press `Enter` in Normal mode to see full package info
4. **Go Back** - Press `q` or `Esc` to return from detail view
5. **Resume Searching** - Press `i` to return to Insert mode and modify your search

### Keyboard Controls

The application uses **vim-style keybindings** with Insert and Normal modes:

#### Insert Mode (Default)
When you start the app, you're in Insert mode and can type to search.

| Key | Action |
|-----|--------|
| **Type** | Search for packages in real-time |
| **↑/↓** | Navigate through search results |
| **Esc** | Switch to Normal mode |
| **q** / **Ctrl+C** | Quit the application |

#### Normal Mode
Press `Esc` to enter Normal mode for vim-style navigation.

| Key | Action |
|-----|--------|
| **i** | Return to Insert mode |
| **j** | Move down |
| **k** | Move up |
| **g** | Jump to top |
| **G** | Jump to bottom |
| **Enter** | View package details |
| **q** | Quit the application |

#### Detail View
When viewing package details:

| Key | Action |
|-----|--------|
| **Esc** / **Backspace** / **b** | Return to search results |
| **q** | Quit the application |

### Example Workflow

```
1. Launch the app
   $ ./bin/ns-tui

2. Type "vim" to search
   (You're in INSERT mode - notice the green indicator)

3. Press Esc to enter NORMAL mode
   (Notice the blue indicator)

4. Press j/k to navigate through results

5. Press Enter on a package to view details
   (Shows description, license, programs, platforms, maintainers, etc.)
   (Shows all 4 installation methods: nix-env, nix profile, NixOS config, nix-shell)

6. Press Esc/Backspace/b to return to search
   (Or press q to quit the application entirely)

7. Press i to modify your search
   (Back to INSERT mode)

8. Press q anytime to quit
```

### Example Searches

Try searching for:
- `vim` - Text editor with detailed configuration options
- `firefox` - Web browser
- `python3` - Python interpreter and packages
- `nodejs` - JavaScript runtime
- `git` - Version control system
- `docker` - Container platform

## How it works

The application uses the same Elasticsearch backend as [search.nixos.org](https://search.nixos.org), providing real-time search across the NixOS package repository. It queries the official NixOS package index and displays results with:

- Package name and attribute name
- Version information
- Package description
- Relevance-based sorting

## Technical Details

### Stack
- Built with [Bubbletea](https://github.com/charmbracelet/bubbletea) - The Elm Architecture for Go
- Styled with [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions for terminal UIs
- Uses [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components for Bubbletea

### Backend
- **API Endpoint**: `https://search.nixos.org/backend`
- **Index**: `latest-44-nixos-unstable`
- **Backend**: Elasticsearch (sponsored by Bonsai.io)
- **Authentication**: Basic auth (read-only public credentials)

### Design
- **Color Scheme**: Catppuccin-inspired palette
- **Layout**: Centered, responsive design
- **Performance**: 300ms debounced search, displays up to 8-12 results per view
- **Navigation**: Vim-style modal interface (Insert/Normal/Detail modes)
- **Scrolling**: Intelligent viewport management with scroll indicators

### Modes
The application implements a vim-like modal interface:
- **Insert Mode**: For searching and typing queries (green indicator)
- **Normal Mode**: For navigating results with vim keys (blue indicator)
- **Detail Mode**: For viewing comprehensive package information

## Development

### Project Structure

```
ns-tui/
├── cmd/
│   └── ns-tui/        # Application entry point
│       └── main.go
├── internal/          # Internal packages
│   ├── api/          # Elasticsearch client
│   │   └── client.go
│   ├── models/       # Data models
│   │   ├── mode.go
│   │   └── package.go
│   ├── styles/       # Lipgloss styles
│   │   └── styles.go
│   └── ui/           # Bubbletea UI components
│       ├── messages.go
│       ├── model.go
│       ├── update.go
│       └── view.go
├── bin/              # Compiled binaries
├── Taskfile.yml      # Task automation
├── flake.nix         # Nix build configuration
└── README.md
```

### Available Tasks

```bash
task              # List all available tasks
task build        # Build the application
task run          # Build and run
task dev          # Run in development mode (hot reload)
task test         # Run tests
task clean        # Clean build artifacts
task fmt          # Format code
task lint         # Run linter
task release      # Build optimized release binary
```

### Making Changes

1. Edit files in `internal/` or `cmd/`
2. Run `task build` to compile
3. Run `task run` to test
4. Run `task fmt` and `task lint` before committing

## License

MIT

## API Credentials

The application uses publicly available read-only credentials for the NixOS search backend. These credentials are intentionally public and embedded in the [search.nixos.org](https://search.nixos.org) frontend source code for community access.

## Acknowledgments

- [Charmbracelet](https://github.com/charmbracelet) for the excellent TUI libraries
- [NixOS](https://nixos.org) for the package search infrastructure and open API
- [Bonsai.io](https://bonsai.io) for sponsoring the Elasticsearch backend
- [Catppuccin](https://github.com/catppuccin/catppuccin) for color scheme inspiration
