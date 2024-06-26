package main

import (
	"bufio"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwirdemann/remember/pkg"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("usage: remember [file]")
		os.Exit(1)
	}
	name := os.Args[1]
	if !strings.HasSuffix(name, ".json") {
		name = fmt.Sprintf("%s.json", name)
	}
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	writer := bufio.NewWriter(f)
	p := tea.NewProgram(pkg.NewListModel(writer), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
