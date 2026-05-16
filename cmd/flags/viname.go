package flags

import (
	"fmt"

	"github.com/Chad-Glazier/edi"
)

type VIName string

const (
	EDI    VIName = "edi"
	ARROW  VIName = "arrow"
	RANDOM VIName = "random"
)

//
// Satisfy the flag interface for Cobra
//
// https://pkg.go.dev/github.com/spf13/pflag#Value
//

func (v *VIName) String() string {
	return string(*v)
}

func (v *VIName) Set(s string) error {
	switch s {
	case "edi", "arrow", "random":
		*v = VIName(s)
		return nil
	default:
		return fmt.Errorf(`must be one of "edi", "arrow", or "random"`)
	}
}

func (v *VIName) Type() string {
	return "VI"
}

//
// Helper function to instantiate a VI from its name.
//

func CreateVI(v VIName) edi.VI {
	switch v {
	case EDI:
		return edi.NewEDI()
	case ARROW:
		return edi.NewArrow()
	case RANDOM:
		return edi.NewRandom()
	default:
		panic(fmt.Sprintf("unmatched VI name \"%s\"", v))
	}
}
