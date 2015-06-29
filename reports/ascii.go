package reports

import (
	"github.com/pemcconnell/amald/defs"
)

type ReportAscii struct{}

// Generate creates an HTML string with all the required data
// in place
func (r *ReportAscii) Generate(results defs.Results) (string, error) {
	output := ""
	return output, nil
}
