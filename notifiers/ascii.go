package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"github.com/pemcconnell/amald/reports"
)

type NotifierAscii struct {
	data []defs.JsonData
}

// Send the message via mailgun
func (n *NotifierAscii) Send() {
	r := &reports.ReportAscii{}
	o, err := r.Generate(n.data)
	if err != nil {
		log.Fatalf("failed to generate ascii report: %s", err)
	}
	log.Infof("\n%s", o)
}
