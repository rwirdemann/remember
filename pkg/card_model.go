package pkg

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type CardModel struct {
	parent   *ListModel
	Question string `json:"question"`
	Answer   string `json:"answer"`
	UUID     string `json:"uuid"`
}

func (m CardModel) Init() tea.Cmd {
	return nil
}

func (m CardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			return m.parent, nil
		}
	}
	return m, nil
}

func (m CardModel) View() string {
	s := fmt.Sprintf("\n%s\n", m.Answer)
	s += "\nenter: return\n"
	return s
}
