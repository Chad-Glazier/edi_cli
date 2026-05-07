package out

import (
	"fmt"

	"github.com/Chad-Glazier/edi/diag"
	"github.com/Chad-Glazier/edi/diag/dmm"
	"github.com/Chad-Glazier/edi_cli/out/rep"
)

// Prints the report from a single turn to the standard output. The startRow
// and startCol coordinates define the location of the top-left corner of the
// report info; the returned values specify the bottom-right corner.
func PrintReport(
	report diag.Report, startRow, startCol int,
) (
	endRow, endCol int,
) {
	switch report.(type) {
	case dmm.AlphaBetaReport:
		return rep.AlphaBeta(report, startRow, startCol)
	default:
		panic(fmt.Sprintf("Report for type %T not implemented.\n", report))
	}
}
