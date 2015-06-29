package storage

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"io/ioutil"
	"time"
)

// MergeData simply takes the scanResults and merges it into the existing data
func MergeData(scanResults []defs.SiteDefinition, olddata map[string][]defs.SiteDefinition) map[string][]defs.SiteDefinition {

	// turn our scanResults into a slightly more informative json string
	newdata := formData(scanResults)

	log.Debugf("olddata %+v", olddata)
	log.Debugf("newdata %+v", newdata)

	// add newdata to olddata
	merged := olddata
	for ts, data := range newdata {
		merged[ts] = data
	}

	log.Debugf("merged %+v", merged)

	return merged
}

// LoadSiteDefsFromStorage will simply load all the recorded SiteDefinitions from storage
func LoadSiteDefsFromStorage(path string) (map[string][]defs.SiteDefinition, error) {
	summary := make(map[string][]defs.SiteDefinition)
	log.Debug("get summary")
	data, err := loadStorageDataFromFile("tmp/test.json") //path)
	if err == nil {
		// turn string into map
		//summaryJson := defs.JsonFormat{}
		summaryJson := defs.TestFormat{}
		err = json.Unmarshal(data, &summaryJson)
		if err == nil {
			log.Debugf("LoadSiteDefsFromStorage results: %+v", summary)
		}
	} else {
		log.Fatalf("Failed to loadStorageDataFromFile: %s", err)
	}
	log.Debugf("x %+v", string(data))
	log.Fatalf("hey %+v", summary)
	return summary, err
}

// loadStorageDataFromFile returns a byte array of all the available
// information from storage
func loadStorageDataFromFile(path string) ([]byte, error) {
	log.Debug("loading storage data from file")
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Problem reading file: %s", err)
		return f, err
	} else {
		log.Debugf("loadStorageDataFromFile results: %s", string(f))
	}
	return f, nil
}

// StoreScan creates a new entry in storage of the current scan
func StoreScan(path string, data map[string][]defs.SiteDefinition) error {
	log.Debug("Storing scan results")

	// convert map to string
	jstr, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to Marshal data: %s", err)
	}

	// open our storage file so that we can append to it
	err = ioutil.WriteFile(path, jstr, 0644)
	if err != nil {
		log.Fatalf("Failed to writefile: %s", err)
	}

	return err
}

// formData takes a series of SiteDefinition and turns them into a format
// we can use for our storage
func formData(scanResults []defs.SiteDefinition) map[string][]defs.SiteDefinition {
	// combine metadata and scan results, then convert to json
	data := make(map[string][]defs.SiteDefinition)
	data[time.Now().UTC().Format(time.RFC3339)] = scanResults
	return data
}
