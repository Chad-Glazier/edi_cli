package ui

import (
	"time"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

//
// Defining the list items.
//

type timeItem struct {
	title    string
	desc     string
	duration time.Duration
}

func (i timeItem) Title() string       { return i.title }
func (i timeItem) Description() string { return i.desc }
func (i timeItem) FilterValue() string { return i.title }

var timeOptions = []list.Item{
	timeItem{
		title: "1s",
		desc:  "1 second per turn (\u2264 92 seconds per game).",
		duration:   time.Second * 1,
	},
	timeItem{
		title: "5s",
		desc:  "5 seconds per turn (\u2264 8 minutes per game). Appropriate for play against humans.",
		duration:   time.Second * 5,
	},
	timeItem{
		title: "30s",
		desc:  "30 seconds per turn (\u2264 46 minutes per game).",
		duration:   time.Second * 30,
	},
	timeItem{
		title: "1m",
		desc: "1 minute per turn (\u2264 92 minutes per game).",
		duration: time.Minute,
	},
	timeItem{
		title: "15m",
		desc: "15 minutes per turn (\u2264 23 hours per game). Similar to timers that would be used in fast computer chess tournaments.",
		duration: time.Minute * 15,
	},
	timeItem{
		title: "1hr",
		desc: "1 hour per turn (\u2264 4 days per game). Similar to high-level computer chess tournaments.",
		duration: time.Hour,
	},
}

//
// Defining the model state.
//

type TimeSelector struct {
	list  list.Model
	TurnTimer    time.Duration
	Style lipgloss.Style
}

func NewTimeSelector() TimeSelector {
	m := TimeSelector{}
	m.list = list.New(timeOptions, list.NewDefaultDelegate(), 0, 0)
	return m
}

//
// Bubbletea methods.
//

func (m TimeSelector) Init() tea.Cmd {
	return nil
}

func (m TimeSelector) Update(msg tea.Msg) (TimeSelector, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			item := m.list.SelectedItem().(timeItem)
			m.TurnTimer = item.duration
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := m.Style.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m TimeSelector) View() tea.View {
	v := tea.NewView(m.Style.Render(m.list.View()))
	v.AltScreen = true
	return v
}
