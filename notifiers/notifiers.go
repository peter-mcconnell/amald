package notifiers

import (
	"github.com/pemcconnell/amald/defs"
)

type NotifierLoader interface {
	Send(map[string]map[string]string, string)
}

func FireNotifiers(cfg defs.Config, results defs.Results) {

	// check to see if ascii has been specified in the config
	if _, ok := cfg.Loaders["ascii"]; ok {
		n := &NotifierAscii{
			results: results,
		}
		n.Send()
	}

	// check to see if mailgun has been specified in the config
	if _, ok := cfg.Loaders["mailgun"]; ok {
		n := &NotifierMailgun{
			results:      results,
			templatepath: cfg.Reports["templates"]["path"],
		}
		n.Send(cfg.Loaders["mailgun"])
	}
}
