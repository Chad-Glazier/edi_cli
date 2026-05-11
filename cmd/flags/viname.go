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

func CreateVI(v VIName) edi.VI {
	switch v {
	case EDI:
		return &edi.EDI{}
	case ARROW:
		return &edi.Arrow{}
	case RANDOM:
		return &edi.Random{}
	default:
		panic(fmt.Sprintf("unmatched VI name \"%s\"", v))
	}
}
