package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		// our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// a map which indicates which choices are selected. we're using
		// the map like a mathematical set. the keys refer to the indexes
		// of the `choices` slice, above
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// just return `nil`, which means 'no I/O right now, please'
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// is it a key press?
	case tea.KeyMsg:

		// cool, what was the actual key pressed?
		switch msg.String() {

		// these keys should exit the program
		case "ctrl+c", "q":
			return m, tea.Quit

		// the "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// the "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// the "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// return the updated model to the Bubble Tea runtime for processing
	// note that we're not returning a command
	return m, nil
}

func (m model) View() string {
	// the header
	s := "What should we but at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// the footer
	s += "\nPress q to quit.\n"

	// send the ui for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Coop, there's been an error: %v", err)
		os.Exit(1)
	}
}
