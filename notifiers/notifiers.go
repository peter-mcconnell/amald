package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
)

// FireNotifiers checks the settings in the Config for known report types and
// calls them if they have been defined
func FireNotifiers(cfg defs.Config, summaries defs.Summaries,
	scanResults []defs.SiteDefinition, testMode bool) map[string]bool {
	log.Debug("Firing notifiers ...")

	fired := make(map[string]bool, 0)

	// check to see if ascii has been specified in the config
	if _, ok := cfg.Reports["ascii"]; ok {
		n := NotifierAscii{
			ScanResults: scanResults,
			Summaries:   summaries,
			Cfg:         cfg,
			TestMode:    testMode,
		}
		if err := n.Fire(); err != nil {
			log.Errorf("Failed to fire: %s", err)
		} else {
			fired["ascii"] = true
		}
	}

	// check to see if mailgun has been specified in the config
	if _, ok := cfg.Reports["mailgun"]; ok {
		n := NotifierMailgun{
			ScanResults: scanResults,
			Summaries:   summaries,
			Cfg:         cfg,
			TestMode:    testMode,
		}
		if err := n.Fire(); err != nil {
			log.Errorf("Failed to fire: %s", err)
		} else {
			fired["mailgun"] = true
		}
	}

	return fired
}
