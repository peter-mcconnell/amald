package reports

import (
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/mgutz/ansi"
	"github.com/olekukonko/tablewriter"
	"github.com/pemcconnell/amald/defs"
	"sort"
)

var (
	buffer bytes.Buffer
)

// Generate creates an HTML string with all the required data
// in place
func (r *Report) GenerateAscii(summaries defs.Summaries) (string, error) {
	log.Debug(summaries)
	colorreset := ansi.ColorCode("reset")
	output := "\n[ SUMMARIES ]\n"
	if len(r.Cfg.SummaryIntervals) != 0 {
		// ensure we're looping summaryintervals in the right order
		var keys []int
		for k := range r.Cfg.SummaryIntervals {
			keys = append(keys, k)
		}
		// ensure we're looping states in the right order
		var statekeys []string
		for k := range defs.StateKeys {
			statekeys = append(statekeys, k)
		}
		sort.Strings(statekeys)
		// now we can loop through the data, with an expected order
		for _, k := range keys {
			title := r.Cfg.SummaryIntervals[k].Title
			color := ansi.ColorCode(r.Cfg.SummaryIntervals[k].Ansii)
			output += "\n " + title + "\n"
			for _, s := range statekeys {
				buffer.Reset()
				output += " ~ " + s + " [since " + title + "]\n"
				table := tablewriter.NewWriter(&buffer)
				table.SetHeader([]string{"URL", "LockedDown", "Status Code"})
				for _, sd := range summaries[k][defs.StateKeys[s]] {
					table.Append([]string{sd.Url, fmt.Sprintf("%t", sd.IsLockedDown), fmt.Sprintf("%d", sd.HttpStatusCode)})
				}
				table.Render()
				if r.AsciiColorEnabled {
					output += color + buffer.String() + colorreset
				} else {
					output += buffer.String()
				}
			}
		}
	} else {
		log.Warn("There have been no summaryintervals defined")
	}
	return output, nil
}
