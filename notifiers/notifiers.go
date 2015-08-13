package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
)

// FireNotifiers checks the settings in the Config for known report types and
// calls them if they have been defined
func FireNotifiers(cfg defs.Config, summaries defs.Summaries,
	scanResults []defs.SiteDefinition) []string {
	log.Debug("Firing notifiers ...")

	fired := make([]string, 0)

	// check to see if ascii has been specified in the config
	if _, ok := cfg.Reports["ascii"]; ok {
		n := NotifierAscii{
			ScanResults: scanResults,
			Summaries:   summaries,
			Cfg:         cfg,
		}
		if err := n.Fire(); err != nil {
			log.Errorf("Failed to fire: %s", err)
		} else {
			fired = append(fired, "ascii")
		}
	}

	// check to see if mailgun has been specified in the config
	if _, ok := cfg.Reports["mailgun"]; ok {
		n := NotifierMailgun{
			ScanResults: scanResults,
			Summaries:   summaries,
			Cfg:         cfg,
		}
		if err := n.Fire(); err != nil {
			log.Errorf("Failed to fire: %s", err)
		} else {
			fired = append(fired, "mailgun")
		}
	}

	return fired
}
