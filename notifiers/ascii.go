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
	TestMode    bool
}

// Fire the Ascii report
func (n *NotifierAscii) Fire() error {
	log.Debug("Firing Ascii Notifier")
	r := &reports.Report{
		AnsiColorEnabled: true,
		Cfg:              n.Cfg,
		ScanResults:      n.ScanResults,
	}
	if o, err := r.GenerateAscii(n.Summaries); err != nil {
		log.Errorf("Failed to generate ascii report: %s", err)
		return err
	} else {
		log.Infof("\n%s", o)
	}
	return nil
}
