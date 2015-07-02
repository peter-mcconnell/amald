package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// LoadConfig returns a Config struct. It builds the config using the provided
// filepath. loadDefaults called after the yaml file is scanned to assign
// defaults if they are needed.
func Load(path string) (defs.Config, error) {
	log.Debug("load config")
	var (
		config defs.Config
	)

	// init Tests map
	config.Tests = make(map[string]bool)

	// get absolute path
	path, err := filepath.Abs(path)
	if err != nil {
		return config, err
	}
	// does file exist?
	if _, err := os.Stat(path); err != nil {
		log.Fatal("Config file not found")
		return config, err
	}
	// read file
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config file could not be read")
		return config, err
	}
	// unmarshal
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatal("Config file could not be unmarshalled")
		return config, err
	}

	// merge defaults
	config, err = loadDefaults(config)
	if err != nil {
		return config, err
	}

	// validate storage
	var valid bool
	valid, err = validateStorageSettings(config)
	if err != nil {
		return config, err
	}
	config.Tests["storage"] = valid

	log.Debugf("config loaded: %+v", config)

	return config, err
}

func validateStorageSettings(config defs.Config) (bool, error) {
	log.Debug("validateStorageSettings")
	var (
		settingsValidated bool  = false
		err               error = nil
	)
	// clear storage settings, if they aren't available
	if _, ok := config.Storage["json"]["path"]; ok {
		log.Debug("storage settings found in config")
		var spath, err = filepath.Abs(config.Storage["json"]["path"])
		if err != nil {
			log.Errorf("Failed to set abs filepath on %s: %s", config.Storage["json"]["path"], err)
		} else {
			if _, err = os.Stat(spath); err != nil {
				log.Errorf("storage settings listed file %s which couldnt be loaded: %s", spath, err)
			} else {
				settingsValidated = true
			}
		}
	}

	log.Debug("config.Storage validated: %s", strconv.FormatBool(settingsValidated))

	return settingsValidated, err
}

// loadDefaults is a placeholder at the moment. If default values are to be set
// to cover missing fields from the config yaml then they should be set here.
func loadDefaults(config defs.Config) (defs.Config, error) {
	if config.Reports == nil {
		config.Reports = make(map[string]map[string]string)
	}
	if _, ok := config.Reports["templates"]; !ok {
		config.Reports["templates"] = make(map[string]string)
	}
	if _, ok := config.Reports["templates"]["path"]; !ok {
		config.Reports["templates"]["path"] = ""
	}
	return config, nil
}
