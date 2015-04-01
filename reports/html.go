package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
)

type ReportHTML struct {
	urls []defs.SiteDefinition
}

// htmlHead creates the start of the HTML template
func htmlHead() string {
	html := `
	<html>
	<body>
		<table width="100%" cellpadding="4">
			<thead>
				<th align="left">URL</th>
				<th>Locked Down</th>
			</thead>
			<tbody>
	`
	return html
}

// htmlBody creates a series of table rows
func htmlBody(data []defs.JsonData) string {
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

// htmlFooter simply closes off the html message
func htmlFooter() string {
	html := `
			</tbody>
		</body>
	</html>
	`
	return html
}

// GenerateHTML creates an HTML string with all the required data
// in place
func (r *ReportHTML) Generate(data []defs.JsonData) (string, error) {
	html := htmlHead()
	html += htmlBody(data)
	html += htmlFooter()
	log.Debug("html:\n", html)
	return html, nil
}
