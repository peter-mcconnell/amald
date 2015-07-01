package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"github.com/pemcconnell/amald/reports"
)

type NotifierAscii struct {
	results defs.Results
}

// Send the message via mailgun
func (n *NotifierAscii) Fire() {
	log.Debug("Firing Ascii Notifier")
	r := &reports.ReportAscii{}
	o, err := r.Generate(n.results)
	if err != nil {
		log.Fatalf("failed to generate ascii report: %s", err)
	}
	log.Infof("\n%s", o)
}
