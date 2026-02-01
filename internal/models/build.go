package models

import (
	"github.com/briheet/ns-tui/internal/styles"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func NewTextInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Search NixOS packages..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 66
	return ti
}

func NewSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(styles.ColorPink)
	return s
}
