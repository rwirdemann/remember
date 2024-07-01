package pkg

import (
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"io"
)

const (
	StateList = iota
	StateSelected
	StateAdd
	StateEdit
)

type card struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	UUID     string `json:"uuid"`
}

var (
	q string
	a string
)

type ListModel struct {
	cards      []card
	cursor     int
	state      int
	form       *huh.Form
	inputFocus int
}

func InitialModel() ListModel {
	return ListModel{form: newForm()}
}

func newForm() *huh.Form {
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Question").Value(&q),
			huh.NewInput().Title("Answer").Value(&a)),
	)
	f.Init()
	return f
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if m.state == StateSelected {
		return updateCard(msg, m)
	}

	if m.state == StateAdd {
		return updateAdd(msg, m)
	}

	if m.state == StateEdit {
		return updateEdit(msg, m)
	}

	return updateList(msg, m)
}

func updateEdit(msg tea.Msg, m ListModel) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		if q != "" && a != "" {
			m.cards[m.cursor] = card{Question: q, Answer: a}
		}
		m.state = StateList
	}

	return m, tea.Batch(cmds...)
}

func updateAdd(msg tea.Msg, m ListModel) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		if q != "" && a != "" {
			m.cards = append(m.cards, card{
				Question: q,
				Answer:   a,
			})
		}
		m.state = StateList
	}

	return m, tea.Batch(cmds...)
}

func updateList(msg tea.Msg, m ListModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "e":
			q = m.cards[m.cursor].Question
			a = m.cards[m.cursor].Answer
			m.form = newForm()
			m.state = StateEdit
			return m, nil
		case "a":
			a = ""
			q = ""
			m.form = newForm()
			m.state = StateAdd
			return m, nil
		case "d":
			m.cards = append(m.cards[:m.cursor], m.cards[m.cursor+1:]...)
			m.cursor -= 1
			if m.cursor < 0 {
				m.cursor = 0
			}
			return m, nil
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.cards)-1 {
				m.cursor++
			}
		case "enter":
			m.state = StateSelected
			return m, nil
		}
	}
	return m, nil
}

func updateCard(msg tea.Msg, m ListModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			m.state = StateList
			return m, nil
		}
	}
	return m, nil
}

func (m ListModel) View() string {
	switch m.state {
	case StateList:
		return listView(m)
	case StateSelected:
		return cardView(m)
	case StateAdd:
		return addView(m)
	case StateEdit:
		return editView(m)
	default:
		return ""
	}
}

func editView(m ListModel) string {
	s := fmt.Sprintf(
		"\n%s\n",
		m.form.View(),
	)

	return s
}

func addView(m ListModel) string {
	s := fmt.Sprintf(
		"\n%s\n",
		m.form.View(),
	)

	return s
}

func listView(m ListModel) string {
	s := "\n"
	for i, card := range m.cards {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, card.Question)
	}
	s += "\na: new card • e: edit • d: delete • t: train • q: quit\n"
	return s
}

func cardView(m ListModel) string {
	c := m.cards[m.cursor]
	s := fmt.Sprintf("\n%s\n", c.Answer)
	s += "\nenter: back\n"
	return s
}

func Write(m ListModel, writer io.Writer) error {
	bb, err := json.Marshal(m.cards)
	if err != nil {
		return err
	}

	_, err = writer.Write(bb)
	if err != nil {
		return err
	}

	return nil
}

func Read(reader io.Reader, m ListModel) (ListModel, error) {
	bb, err := io.ReadAll(reader)
	if err != nil {
		return ListModel{}, err
	}
	if len(bb) == 0 {
		return ListModel{}, err
	}
	var cards []card
	if err := json.Unmarshal(bb, &cards); err != nil {
		return ListModel{}, err
	}
	m.cards = cards
	return m, nil
}
