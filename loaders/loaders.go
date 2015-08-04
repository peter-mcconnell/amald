package loaders

import (
	log "github.com/Sirupsen/logrus"
)

type UrlsLoader interface {
	FetchUrls() ([]string, error)
}

var (
	loaders = make(map[string]UrlsLoader)
)

// GetLoaders will check through the activeloaders it's passed and enable them
// accordingly. This is set up so that these 'loaders' can be enabled via the
// config.yaml
func GetLoaders(activeloaders map[string]map[string]string) {

	log.Debugf("get loaders: %+v", activeloaders)
	// check to see if gcloudcli has been specified in the config
	if _, ok := activeloaders["gcloudcli"]; ok {
		// gcloudcli has been specified. See if we can run it
		if gcloudcliLoaderAvailable() {
			log.Debug("gcloud cli loader available")
			loaders["gcloudcli"] = &LoaderGcloudCLI{}
		}
	}

	// check to see if a textfile has been provided
	if _, ok := activeloaders["textfile"]; ok {
		// textfile has been specified. See if we can run it
		if textfileLoaderAvailable(activeloaders["textfile"]) {
			log.Debug("textfile loader available")
			loaders["textfile"] = &LoaderTextfile{}
		}
	}
}

// Collect URL information from each of the loaders
func CollectUrls() ([]string, error) {
	log.Debug("collect urls")
	m := []string{}
	for _, loader := range loaders {
		if f, err := loader.FetchUrls(); err == nil {
			for _, v := range f {
				m = append(m, v)
			}
		} else {
			log.Fatalf("FetchUrls failed: %s", err)
			return m, err
		}
	}
	log.Debug("found these urls:\n", m)
	return m, nil
}
