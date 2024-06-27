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
	memorized  int
	checked    map[int]struct{}
}

func NewTrainModel(parent *ListModel, cards []CardModel) TrainModel {
	m := TrainModel{parent: parent, cards: cards}
	m.checked = make(map[int]struct{})
	m.selected = rand.Intn(len(cards))
	return m
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
		case "f":
			m.selected = m.selectCard()
		case "c":
			m.checked[m.selected] = struct{}{}
			m.memorized = m.memorized + 1
			if m.memorized == len(m.cards) {
				return m.parent, nil
			}
			m.selected = m.selectCard()
		case "enter":
			return m.parent, nil
		}
	}
	return m, nil
}

func (m TrainModel) selectCard() int {
	for {
		i := rand.Intn(len(m.cards))
		if _, ok := m.checked[i]; !ok {
			return i
		}
	}
}

func (m TrainModel) View() string {
	s := fmt.Sprintf("%d of %d cards memorized\n", m.memorized, len(m.cards))
	if m.showAnswer {
		s += fmt.Sprintf("\n%s\n", m.cards[m.selected].Answer)
	} else {
		s += fmt.Sprintf("\n%s\n", m.cards[m.selected].Question)
	}
	s += "\na: show answer • c: correct • f: false • enter: return\n"
	return s
}
