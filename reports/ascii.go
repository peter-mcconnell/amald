package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
	"github.com/pemcconnell/amald/defs"
)

type ReportAscii struct{}

var (
	buffer bytes.Buffer
)

// Generate creates an HTML string with all the required data
// in place
func (r *ReportAscii) Generate(summaries defs.Summaries) (string, error) {
	log.Debug(summaries)
	output := "\n[ SUMMARIES ]\n"
	for title, summary := range summaries {
		table := tablewriter.NewWriter(&buffer)
		for ks, k := range defs.StateKeys {
			table.SetHeader([]string{ks})
			table.SetHeader([]string{"URL", "LockedDown", "Status Code", "Status"})
			for _, sd := range summary[k] {
				table.Append([]string{sd.Url, "x", "y", "z"})
			}
		}
		table.Render()
		output += "\n " + title + "\n"
		output += buffer.String()
	}
	return output, nil
}
