package run

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/Chad-Glazier/edi"
	"github.com/Chad-Glazier/edi_cli/ui"
)

//
// Defining the model state.
//

type gameModel struct {
	height int
	width  int

	white         edi.VI
	black         edi.VI
	whiteSelector ui.VISelector
	blackSelector ui.VISelector
	board         ui.BoardModel
	turnTimer     *time.Duration
}

func NewGameModel(white, black edi.VI, turnTimer *time.Duration) gameModel {
	return gameModel{
		white:         white,
		black:         black,
		turnTimer:     turnTimer,
		whiteSelector: ui.NewVISelector(ui.WHITE),
		blackSelector: ui.NewVISelector(ui.BLACK),
		board:         ui.NewBoardModel(),
	}
}

//
// Bubbletea methods.
//

func (m gameModel) Init() tea.Cmd {
	return tea.Batch(
		m.blackSelector.Init(),
		m.whiteSelector.Init(),
		m.board.Init(),
	)
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch {
	case m.white == nil:
		m.whiteSelector, _ = m.whiteSelector.Update(msg)
		m.white = m.whiteSelector.VI
	case m.black == nil:
		m.blackSelector, _ = m.blackSelector.Update(msg)
		m.black = m.blackSelector.VI
	default:
		m.board, _ = m.board.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.whiteSelector, _ = m.whiteSelector.Update(msg)
		m.blackSelector, _ = m.blackSelector.Update(msg)
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m gameModel) View() tea.View {
	switch {
	case m.width == 0:
		return tea.NewView(ui.FgBrightBlack("Loading..."))
	case m.white == nil:
		return m.whiteSelector.View()
	case m.black == nil:
		return m.blackSelector.View()
	default:
		return m.board.View()
	}
}
