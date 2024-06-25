package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwirdemann/remember/pkg/remember"
	"os"
)

func main() {
	p := tea.NewProgram(remember.Model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
