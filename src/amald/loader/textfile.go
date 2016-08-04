package loader

import "amald/config"

/*
Textfile implements the Loader interface and is designed to
load urls from a textfile
*/
type Textfile struct{}

func (g Textfile) load(loader config.Loader) ([]config.Target, error) {
	var targets []config.Target

	return targets, nil
}

func (g Textfile) validate(loader Loader, config config.Loader) error {
	return nil
}
