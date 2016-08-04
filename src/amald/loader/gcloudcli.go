package loader

import "amald/config"

/*
Gcloudcli implementes the Loader interface and is designed to
load urls from the gcloud cli
*/
type Gcloudcli struct{}

func (g Gcloudcli) load(loader config.Loader) ([]config.Target, error) {
	var targets []config.Target

	return targets, nil
}

func (g Gcloudcli) validate(loader Loader, config config.Loader) error {
	return nil
}
