package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type card struct {
	question string
	answer   string
}

func (c card) Init() tea.Cmd {
	return nil
}

func (c card) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			return m, nil
		}
	}
	return c, nil
}

func (c card) View() string {
	return fmt.Sprintf("%s\n", c.answer)
}

type model struct {
	cards    []card
	cursor   int              // which card our cursor is pointing at
	selected map[int]struct{} // which cards are selected
}

var m model

func init() {
	m = model{
		cards: []card{
			{
				question: "Wie verändert man das JSON-Marshaling-Verhalten eines Typs?",
				answer:   "Indem das Interface json.Marshaler implementiert wird.",
			},
			{
				question: "Wie werden die Methoden einen eingebnetteten Typs aufgerufen?",
				answer:   "Direkt auf dem umschliessenden Typ.",
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.cards)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			return m.cards[m.cursor], nil
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil

}

func (m model) View() string {
	// The header
	s := ""

	for i, card := range m.cards {

		// Is the cursor pointing at this card?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, card.question)
	}

	// The footer
	s += "\nPress q to quit, n to add a new card\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
