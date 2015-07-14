package reports

import (
	"github.com/pemcconnell/amald/defs"
)

type ReportHTML struct {
	Templatepath string
}

// Generate creates an HTML string with all the required data
// in place
func (r *ReportHTML) Generate(summaries defs.Summaries) (string, error) {
	output := ""
	return output, nil
}
