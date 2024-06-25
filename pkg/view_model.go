package pkg

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewModel struct {
	Question string
	Answer   string
}

func (c ViewModel) Init() tea.Cmd {
	return nil
}

func (c ViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			return Model, nil
		}
	}
	return c, nil
}

func (c ViewModel) View() string {
	return fmt.Sprintf("%s\n", c.Answer)
}
