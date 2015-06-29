// TODO: create a proper template for these. The strings are nasty
package reports

import (
	"github.com/pemcconnell/amald/defs"
)

type ReportHTML struct {
	Templatepath string
}

// Generate creates an HTML string with all the required data
// in place
func (r *ReportHTML) Generate(results defs.Results) (string, error) {
	output := ""
	return output, nil
}
