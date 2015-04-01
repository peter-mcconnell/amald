package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"time"
)

type ReportHTML struct{}

// htmlHead creates the start of the HTML template
func htmlHead() string {
	html := `
	<html>
		<body>
	`
	return html

}

func summary() string {
	html := `
	<h3>Amald summary</h3>
	<i>Below is a list of changes to the scans that Amald has been performing &
	stored</i>
	<h4>since yesterday</h4>
	<strong>new urls</strong>
	<table width="100%" cellpadding="4">
	<thead>
	<tr>
		<th align="left">URL</th>
		<th>Locked Down</th>
	</tr>
	</thead>
	<tbody>
	<tr>
		<td>URL</th>
		<td>YES</th>
	</tr>
	</tbody>
	</table>
	<strong>removed urls</strong>
	<table width="100%" cellpadding="4">
	<thead>
	<tr>
		<th align="left">URL</th>
		<th>Locked Down</th>
	</tr>
	</thead>
	<tbody>
	<tr>
		<td>URL</th>
		<td>YES</th>
	</tr>
	</tbody>
	</table>
	<strong>updated urls</strong>
	<table width="100%" cellpadding="4">
	<thead>
	<tr>
		<th align="left">URL</th>
		<th>Locked Down</th>
	</tr>
	</thead>
	<tbody>
	<tr>
		<td>URL</th>
		<td>YES</th>
	</tr>
	</tbody>
	</table>
	`
	return html
}

// listHead creates the start of the list template
func listHead() string {
	html := `
		<h4>All apps</h4>
		<i>Below is a list of the current state of all apps that Amald scanned 
		on its most recent sweep</i>
		<table width="100%" cellpadding="4">
			<thead>
			<tr>
				<th align="left">URL</th>
				<th>Locked Down</th>
			</tr>
			</thead>
			<tbody>
	`
	return html
}

// listBody creates a series of table rows
func listBody(data []defs.JsonData) string {
	var buffer bytes.Buffer

	// Pick most recent item (active scan)
	urls := data[0].Data
	for _, url := range urls {
		buffer.WriteString("<tr>")
		// url
		buffer.WriteString("<td><a href=\"" + url.Url + "\">" + url.Url + "</a></td>")
		// locked down
		lockeddown := "Yes"
		lockedcolor := "00CC00"
		if !url.IsLockedDown {
			lockeddown = "No"
			lockedcolor = "FF0000"
		}
		buffer.WriteString("<td style=\"background:#" + lockedcolor + "\" align=\"center\"><strong>" + lockeddown + "</strong></td>")
		buffer.WriteString("</tr>")
	}
	return buffer.String()
}

// listFooter simply closes off the list template
func listFooter() string {
	html := `
			</tbody>
		</table>`
	return html
}

func htmlFooter() string {
	html := `
			</body>
		</html>`
	return html
}

// DiffData is designed to scan all of the available data to try and pull out
// some comparitive differences in the data sets.
// data is in order (newest first) which gives us some advantages for marking
// out which data to use for
func diffData(data []defs.JsonData) (map[string][]defs.JsonData, error) {
	var err error
	var summary = make(map[string][]defs.JsonData)
	for _, d := range data {
		log.Info(time.Parse("2010-12-20 21:16:26.371Z", d.Meta["utc_timestamp"]))
	}
	return summary, err
}

// GenerateHTML creates an HTML string with all the required data
// in place
func (r *ReportHTML) Generate(data []defs.JsonData) (string, error) {
	_, err := diffData(data)
	if err != nil {
		log.Errorf("failed to diffData: %s", err)
	}
	html := htmlHead()
	html += summary()
	html += listHead()
	html += listBody(data)
	html += listFooter()
	html += htmlFooter()
	log.Debug("html:\n", html)
	return html, nil
}
