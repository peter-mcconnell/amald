package loaders

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
)

type UrlsLoader interface {
	FetchUrls() map[string]defs.SiteDefinition
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
func CollectUrls() map[string]defs.SiteDefinition {
	m := map[string]defs.SiteDefinition{}
	for _, loader := range loaders {
		f := loader.FetchUrls()
		for k, v := range f {
			m[k] = v
		}
	}
	return m
}
