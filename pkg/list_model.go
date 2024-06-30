package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"io"
)

const (
	StateList = iota
	StateSelected
	StateAdd
)

type card struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	UUID     string `json:"uuid"`
}

type ListModel struct {
	Cards      []card
	Cursor     int
	State      int
	question   textinput.Model
	answer     textinput.Model
	inputFocus int
}

func InitialModel() ListModel {
	tiq := textinput.New()
	tiq.Placeholder = "Question"
	tiq.Focus()

	tia := textinput.New()
	tia.Placeholder = "Answer"

	return ListModel{
		question: tiq,
		answer:   tia,
	}
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

	if m.State == StateSelected {
		return updateCard(msg, m)
	}

	if m.State == StateAdd {
		return updateAdd(msg, m)
	}

	return updateList(msg, m)
}

func updateAdd(msg tea.Msg, m ListModel) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			if m.inputFocus == 1 {
				if m.question.Value() != "" && m.answer.Value() != "" {
					c := card{
						Question: m.question.Value(),
						Answer:   m.answer.Value(),
						UUID:     uuid.NewString(),
					}
					m.question.SetValue("")
					m.answer.SetValue("")
					m.Cards = append(m.Cards, c)
				}
				m.State = StateList
				m.question.Focus()
				m.answer.Blur()
				m.inputFocus = 0
				return m, nil
			}
			m.inputFocus++
			m.question.Blur()
			return m, m.answer.Focus()
		}
	}

	if m.inputFocus == 0 {
		m.question, cmd = m.question.Update(msg)
	} else {
		m.answer, cmd = m.answer.Update(msg)
	}
	return m, cmd
}

func updateList(msg tea.Msg, m ListModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			m.State = StateAdd
			return m, nil
		case "d":
			m.Cards = append(m.Cards[:m.Cursor], m.Cards[m.Cursor+1:]...)
			m.Cursor -= 1
			if m.Cursor < 0 {
				m.Cursor = 0
			}
			return m, nil
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Cards)-1 {
				m.Cursor++
			}
		case "enter":
			m.State = StateSelected
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
			m.State = StateList
			return m, nil
		}
	}
	return m, nil
}

func (m ListModel) View() string {
	switch m.State {
	case StateList:
		return listView(m)
	case StateSelected:
		return cardView(m)
	case StateAdd:
		return addView(m)
	default:
		return ""
	}
}

func addView(m ListModel) string {
	s := fmt.Sprintf(
		"Question?\n\n%s",
		m.question.View(),
	) + "\n"

	s = s + fmt.Sprintf(
		"\nAnswer?\n\n%s\n",
		m.answer.View(),
	)

	s += "\ntab: focus next\n"

	return s
}

func listView(m ListModel) string {
	s := "\n"
	for i, card := range m.Cards {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, card.Question)
	}
	s += "\na: new card • e: edit • d: delete • t: train • q: quit\n"
	return s
}

func cardView(m ListModel) string {
	c := m.Cards[m.Cursor]
	s := fmt.Sprintf("\n%s\n", c.Answer)
	s += "\nenter: back\n"
	return s
}

func Write(m ListModel, writer io.Writer) error {
	bb, err := json.Marshal(m.Cards)
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
	m.Cards = cards
	return m, nil
}
