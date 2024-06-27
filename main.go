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

	// create model and initialize it with contents from file
	in, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	model := &pkg.ListModel{}
	if err := model.Read(bufio.NewReader(in)); err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// write updated model content back to file
	out, err := os.OpenFile(name, os.O_WRONLY|os.O_TRUNC, 0666)
	defer out.Close()
	writer := bufio.NewWriter(out)
	if err := model.Write(writer); err != nil {
		log.Fatal(err)
	}
	if err := writer.Flush(); err != nil {
		log.Fatal(err)
	}
}
