package pkg

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewModel struct {
	parent   *ListModel
	Question string
	Answer   string
}

func NewViewModel(parent *ListModel, question, answer string) ViewModel {
	return ViewModel{
		parent:   parent,
		Question: question,
		Answer:   answer,
	}
}

func (m ViewModel) Init() tea.Cmd {
	return nil
}

func (m ViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			return m.parent, nil
		}
	}
	return m, nil
}

func (m ViewModel) View() string {
	s := fmt.Sprintf("\n%s\n", m.Answer)

	// The footer
	s += "\nenter: return\n"
	return s
}
