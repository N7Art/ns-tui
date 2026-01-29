package ui

import (
	"time"

	"ns-tui/internal/api"
	"ns-tui/internal/models"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the application state
type Model struct {
	textInput            textinput.Model
	packages             []models.Package
	cursor               int
	scrollOffset         int
	loading              bool
	err                  error
	width                int
	height               int
	lastQuery            string
	searchTimer          *time.Timer
	mode                 models.Mode
	selectedPackage      *models.Package
	apiClient            *api.Client
	selectedInstallMethod int // 0-3 for the 4 install methods
}

// NewModel creates a new application model
func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Search NixOS packages..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 66

	return Model{
		textInput: ti,
		packages:  []models.Package{},
		cursor:    0,
		loading:   false,
		mode:      models.InsertMode,
		apiClient: api.NewClient(),
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// performSearch executes the search in a goroutine
func performSearch(client *api.Client, query string) tea.Cmd {
	return func() tea.Msg {
		packages, err := client.SearchPackages(query)
		return searchResultMsg{packages: packages, err: err}
	}
}
