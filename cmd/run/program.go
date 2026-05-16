package run

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Chad-Glazier/edi"
	"github.com/Chad-Glazier/edi/state"
	"github.com/Chad-Glazier/edi_cli/cmd/flags"
	"github.com/Chad-Glazier/edi_cli/sim"
	"github.com/Chad-Glazier/edi_cli/ui"
)

//
// Defining the model state.
//

type gameModel struct {
	height int
	width  int

	white     edi.VI
	black     edi.VI
	turnTimer time.Duration
	game      <-chan state.Board
	winner        *state.PlayerColor

	whiteSelector ui.VISelector
	blackSelector ui.VISelector
	board         ui.BoardModel
	timeSelector ui.TimeSelector
}

func NewGameModel(white, black flags.VI, turnTimer time.Duration) gameModel {
	return gameModel{
		white:         white.VI(),
		black:         black.VI(),
		turnTimer:     turnTimer,
		whiteSelector: ui.NewVISelector(ui.WHITE),
		blackSelector: ui.NewVISelector(ui.BLACK),
		timeSelector: ui.NewTimeSelector(),
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
		m.timeSelector.Init(),
		m.board.Init(),
	)
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch {
	case m.turnTimer == 0:
		m.timeSelector, _ = m.timeSelector.Update(msg)
		m.turnTimer = m.timeSelector.TurnTimer
	case m.white == nil:
		m.whiteSelector, _ = m.whiteSelector.Update(msg)
		m.white = m.whiteSelector.VI
	case m.black == nil:
		m.blackSelector, _ = m.blackSelector.Update(msg)
		m.black = m.blackSelector.VI
	default:
		switch msg := msg.(type) {
		case ui.SetBoardMsg:
			m.board, _ = m.board.Update(msg)
			return m, awaitGameUpdate(&m)
		case GameOverMsg:
			m.winner = &msg.winner
			return m, nil
		}
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.whiteSelector, _ = m.whiteSelector.Update(msg)
		m.blackSelector, _ = m.blackSelector.Update(msg)
		m.timeSelector, _ = m.timeSelector.Update(msg)
		m.board, _ = m.board.Update(msg)
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	if m.white != nil && m.black != nil && m.game == nil {
		m.game = sim.Game(m.white, m.black, m.turnTimer)
		return m, awaitGameUpdate(&m)
	}

	return m, nil
}

func (m gameModel) View() tea.View {
	switch {
	case m.width == 0:
		return tea.NewView(ui.FgBrightBlack("Loading..."))
	case m.turnTimer == 0:
		return m.timeSelector.View()
	case m.white == nil:
		return m.whiteSelector.View()
	case m.black == nil:
		return m.blackSelector.View()
	case m.winner == nil:
		v := tea.NewView(lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			m.board.View(),
		))
		v.AltScreen = true
		return v
	default:
		winner := ""
		switch *m.winner {
		case state.WHITE:
			winner = ui.FgBrightCyan(m.white.Id())
		case state.BLACK:
			winner = ui.FgBrightRed(m.black.Id())
		}
		return tea.NewView(lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				m.board.View(),
				fmt.Sprintf("%s Wins!", winner),
			),
		))
	}
}

//
// Custom commands.
//

func awaitGameUpdate(m *gameModel) tea.Cmd {
	return func() tea.Msg {
		updatedGameState, ok := <-m.game
		if !ok {
			switch m.board.State.Player {
			case state.WHITE:
				return GameOverMsg{winner: state.BLACK}
			case state.BLACK:
				return GameOverMsg{winner: state.WHITE}
			}
		}
		return ui.SetBoardMsg(updatedGameState)
	}
}

//
// Custom messages.
//

type GameOverMsg struct {
	winner state.PlayerColor
}

//
// Styles.
//


