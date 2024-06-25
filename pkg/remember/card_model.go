package remember

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
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
			return Model, nil
		}
	}
	return c, nil
}

func (c card) View() string {
	return fmt.Sprintf("%s\n", c.answer)
}
