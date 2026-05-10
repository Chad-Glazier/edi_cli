package ui

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Chad-Glazier/edi"
)

// Holds the currently-selected VI title.
var whiteTitle string
var blackTitle string

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list          list.Model
	selectedWhite bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			i := m.list.SelectedItem().(item)
			if !m.selectedWhite {
				whiteTitle = i.title
				m.selectedWhite = true
				setBlackStyles(&m)
			} else {
				blackTitle = i.title
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	v := tea.NewView(docStyle.Render(m.list.View()))
	v.AltScreen = true
	return v
}

// Prompts the user to select two VIs to play a game and returns them. The
// first VI should play White and the second should play Black.
func SelectPlayers() (edi.VI, edi.VI) {
	items := []list.Item{
		item{
			title: "EDI",
			desc:  "Alpha-beta search with the k-mindist heuristic for leaf evaluation and the History Heuristic for move ordering.",
		},
		item{
			title: "Arrow",
			desc:  "Alpha-beta search with the q-mindist heuristic for leaf evaluation and no move ordering.",
		},
		item{
			title: "Random",
			desc:  "Chooses moves at random.",
		},
	}

	d := list.NewDefaultDelegate()
	m := model{list: list.New(items, d, 0, 0)}

	//
	// Set styles for picking white.
	//

	setWhiteStyles(&m)

	//
	// Run the program
	//

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	var whiteVI, blackVI edi.VI

	switch whiteTitle {
	case "EDI":
		whiteVI = &edi.EDI{}
	case "Arrow":
		whiteVI = &edi.Arrow{}
	case "Random":
		whiteVI = &edi.Random{}
	}

	switch blackTitle {
	case "EDI":
		blackVI = &edi.EDI{}
	case "Arrow":
		blackVI = &edi.Arrow{}
	case "Random":
		blackVI = &edi.Random{}
	}

	whiteTitle = ""
	blackTitle = ""

	return whiteVI, blackVI
}

func setWhiteStyles(m *model) {
	m.list.Title = " Select who plays White "

	blue := lipgloss.Color("#007bff")
	cyan := lipgloss.Color("#03befc")
	lightgray := lipgloss.Color("#eeeeee")
	white := lipgloss.Color("#ffffff")

	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.SelectedTitle.
		Foreground(cyan).
		BorderLeftForeground(cyan)
	d.Styles.SelectedDesc = d.Styles.SelectedDesc.
		Foreground(lightgray).
		BorderLeftForeground(cyan)
	m.list.SetDelegate(d)

	m.list.Styles.Title = lipgloss.NewStyle().
		Background(blue).
		Foreground(white)
}

func setBlackStyles(m *model) {
	m.list.Title = " Select who plays Black "

	red := lipgloss.Color("#c60000")
	lightRed := lipgloss.Color("#ff7070")
	lightgray := lipgloss.Color("#eeeeee")
	white := lipgloss.Color("#ffffff")

	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.SelectedTitle.
		Foreground(lightRed).
		BorderLeftForeground(lightRed)
	d.Styles.SelectedDesc = d.Styles.SelectedDesc.
		Foreground(lightgray).
		BorderLeftForeground(lightRed)
	m.list.SetDelegate(d)

	m.list.Styles.Title = lipgloss.NewStyle().
		Background(red).
		Foreground(white)
}
