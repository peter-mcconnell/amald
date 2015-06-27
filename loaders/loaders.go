package loaders

import (
	log "github.com/Sirupsen/logrus"
)

type UrlsLoader interface {
	FetchUrls() []string
}

var (
	loaders = make(map[string]UrlsLoader)
)

func GetLoaders(activeloaders map[string]map[string]string) {
	log.Debug("get loaders")
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
func CollectUrls() []string {
	log.Debug("collect urls")
	m := []string{}
	for _, loader := range loaders {
		f := loader.FetchUrls()
		for _, v := range f {
			m = append(m, v)
		}
	}
	log.Debug("found these urls:\n", m)
	return m
}
