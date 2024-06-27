package pkg

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	question   textinput.Model
	answer     textinput.Model
	uuid       string
	focusIndex int
	err        error
	parent     *ListModel
}

func NewModel(parent *ListModel, model ViewModel) InputModel {
	tiq := textinput.New()
	tiq.Placeholder = "Question"
	tiq.Focus()

	tia := textinput.New()
	tia.Placeholder = "Answer"

	uuid := ""
	if len(model.UUID) > 0 {
		tiq.SetValue(model.Question)
		tia.SetValue(model.Answer)
		uuid = model.UUID
	}

	return InputModel{
		parent:   parent,
		question: tiq,
		answer:   tia,
		uuid:     uuid,
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
			if m.focusIndex == 1 {
				vm := ViewModel{
					parent:   m.parent,
					Question: m.question.Value(),
					Answer:   m.answer.Value(),
					UUID:     m.uuid,
				}
				vm.parent.AddOrUpdate(vm)
				m.parent.Cursor = len(m.parent.Cards) - 1
				return m.parent, nil
			}

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
		"\nAnswer?\n\n%s\n",
		m.answer.View(),
	)

	s += "\nenter: focus next\n"

	return s
}
