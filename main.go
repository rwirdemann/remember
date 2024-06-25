package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwirdemann/remember/pkg"
	"os"
)

func main() {
	p := tea.NewProgram(pkg.Remember, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
