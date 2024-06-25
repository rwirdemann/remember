package pkg

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type ListModel struct {
	Cards  []ViewModel
	Cursor int // which card our cursor is pointing at
}

var Model ListModel

func init() {
	Model = ListModel{
		Cards: []ViewModel{
			{
				Question: "Wie verändert man das JSON-Marshaling-Verhalten eines Typs?",
				Answer:   "Indem das Interface json.Marshaler implementiert wird.",
			},
			{
				Question: "Wie werden die Methoden einen eingebnetteten Typs aufgerufen?",
				Answer:   "Direkt auf dem umschliessenden Typ.",
			},
		},
	}
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// These keys should exit the program.
		case "n":
			return NewModel(), nil

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Cursor < len(m.Cards)-1 {
				m.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			return m.Cards[m.Cursor], nil
		}
	}

	// Return the updated InputModel to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil

}

func (m ListModel) View() string {
	// The header
	s := ""

	for i, card := range m.Cards {

		// Is the cursor pointing at this card?
		cursor := " " // no cursor
		if m.Cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, card.Question)
	}

	// The footer
	s += "\nn: new card • q: quit\n"

	// Send the UI for rendering
	return s
}
