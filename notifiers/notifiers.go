package notifiers

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/loaders"
	"strings"
)

type NotifierLoader interface {
	Send(map[string]map[string]string, string)
}

type JsonData struct {
	Meta map[string]string        `json:"Meta"`
	Data []loaders.SiteDefinition `json:"Data"`
}

func FireNotifiers(jsonbytes []byte,
	activeloaders map[string]map[string]string) error {

	data, err := loadData(jsonbytes)

	// Data from the first item forms the active report
	urls := data[0].Data
	if err != nil {
		log.Fatalf("there was a problem loading the existing data: %s", err)
	}

	// check to see if ascii has been specified in the config
	if _, ok := activeloaders["ascii"]; ok {
		n := &NotifierAscii{urls: urls}
		n.Send()
	}

	// check to see if mailgun has been specified in the config
	if _, ok := activeloaders["mailgun"]; ok {
		n := &NotifierMailgun{urls: urls}
		n.Send(activeloaders["mailgun"])
	}

	return nil
}

func loadData(jsonbytes []byte) ([]JsonData, error) {
	var (
		datalist []JsonData
		data     JsonData
		err      error
	)

	rows := strings.Split(string(jsonbytes), "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		jsonParser := json.NewDecoder(bytes.NewBufferString(row))
		if err = jsonParser.Decode(&data); err != nil {
			log.Fatalf("error parsing data file: %s", err)
		}
		datalist = append(datalist, data)
	}

	return datalist, err
}
