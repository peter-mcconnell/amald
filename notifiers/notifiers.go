package notifiers

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"strings"
)

type NotifierLoader interface {
	Send(map[string]map[string]string, string)
}

func FireNotifiers(jsonbytes []byte,
	activeloaders map[string]map[string]string) error {

	data, err := loadData(jsonbytes)
	if err != nil {
		log.Fatalf("there was a problem loading the existing data: %s", err)
	}

	// check to see if ascii has been specified in the config
	if _, ok := activeloaders["ascii"]; ok {
		n := &NotifierAscii{data: data}
		n.Send()
	}

	// check to see if mailgun has been specified in the config
	if _, ok := activeloaders["mailgun"]; ok {
		n := &NotifierMailgun{data: data}
		n.Send(activeloaders["mailgun"])
	}

	return nil
}

func loadData(jsonbytes []byte) ([]defs.JsonData, error) {
	var (
		datalist []defs.JsonData
		data     defs.JsonData
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
