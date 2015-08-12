package storage

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"io/ioutil"
)

// StoreScan creates a new entry in storage of the current scan
func StoreScan(path string, data defs.Records) error {
	log.Debug("Storing scan results")

	// convert map to string
	jstr, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Failed to Marshal data: %s", err)
	}

	// open our storage file so that we can append to it
	err = ioutil.WriteFile(path, jstr, 0644)
	if err != nil {
		log.Errorf("Failed to writefile: %s", err)
	}

	return err
}

// LoadSiteDefsFromStorage will simply load all the recorded SiteDefinitions from storage
func LoadSiteDefsFromStorage(path string) (defs.Records, error) {
	log.Debug("get summary")
	summary := defs.Records{}
	data, err := loadStorageDataFromFile(path)
	if err == nil {
		if string(data) != "" {
			// turn string into map
			err = json.Unmarshal(data, &summary)
			if err == nil {
				log.Debugf("LoadSiteDefsFromStorage results: %+v", summary)
			}
		}
	} else {
		log.Errorf("Failed to loadStorageDataFromFile: %s", err)
	}
	return summary, err
}

// loadStorageDataFromFile returns a byte array of all the available
// information from storage
func loadStorageDataFromFile(path string) ([]byte, error) {
	log.Debug("loading storage data from file")
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("Problem reading file: %s", err)
		return f, err
	} else {
		log.Debugf("loadStorageDataFromFile results: %s", string(f))
	}
	return f, nil
}
