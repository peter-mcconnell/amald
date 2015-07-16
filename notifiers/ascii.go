package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"github.com/pemcconnell/amald/reports"
)

type NotifierAscii struct {
	Summaries defs.Summaries
	Cfg       defs.Config
}

// Send the message via mailgun
func (n *NotifierAscii) Fire() {
	log.Debug("Firing Ascii Notifier")
	r := &reports.Report{}
	r.SetCfg(n.Cfg)
	o, err := r.GenerateAscii(n.Summaries)
	if err != nil {
		log.Fatalf("failed to generate ascii report: %s", err)
	}
	log.Infof("\n%s", o)
}
