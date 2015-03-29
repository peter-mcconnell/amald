package notifiers

import (
	"github.com/pemcconnell/amald/loaders"
)

type NotifierLoader interface {
	Send(map[string]map[string]string, string)
}

func FireNotifiers(urls []loaders.SiteDefinition,
	activeloaders map[string]map[string]string) error {

	// check to see if ascii has been specified in the config
	if _, ok := activeloaders["ascii"]; ok {
		n := &NotifierAscii{urls: urls}
		n.Send()
	}

	// check to see if mailgun has been specified in the config
	if _, ok := activeloaders["mailgun"]; ok {
		n := &NotifierMailgun{urls: urls}
		n.Send(activeloaders["mailgun"])
	}

	return nil
}
