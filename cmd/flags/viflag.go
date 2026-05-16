package flags

import (
	"fmt"

	"github.com/Chad-Glazier/edi"
)

const VI_USAGE = "a VI player: edi, arrow, or random"

type VI struct {
	vi edi.VI
}

func (v *VI) VI() edi.VI {
	return v.vi
}

//
// Satisfy the flag interface for Cobra.
//
// https://pkg.go.dev/github.com/spf13/pflag#Value
//

func (v *VI) String() string {
	if v.vi == nil {
		return ""
	}
	return v.vi.Id()
}

func (v *VI) Set(s string) error {
	switch s {
	case "edi":
		v.vi = edi.NewEDI()
	case "arrow":
		v.vi = edi.NewArrow()
	case "random":
		v.vi = edi.NewRandom()
	default:
		return fmt.Errorf(`must be one of "edi", "arrow", or "random"`)
	}
	return nil
}

func (v *VI) Type() string {
	return "VI"
}
