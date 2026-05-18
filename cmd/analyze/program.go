package analyze

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
	winner    *state.PlayerColor

	viSelector   ui.VISelector
	board        ui.BoardModel
	timeSelector ui.TimeSelector
}

func NewGameModel(vi flags.VI, turnTimer time.Duration) gameModel {
	m := gameModel{
		turnTimer:    turnTimer,
		viSelector:   ui.NewVISelector(ui.NEUTRAL),
		timeSelector: ui.NewTimeSelector(),
		board:        ui.NewBoardModel(),
	}

	if vi.New != nil {
		m.white = vi.New()
		m.black = vi.New()
	}

	return m
}

//
// Bubbletea methods.
//

func (m gameModel) Init() tea.Cmd {
	return tea.Batch(
		m.viSelector.Init(),
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
		m.viSelector, _ = m.viSelector.Update(msg)
		if m.viSelector.NewVI != nil {
			m.white = m.viSelector.NewVI()
			m.black = m.viSelector.NewVI()			
		}
	default:
		switch msg := msg.(type) {
		case ui.SetBoardMsg:
			m.board, _ = m.board.Update(msg)
			return m, awaitGameUpdate(&m)
		case GameOverMsg:
			m.winner = &msg.winner
			return m, tea.Quit
		}
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viSelector, _ = m.viSelector.Update(msg)
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
		return m.viSelector.View()
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
