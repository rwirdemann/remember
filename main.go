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

	var f *os.File
	var err error
	if exists(name) {
		f, err = os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f, err = os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	writer := bufio.NewWriter(f)

	model, err := pkg.NewListModel(reader, writer)
	if err != nil {
		log.Fatal(err)
	}
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func exists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}
