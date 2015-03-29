package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/loaders"
)

type ReportHTML struct {
	urls []loaders.SiteDefinition
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
func htmlBody(urls []loaders.SiteDefinition) string {
	var buffer bytes.Buffer
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
func (r *ReportHTML) Generate(urls []loaders.SiteDefinition) (string, error) {
	html := htmlHead()
	html += htmlBody(urls)
	html += htmlFooter()
	log.Debug("html:\n", html)
	return html, nil
}
