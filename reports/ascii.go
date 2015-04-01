package reports

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
	"github.com/pemcconnell/amald/defs"
)

type ReportAscii struct{}

// GenerateHTML creates an HTML string with all the required data
// in place
func (r *ReportAscii) Generate(data []defs.JsonData) (string, error) {
	var buffer bytes.Buffer

	table := tablewriter.NewWriter(&buffer)
	table.SetHeader([]string{"URL", "LockedDown"})

	// Pick most recent item (active scan)
	urls := data[0].Data
	for _, url := range urls {
		//m := make()
		lockeddown := "Yes"
		if !url.IsLockedDown {
			lockeddown = "No"
		}
		table.Append([]string{url.Url, lockeddown})
	}
	table.Render()

	ascii := buffer.String()
	return ascii, nil
}
