package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"github.com/pemcconnell/amald/reports"
)

type NotifierAscii struct {
	ScanResults []defs.SiteDefinition
	Summaries   defs.Summaries
	Cfg         defs.Config
}

// Fire the Ascii report
func (n *NotifierAscii) Fire() {
	log.Debug("Firing Ascii Notifier")
	r := &reports.Report{
		AnsiColorEnabled: true,
		Cfg:              n.Cfg,
		ScanResults:      n.ScanResults,
	}
	o, err := r.GenerateAscii(n.Summaries)
	if err != nil {
		log.Fatalf("failed to generate ascii report: %s", err)
	}
	log.Infof("\n%s", o)
}
