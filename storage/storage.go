package storage

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/loaders"
	"io/ioutil"
	"time"
)

var (
	urls []loaders.SiteDefinition
)

func LoadData() {

}

func StoreData(passedurls []loaders.SiteDefinition,
	config map[string]map[string]string) error {
	urls = passedurls

	// store data to a json file
	if _, ok := config["json"]; ok {
		err := jsonStoreData(config["json"])
		if err != nil {
			log.Fatalf("jsonStoreData failed: %s", err)
		}
	}

	return nil
}

func CompareData() error {
	// load old data
	LoadData()
	return nil
}

func jsonStoreData(config map[string]string) error {

	// create a meta map to allow us to throw some extra info in
	meta := make(map[string]string)
	meta["utc_timestamp"] = time.Now().UTC().String()

	// marshal data
	jstr, err := json.Marshal(struct {
		Meta map[string]string
		Data []loaders.SiteDefinition
	}{
		Meta: meta,
		Data: urls,
	})
	if err != nil {
		log.Fatalf("failed to marshal urls in StoreData: %s", err)
	}

	// write to disk
	err = ioutil.WriteFile(config["path"], jstr, 0644)
	if err != nil {
		log.Fatalf("failed to writefile: %s", err)
	}
	return err
}
