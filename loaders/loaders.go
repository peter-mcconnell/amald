package loaders

import (
	log "github.com/Sirupsen/logrus"
)

type SiteDefinition struct {
	Url          string
	IsLockedDown bool
}

type UrlsLoader interface {
	FetchUrls() []SiteDefinition
}

var (
	loaders = make(map[string]UrlsLoader)
)

// Grab each loader
func GetLoaders(activeloaders map[string]map[string]string) error {

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

	return nil
}

// Collect URL information from each of the loaders
func CollectUrls() []SiteDefinition {
	m := []SiteDefinition{}
	for _, loader := range loaders {
		m = append(m, loader.FetchUrls()...)
	}
	return m
}
