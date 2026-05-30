package ui

import (
	"fmt"

	"github.com/Chad-Glazier/edi"
)

// Returns a formatted string representation of a VI's analytics. If the VI
// yields unsupported analytics, an error is returned.
func AnalyticsView(vi edi.VI) (string, error) {
	switch a := vi.GetAnalytics().(type) {
	case edi.EDIAnalytics:
		return fmt.Sprintf("EDI analytics: %v", a), nil
	case edi.ArrowAnalytics:
		return fmt.Sprintf("Arrow analytics: %v", a), nil
	case edi.SparrowAnalytics:
		return fmt.Sprintf("Sparrow analytics: %v", a), nil
	default:
		return "", fmt.Errorf("analyticsView: unhandled analytics type %T", a)
	}
}

//
// Single search alpha-beta analytics.
//

//
// Aggregate alpha-beta analytics.
//
