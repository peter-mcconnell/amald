package defs

import (
	log "github.com/Sirupsen/logrus"
	"math"
	"time"
)

type Config struct {
	Loaders          map[string]map[string]string `json:loaders,omitempty`
	Reports          map[string]map[string]string `json:reports,omitempty`
	Storage          map[string]map[string]string `json:storage,omitempty`
	SummaryIntervals []IntervalSettings           `json:summary_intervals,omitempty`
	Tests            map[string]bool
}

type IntervalSettings struct {
	Title        string `json:"title"`
	DistanceDays int    `json:"distance_days"`
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

// SiteDefinitionsToRecords takes a series of SiteDefinition and turns them into a format
// we can use for our storage
func SiteDefinitionsToRecords(scanResults []SiteDefinition) Records {
	// combine metadata and scan results, then convert to json
	records := Records{}
	records.Records = append(records.Records, Results{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Results:   scanResults,
	})
	return records
}

// AnalyseRecords compares the most recent result against other entries
func AnalyseRecords(cfg Config, r Records) Analysis {
	log.Debug("Analysing data")
	analysis := Analysis{}

	//@TODO: implement sort.Interface on Analysis

	return analysis
}
