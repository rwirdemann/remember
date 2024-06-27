package pkg

import (
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"io"
)

type ListModel struct {
	Cards  []CardModel `json:"cards"`
	Cursor int         `json:"-"`
}

func (m *ListModel) Init() tea.Cmd {
	return nil
}

func (m *ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "a":
			return NewModel(m, CardModel{}), nil

		case "e":
			return NewModel(m, m.Cards[m.Cursor]), nil

		case "t":
			trainModel := NewTrainModel(m, m.Cards)
			return trainModel, nil

		case "d":
			m.Cards = append(m.Cards[:m.Cursor], m.Cards[m.Cursor+1:]...)
			m.Cursor -= 1
			if m.Cursor < 0 {
				m.Cursor = 0
			}
			return m, nil

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

func (m *ListModel) View() string {
	// The header
	s := "\n"

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
	s += "\na: new card • e: edit • d: delete • t: train • q: quit\n"

	// Send the UI for rendering
	return s
}

func (m *ListModel) Write(writer io.Writer) error {
	bb, err := json.Marshal(m.Cards)
	if err != nil {
		return err
	}

	_, err = writer.Write(bb)
	if err != nil {
		return err
	}

	return nil
}

func (m *ListModel) Read(reader io.Reader) error {
	bb, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	if len(bb) == 0 {
		return nil
	}
	var cards []CardModel
	if err := json.Unmarshal(bb, &cards); err != nil {
		return err
	}
	m.Cards = cards
	for i, _ := range m.Cards {
		m.Cards[i].parent = m
	}
	return nil
}

func (m *ListModel) AddOrUpdate(vm CardModel) {
	if len(vm.UUID) == 0 {
		vm.UUID = uuid.NewString()
		m.Cards = append(m.Cards, vm)
		return
	}
	for i, c := range m.Cards {
		if c.UUID == vm.UUID {
			m.Cards[i] = vm
			break
		}
	}
}
