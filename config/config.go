package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
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
		log.Error("Config file not found")
		return config, err
	}
	// read file
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("Config file could not be read")
		return config, err
	}
	// unmarshal
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Error("Config file could not be unmarshalled")
		return config, err
	}

	// merge defaults
	config = loadDefaults(config)

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

// validateStorageSettings inspects a Config and returns a boolean based on the
// Config being valid. This is designed to try and make sure the user has
// inserted the correct values in their config.yaml
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
			log.Errorf("Failed to set abs filepath on %s: %s",
				config.Storage["json"]["path"], err)
		} else {
			if _, err = os.Stat(spath); err != nil {
				log.Errorf("storage settings listed file %s"+
					"which couldnt be loaded: %s", spath,
					err)
				return settingsValidated, err
			} else {
				settingsValidated = true
			}
		}
	}
	return settingsValidated, err
}

// loadDefaults is a method that allows you to define some default values for
// the Config. Any missing fields from the config yaml then they should be set
// here.
func loadDefaults(config defs.Config) defs.Config {
	if config.Reports == nil {
		config.Reports = make(map[string]map[string]string)
	}
	return config
}
