package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
)

func FireNotifiers(cfg defs.Config, summaries defs.Summaries, scanResults []defs.SiteDefinition) {
	log.Debug("Firing notifiers ...")

	// check to see if ascii has been specified in the config
	if _, ok := cfg.Reports["ascii"]; ok {
		n := NotifierAscii{
			ScanResults: scanResults,
			Summaries:   summaries,
			Cfg:         cfg,
		}
		n.Fire()
	}

	// check to see if mailgun has been specified in the config
	if _, ok := cfg.Reports["mailgun"]; ok {
		n := NotifierMailgun{
			ScanResults: scanResults,
			Summaries:   summaries,
			Cfg:         cfg,
		}
		n.Fire()
	}
}
