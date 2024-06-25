package remember

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type createCardModel struct {
	question textinput.Model
	err      error
}

func initialCreateCardModel() createCardModel {
	ti := textinput.New()
	ti.Placeholder = "Question"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return createCardModel{
		question: ti,
		err:      nil,
	}
}

func (m createCardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m createCardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			Model.cards = append(Model.cards, card{
				question: m.question.Value(),
				answer:   "",
			})

			return Model, nil
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.question, cmd = m.question.Update(msg)
	return m, cmd
}

func (m createCardModel) View() string {
	return fmt.Sprintf(
		"What’s the Question?\n\n%s\n\n%s",
		m.question.View(),
		"(esc to quit)",
	) + "\n"
}
