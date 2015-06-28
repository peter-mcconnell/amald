package storage

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"io/ioutil"
	"os"
	"time"
)

func GetSummary(path string) (map[string][]defs.SiteDefinition, error) {
	log.Debug("get summary")
	if data, err := loadStorageDataFromFile(path); err != nil {
		log.Fatalf("Failed to loadStorageDataFromFile: %s", err)
	} else {
		// turn string into map
		summary := defs.JsonFormat{}
		err := json.Unmarshal(data, &summary)
		if err != nil {
			log.Fatalf("Failted to unmarshall: %s", err)
			return summary, err
		}
	}
	return summary, nil
}

// loadStorageDataFromFile returns a byte array of all the available
// information from storage
func loadStorageDataFromFile(path string) ([]byte, error) {
	log.Debug("loading storage data from file")
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Problem reading file: %s", err)
		return f, err
	}
	return f, nil
}

// StoreScan creates a new entry in storage of the current scan
func StoreScan(path string, scanResults []defs.SiteDefinition) {
	log.Debug("Storing scan results")
	// turn our scanResults into a slightly more informative json string
	jstr, err := formData(scanResults)
	if err != nil {
		log.Fatalf("Failed to formData: %s", err)
	}

	// open our storage file so that we can append to it
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Problem closing file: %s", err)
		}
	}()

	if _, err := f.WriteString(jstr); err != nil {
		log.Fatalf("Failed to write to storage: %s", err)
	}
}

func formData(scanResults []defs.SiteDefinition) (string, error) {
	// create a meta map so that we can store some additional information
	meta := make(map[string]string)
	meta["timestamp"] = time.Now().UTC().Format(time.RFC3339)
	// combine metadata and scan results, then convert to json
	j, err := json.Marshal(defs.JsonFormat{
		Meta: meta,
		Data: scanResults,
	})
	// add a newline to the end of this scan entry
	j = append(j, []byte("\n")...)
	return string(j), err
}
