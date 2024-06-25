package pkg

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	question   textinput.Model
	answer     textinput.Model
	focusIndex int
	err        error
}

func NewModel() InputModel {
	tiq := textinput.New()
	tiq.Placeholder = "Question"
	tiq.Focus()

	tia := textinput.New()
	tia.Placeholder = "Answer"

	return InputModel{
		question: tiq,
		answer:   tia,
		err:      nil,
	}
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			Remember.Cards = append(Remember.Cards, ViewModel{
				Question: m.question.Value(),
				Answer:   m.answer.Value(),
			})
			Remember.Cursor = len(Remember.Cards) - 1
			return Remember, nil
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

func (m InputModel) View() string {
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
