package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/loaders"
	"github.com/pemcconnell/amald/reports"
)

type NotifierAscii struct {
	urls []loaders.SiteDefinition
}

// Send the message via mailgun
func (n *NotifierAscii) Send() {
	r := reports.ReportAscii{}
	o, err := r.Generate(n.urls)
	if err != nil {
		log.Fatalf("failed to generate ascii report: %s", err)
	}
	log.Infof("\n%s", o)
}
