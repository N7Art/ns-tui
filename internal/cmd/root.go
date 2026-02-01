package cmd

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/briheet/ns-tui/internal/styles"
	"github.com/briheet/ns-tui/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) int {
	var (
		theme   string
		profile bool
	)

	var cpuFile *os.File

	rootCmd := &cobra.Command{
		Use:   "ns-tui",
		Short: "A terminal UI for searching NixOS packages and Home Manager options",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if !profile {
				return nil
			}

			var err error
			cpuFile, err = os.Create("cpu.pprof")
			if err != nil {
				return fmt.Errorf("creating cpu.pprof: %w", err)
			}

			if err := pprof.StartCPUProfile(cpuFile); err != nil {
				cpuFile.Close()
				return fmt.Errorf("starting CPU profile: %w", err)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if theme != "" {
				styles.SetTheme(theme)
			}

			baseModel := ui.NewModel()
			p := tea.NewProgram(baseModel, tea.WithAltScreen(), tea.WithContext(ctx))

			if _, err := p.Run(); err != nil {
				return fmt.Errorf("running program: %w", err)
			}

			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			if !profile {
				return nil
			}

			pprof.StopCPUProfile()
			cpuFile.Close()

			runtime.GC()

			memFile, err := os.Create("mem.pprof")
			if err != nil {
				return fmt.Errorf("creating mem.pprof: %w", err)
			}
			defer memFile.Close()

			if err := pprof.WriteHeapProfile(memFile); err != nil {
				return fmt.Errorf("writing heap profile: %w", err)
			}

			return nil
		},
	}

	rootCmd.Flags().StringVarP(&theme, "theme", "t", "", "catppuccin theme (mocha, latte, frappe, macchiato)")
	rootCmd.PersistentFlags().BoolVarP(&profile, "profile", "p", false, "write cpu.pprof and mem.pprof profiles")

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
