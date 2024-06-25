package remember

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type createCardModel struct {
	question   textinput.Model
	answer     textinput.Model
	focusIndex int
	err        error
}

func initialCreateCardModel() createCardModel {
	tiq := textinput.New()
	tiq.Placeholder = "Question"
	tiq.Focus()

	tia := textinput.New()
	tia.Placeholder = "Answer"

	return createCardModel{
		question: tiq,
		answer:   tia,
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
		case tea.KeyTab, tea.KeyEnter:
			m.focusIndex++
			if m.focusIndex > 1 {
				m.focusIndex = 0
			}
			if m.focusIndex == 0 {
				m.answer.Blur()
				return m, m.question.Focus()
			}
			m.question.Blur()
			return m, m.answer.Focus()
		case tea.KeyEsc:
			Model.cards = append(Model.cards, card{
				question: m.question.Value(),
				answer:   m.answer.Value(),
			})
			return Model, nil
		}

	case error:
		m.err = msg
		return m, nil
	}

	if m.focusIndex == 0 {
		m.question, cmd = m.question.Update(msg)
	} else {
		m.answer, cmd = m.answer.Update(msg)
	}
	return m, cmd
}

func (m createCardModel) View() string {
	s := fmt.Sprintf(
		"Question?\n\n%s",
		m.question.View(),
	) + "\n"

	s = s + fmt.Sprintf(
		"\nAnswer?\n\n%s\n\n%s",
		m.answer.View(),
		"(esc to quit)",
	) + "\n"
	return s
}
