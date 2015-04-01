package storage

import (
	"bufio"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var (
	urls        []defs.SiteDefinition
	recordLimit int = 100 // default value (if not set in config)
)

// StoreData adds a new entry of urls to the specified json file. It returns
// a byte array of ALL records and an error type
func StoreData(passedurls []defs.SiteDefinition,
	config map[string]map[string]string) ([]byte, error) {
	urls = passedurls
	var (
		data []byte
		err  error
	)

	// do we need to store data to a json file?
	if _, ok := config["json"]; ok {
		// do we need to update the recordLimit?
		if _, ok := config["json"]["recordlimit"]; ok {
			if i, err := strconv.Atoi(config["json"]["recordlimit"]); err == nil {
				recordLimit = i
			} else {
				log.Warnf("Failed to conver recordlimit value to int: %s", err)
			}
		}
		// store data to a json file
		data, err = jsonStoreData(config["json"])
		if err != nil {
			log.Fatalf("jsonStoreData failed: %s", err)
		}
	}

	return data, nil
}

// jsonStoreData takes any existing information in the json file and adds in
// an entry for the active scan
func jsonStoreData(config map[string]string) ([]byte, error) {

	// create a meta map to allow us to throw some extra info in
	meta := make(map[string]string)
	meta["utc_timestamp"] = time.Now().UTC().Format("")

	// marshal data
	jstr, err := json.Marshal(struct {
		Meta map[string]string
		Data []defs.SiteDefinition
	}{
		Meta: meta,
		Data: urls,
	})
	if err != nil {
		log.Fatalf("failed to marshal urls in StoreData: %s", err)
	}
	jstr = append(jstr, []byte("\n")...) // NL gives each entry its own row

	// load existing data
	datab, err := loadExistingTextfileData(config)
	if err != nil {
		// Only a warning as this will fail if a json file hasn't been stored
		// previously
		log.Warnf("Failed to load data file: %s", err)
	}
	data := append(jstr, datab...)

	// write to disk
	err = ioutil.WriteFile(config["path"], data, 0644)
	if err != nil {
		log.Fatalf("failed to writefile: %s", err)
	}
	return data, err
}

// loadExistingData creates an array of bytes from the specified data path
// and trims it to a certain number of lines (recordLimit) if needs be
func loadExistingTextfileData(config map[string]string) ([]byte, error) {
	file, err := os.Open(config["path"])
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []byte
	scanner := bufio.NewScanner(file)
	i := 2 // account for adding another line in for the active scan
	for scanner.Scan() {
		lines = append(lines, scanner.Bytes()...)
		// make sure there's still a newline between each row
		lines = append(lines, []byte("\n")...)
		// only capture enough lines required
		i++
		if i > recordLimit {
			break
		}
	}
	return lines, scanner.Err()
}
