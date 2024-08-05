package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
)

// bubble tea programs composed of four things:

// a model
type model struct {
	info     []string
	choices  []string
	cursor   int
	selected map[int]struct{}
	view     int
}

// describes application state

func main() {
	p := tea.NewProgram(init_model())
	if _, err := p.Run(); err != nil {
		fmt.Printf("aye bro theres been an error, %v", err)
		os.Exit(1)
	}

}

// returns initial state for program to run
func init_model() model {
	return model{choices: []string{"About", "Education", "Projects"}, selected: make(map[int]struct{}), view: 0}
}

func default_view(m *model) {
	m.choices = []string{"About", "Education", "Projects"}
	m.selected = make(map[int]struct{})
	m.info = []string{}
	m.view = 0
}

func view_one(m *model) {
	m.choices = []string{"Return to menu"}
	m.selected = make(map[int]struct{})
	m.cursor = 0
	m.view = 1
	m.info = []string{
		wordwrap.String("I'm Jossaya, a junior at FAU High School majoring in Data Science and Analytics.", 80),
		wordwrap.String("My main interests lie in software engineering, AI, computer vision, cybersecurity, and data analytics.", 80),
		wordwrap.String("I'm currently learning Go, React Native, and C++ among other things.", 80),
		wordwrap.String("I've learned Python, JavaScript/TypeScript, React and am familiar with C.", 80),
		wordwrap.String("Outside of computer science, I enjoy playing video games, am an LA at FAU, and would like to learn the piano.", 80),
	}
	m.choices = []string{
		"Back to home",
	}
}
func view_two(m *model) {
	m.choices = []string{"Return to menu"}
	m.selected = make(map[int]struct{})
	m.cursor = 0
	m.view = 1
	m.info = []string{
		wordwrap.String("Florida Atlantic University High School -- 2023-present", 80),
		wordwrap.String("--Junior at FAU High School", 80),
		wordwrap.String("--Currently working on progressing into undergraduate research on AI in Healthcare with Dr. Behnaz Ghoraani", 80),
		wordwrap.String("--Completed coursework like Programming 2 in C++, Calculus 2, and General Physics 1", 80),
	}
	m.choices = []string{
		"Back to home",
	}
}

// intializes program
func (m model) Init() tea.Cmd {
	return nil
}

// update method
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { // should probably leave this alone for now lol
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "w":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "s":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ", "\n":
			if m.view == 0 {
				if m.cursor == 0 {
					view_one(&m)
				} else if m.cursor == 1 {
					view_two(&m)
				} else if m.cursor == 2 {
					m.view = 3
				}
			} else if m.view == 1 {
				default_view(&m)
			} else if m.view == 2 {
				default_view(&m)
			}

		}
	}

	// returns an updated model
	return m, nil

}

func (m model) View() string { // probably gna end up messing with this later
	s := " "
	if m.view == 0 {
		s = "Welcome to my portfolio! I'm Jossaya Camille.\n"
	} else if m.view == 1 {
		s = "About me\n"
	} else if m.view == 2 {
		s = "Education\n"
	}
	for _, detail := range m.info {
		s += fmt.Sprintf("%s \n", detail)
	}
	s += "\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = ">>"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress Q to quit.\n"
	return s

}
