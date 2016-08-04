package loader

import (
	"log"

	"amald/config"
)

/*
Loader's should be able to validate() and load()
*/
type Loader interface {
	validate(loader Loader, config config.Loader) error
	load(config.Loader) ([]config.Target, error)
}

/*
Load accepts a list of Loaders, iterates over each and performs
the relevant load method. Throws a fatal error if the loader is
invalid
*/
func Load(configLoaders []config.Loader) []config.Target {
	var targets []config.Target
	var loadMethod Loader
	for _, config := range configLoaders {
		switch config.Type {
		case "gcloudcli":
			loadMethod = Gcloudcli{}
		case "textfile":
			loadMethod = Textfile{}
		default:
			log.Fatalf("Unrecognised loader type '%s'. Please check config", config.Type)
		}

		if err := loadMethod.validate(loadMethod, config); err != nil {
			log.Fatalf("Failed to validate loader:\n%s", err)
		}

		urls, err := loadMethod.load(config)
		if err != nil {
			log.Fatalf("Failed to load from loader:\n%s", err)
		}

		targets = append(targets, urls...)
	}

	return targets
}
