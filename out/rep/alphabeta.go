package rep

import (
	"fmt"
	"time"

	"github.com/Chad-Glazier/edi/diag"
	"github.com/Chad-Glazier/edi_cli/ansi"
)

// Prints the report from a single turn to the standard output. The startRow
// and startCol coordinates define the location of the top-left corner of the
// report info; the returned values specify the bottom-right corner.
func AlphaBeta(
	report diag.Report,
	startRow, startCol int,
) (
	endRow, endCol int,
) {
	type depthStats struct {
		depth    int
		duration time.Duration
		leaves   uint64
	}

	const (
		BAR_WIDTH = 28
	)

	set := func(r, c int) {
		ansi.SetCursor(startRow+r, startCol+c)
	}

	maxDepth := report.GreatestDepth()

	// ------------------------------------------------------------
	// Gather stats
	// ------------------------------------------------------------

	stats := make([]depthStats, 0, maxDepth)

	var (
		maxLeaves   uint64
		maxDuration time.Duration
		totalLeaves uint64
		maxCutoff   uint64
	)

	for depth := 1; depth <= maxDepth; depth++ {
		s := depthStats{
			depth:    depth,
			duration: report.SearchDuration(depth),
			leaves:   report.Leaves(depth),
		}

		stats = append(stats, s)

		totalLeaves += s.leaves

		if s.leaves > maxLeaves {
			maxLeaves = s.leaves
		}

		if s.duration > maxDuration {
			maxDuration = s.duration
		}

		for cutoffDepth := 1; cutoffDepth <= depth; cutoffDepth++ {
			v := report.Cutoffs(depth, cutoffDepth)

			if v > maxCutoff {
				maxCutoff = v
			}
		}
	}

	// ------------------------------------------------------------
	// Rendering
	// ------------------------------------------------------------

	row := 0
	maxWidth := 0

	write := func(r, c int, s string) {
		set(r, c)
		fmt.Print(s)

		if c+len([]rune(s)) > maxWidth {
			maxWidth = c + len([]rune(s))
		}
	}

	// ------------------------------------------------------------
	// Header
	// ------------------------------------------------------------

	write(
		row,
		0,
		ansi.Bold(ansi.FgBrightWhite(" Alpha-Beta Search Report ")),
	)
	row += 2

	write(
		row,
		0,
		ansi.Bold(" Recommended Move: ")+
			ansi.FgBrightCyan(fmt.Sprintf("%v", report.Move())),
	)
	row += 2

	// ------------------------------------------------------------
	// Iterative deepening
	// ------------------------------------------------------------

	write(
		row,
		0,
		ansi.Bold(ansi.FgBrightWhite(" Iterative Deepening ")),
	)
	row += 2

	for _, s := range stats {
		durationWidth := 0
		leafWidth := 0

		if maxDuration > 0 {
			durationWidth =
				int(float64(s.duration) /
					float64(maxDuration) *
					BAR_WIDTH)
		}

		if maxLeaves > 0 {
			leafWidth =
				int(float64(s.leaves) /
					float64(maxLeaves) *
					BAR_WIDTH)
		}

		durationBar :=
			ansi.FgBrightBlue(ansi.Repeat(durationWidth, "█")) +
				ansi.Dim(ansi.Repeat(BAR_WIDTH-durationWidth, "░"))

		leafBar :=
			ansi.FgBrightGreen(ansi.Repeat(leafWidth, "█")) +
				ansi.Dim(ansi.Repeat(BAR_WIDTH-leafWidth, "░"))

		write(
			row,
			0,
			fmt.Sprintf(
				" Depth %-2d  %s  %8s",
				s.depth,
				durationBar,
				ansi.FgBrightBlue(s.duration.String()),
			),
		)

		row++

		write(
			row,
			0,
			fmt.Sprintf(
				"           %s  %s leaves",
				leafBar,
				ansi.FgBrightGreen(fmt.Sprintf("%d", s.leaves)),
			),
		)

		row += 2
	}

	// ------------------------------------------------------------
	// Cutoff heatmap
	// ------------------------------------------------------------

	write(
		row,
		0,
		ansi.Bold(ansi.FgBrightWhite(" Cutoff Distribution ")),
	)
	row += 2

	header := " SD\\CD "

	for cutoffDepth := 1; cutoffDepth <= maxDepth; cutoffDepth++ {
		header += fmt.Sprintf(" %2d ", cutoffDepth)
	}

	write(row, 0, header)
	row++

	for searchDepth := 1; searchDepth <= maxDepth; searchDepth++ {
		line := fmt.Sprintf("  %2d   ", searchDepth)

		for cutoffDepth := 1; cutoffDepth <= maxDepth; cutoffDepth++ {
			if cutoffDepth > searchDepth {
				line += "    "
				continue
			}

			v := report.Cutoffs(searchDepth, cutoffDepth)

			if v == 0 {
				line += ansi.Dim(" ·  ")
				continue
			}

			ratio := float64(v) / float64(maxCutoff)

			cell := ansi.FgBrightBlack("░")

			switch {
			case ratio > 0.75:
				cell = ansi.FgBrightRed("█")
			case ratio > 0.50:
				cell = ansi.FgBrightYellow("▓")
			case ratio > 0.25:
				cell = ansi.FgBrightGreen("▒")
			default:
				cell = ansi.FgBrightBlue("░")
			}

			line += fmt.Sprintf(" %s ", cell)
		}

		write(row, 0, line)
		row++
	}

	row++

	// ------------------------------------------------------------
	// Footer
	// ------------------------------------------------------------

	write(
		row,
		0,
		ansi.Bold(" Completed Depth: ")+
			ansi.FgBrightWhite(fmt.Sprintf("%d", maxDepth)),
	)
	row++

	write(
		row,
		0,
		ansi.Bold(" Total Leaves:    ")+
			ansi.FgBrightGreen(fmt.Sprintf("%d", totalLeaves)),
	)

	// ------------------------------------------------------------
	// Final bounds
	// ------------------------------------------------------------

	endRow = startRow + row
	endCol = startCol + maxWidth
	return
}
