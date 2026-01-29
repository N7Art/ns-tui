package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/briheet/ns-tui/internal/styles"
	"github.com/briheet/ns-tui/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

var version = "dev"

func main() {
	theme := flag.String("theme", "", "catppuccin theme (mocha, latte, frappe, macchiato)")
	flag.Parse()

	if *theme != "" {
		styles.SetTheme(*theme)
	}

	m := ui.NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
