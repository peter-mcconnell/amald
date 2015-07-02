package defs

import (
	log "github.com/Sirupsen/logrus"
	"time"
)

type Config struct {
	Loaders map[string]map[string]string `json:loaders,omitempty`
	Reports map[string]map[string]string `json:reports,omitempty`
	Storage map[string]map[string]string `json:storage,omitempty`
	Tests   map[string]bool
}

type SiteDefinition struct {
	Url          string `json:"url"`
	IsLockedDown bool   `json:"islockeddown"`
}

type Results struct {
	Timestamp string           `json:"timestamp"`
	Results   []SiteDefinition `json:"results"`
}

type Records struct {
	Records []Results `json:"records"`
}

type Analysis struct {
	Since map[string][]SiteDefinition
}

// SiteDefinitionsToResults takes a series of SiteDefinition and turns them into a format
// we can use for our storage
func SiteDefinitionsToResults(scanResults []SiteDefinition) Results {
	// combine metadata and scan results, then convert to json
	data := Results{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Results:   scanResults,
	}
	return data
}

func AnalyseData(results Results) Analysis {
	log.Debug("Analysing data")
	analysis := Analysis{}
	for _, sd := range results.Results {
		log.Debugf("%s is %b", sd.Url, sd.IsLockedDown)
	}
	return analysis
}
