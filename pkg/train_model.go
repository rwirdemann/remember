package pkg

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
)

type TrainModel struct {
	parent     *ListModel
	cards      []CardModel
	selected   int
	showAnswer bool
}

func (m TrainModel) Init() tea.Cmd {
	return nil
}

func (m TrainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.showAnswer = false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			m.showAnswer = true
		case "c", "f":
			m.selected = rand.Intn(len(m.cards))
		case "enter":
			return m.parent, nil
		}
	}
	return m, nil
}

func (m TrainModel) View() string {
	s := ""
	if m.showAnswer {
		s = fmt.Sprintf("\n%s\n", m.cards[m.selected].Answer)
	} else {
		s = fmt.Sprintf("\n%s\n", m.cards[m.selected].Question)
	}
	s += "\na: show answer • c: correct • f: false • enter: return\n"
	return s
}
