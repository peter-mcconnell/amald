package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/mgutz/ansi"
	"github.com/olekukonko/tablewriter"
	"github.com/pemcconnell/amald/defs"
)

type ReportAscii struct{}

var (
	buffer bytes.Buffer
)

func createSummaryForType(data DataDiff, keyprefix string) map[string]string {

	ret := make(map[string]string)

	colorreset := ansi.ColorCode("reset")

	for i := 0; i < 3; i++ {
		buffer.Reset()
		var (
			dat        map[string]bool
			keypostfix string
			color      string
		)
		switch i {
		default:
			dat = data.AddedApps
			keypostfix = "added"
			color = ansi.ColorCode("green+h:black")
		case 1:
			dat = data.UpdatedApps
			keypostfix = "updated"
			color = ansi.ColorCode("blue+h:black")
		case 2:
			dat = data.RemovedApps
			keypostfix = "removed"
			color = ansi.ColorCode("red+h:black")
		}
		if len(dat) != 0 {
			table := tablewriter.NewWriter(&buffer)
			table.SetHeader([]string{"URL", "LockedDown"})

			for url, ld := range dat {
				lockeddown := "Yes"
				if !ld {
					lockeddown = "No"
				}
				table.Append([]string{url, lockeddown})
			}
			table.Render()
			ret[keyprefix+"."+keypostfix] = color + buffer.String() + colorreset
		}
	}
	buffer.Reset()
	return ret
}

// GenerateHTML creates an HTML string with all the required data
// in place
func (r *ReportAscii) Generate(data []defs.JsonData) (string, error) {

	diff, err := FindDiffs(data)
	if err != nil {
		log.Errorf("failed to diffData: %s", err)
	}
	type TemplateData struct {
		Summary map[string]DataDiff
		List    defs.JsonData
	}
	tmpldata := TemplateData{}
	tmpldata.Summary = diff
	tmpldata.List = data[0]

	// start to create the ascii output
	var output string

	// summary
	summaries := make(map[string]map[string]string)

	summaries["yesterday"] = createSummaryForType(tmpldata.Summary["yesterday"], "yesterday")
	summaries["lastweek"] = createSummaryForType(tmpldata.Summary["lastweek"], "lastweek")
	summaries["thirtydays"] = createSummaryForType(tmpldata.Summary["thirtydays"], "thirtydays")

	summaryout := ""
	for _, dat := range summaries {
		if len(dat) != 0 {
			for k, out := range dat {
				summaryout += "\n ~ " + k + "\n"
				summaryout += out
			}
		}
	}
	if summaryout != "" {
		output += "\n[ SUMMARIES ]\n" + summaryout
	}

	// latest scan
	output += "\n[ LATEST SCAN ]\n"

	buffer.Reset()
	table := tablewriter.NewWriter(&buffer)
	table.SetHeader([]string{"URL", "LockedDown"})

	// Pick most recent item (active scan)
	for _, url := range tmpldata.List.Data {
		//m := make()
		lockeddown := "Yes"
		if !url.IsLockedDown {
			lockeddown = "No"
		}
		table.Append([]string{url.Url, lockeddown})
	}
	table.Render()

	output += buffer.String()
	return output, nil
}
