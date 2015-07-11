package defs

import (
	log "github.com/Sirupsen/logrus"
	"sort"
	"time"
)

type Config struct {
	Loaders          map[string]map[string]string `json:"loaders",omitempty`
	Reports          map[string]map[string]string `json:"reports",omitempty`
	Storage          map[string]map[string]string `json:"storage",omitempty`
	SummaryIntervals []IntervalSettings           `json:"summaryintervals"`
	Tests            map[string]bool
}

type IntervalSettings struct {
	Title        string `json:"title"`
	DistanceDays int    `json:"distancedays"`
}

type SiteDefinition struct {
	Url            string `json:"url"`
	IsLockedDown   bool   `json:"islockeddown"`
	HttpStatusCode int    `json:"httpstatuscode",omitempty`
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

// Implement sort interface on our Records struct
func (r Records) Len() int {
	return len(r.Records)
}
func (r Records) Less(i, j int) bool {
	b, err := time.Parse(time.RFC3339, r.Records[i].Timestamp)
	if err != nil {
		log.Errorf("Failed to parse time 'b': %s", err)
	}
	a, err := time.Parse(time.RFC3339, r.Records[j].Timestamp)
	if err != nil {
		log.Errorf("Failed to parse time 'a': %s", err)
	}
	return b.After(a)
}
func (r Records) Swap(i, j int) {
	r.Records[i], r.Records[j] = r.Records[j], r.Records[i]
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

type SiKeyStore []int

// Add a sort interface to SiKeyStore
func (s SiKeyStore) Len() int           { return len(s) }
func (s SiKeyStore) Less(i, j int) bool { return s[i] < s[j] }
func (s SiKeyStore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// AnalyseRecords compares the most recent result against other entries
func AnalyseRecords(cfg Config, r Records) Analysis {
	log.Debug("Analysing data")
	analysis := Analysis{}
	si := SiKeyStore{}
	// keep a map of interval values
	//for k, interval := range cfg.SummaryIntervals {
	//	si[interval.DistanceDays] = k
	//}
	sort.Sort(si)
	log.Debug(cfg.SummaryIntervals)
	log.Fatal(si)
	// lets sort the Records (newest first)
	sort.Sort(r)
	// the first item is the scan which we just performed
	now := r.Records[0]
	for distance, k := range si {
		log.Debugf("distance %s, key %s", distance, k)
	}
	// delete:
	for _, rec := range r.Records[1:] {
		if distance_days, err := distanceDays(now.Timestamp, rec.Timestamp); err == nil {
			log.Debugf("distance: %s, err: %s", distance_days, err)
		}
	}

	return analysis
}

// distanceDays takes two timestamp strings and returns the difference in days
func distanceDays(a, b string) (int, error) {
	ta, err := time.Parse(time.RFC3339, a)
	if err != nil {
		log.Errorf("Failed to parse time: %s")
		return 0, err
	}
	tb, err := time.Parse(time.RFC3339, b)
	if err != nil {
		log.Errorf("Failed to parse time: %s")
		return 0, err
	}
	return int(ta.Sub(tb).Hours()), nil
}
