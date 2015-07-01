package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
)

type NotifierLoader interface {
	Fire(map[string]map[string]string, string)
}

func FireNotifiers(cfg defs.Config, results defs.Results) {
	log.Debug("Firing notifiers ...")
	// check to see if ascii has been specified in the config
	if _, ok := cfg.Reports["ascii"]; ok {
		n := &NotifierAscii{
			results: results,
		}
		n.Fire()
	}

	// check to see if mailgun has been specified in the config
	if _, ok := cfg.Reports["mailgun"]; ok {
		n := &NotifierMailgun{
			results:      results,
			templatepath: cfg.Reports["templates"]["path"],
		}
		n.Fire(cfg.Loaders["mailgun"])
	}
}
