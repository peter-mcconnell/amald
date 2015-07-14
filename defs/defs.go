package defs

import (
	log "github.com/Sirupsen/logrus"
	"sort"
	"time"
)

var StateKeys = map[string]int{
	"removed": 0,
	"created": 1,
	"updated": 2,
	"same":    3,
}

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

type Analysis map[int][]SiteDefinition // map[int] = 0-3 (deleted, created, updated, same)

type Summaries map[string]Analysis

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
func AnalyseRecords(cfg Config, r Records) Summaries {
	log.Debug("Analysing data")
	// lets sort the Records (newest first)
	sort.Sort(r)
	// the first item is the scan which we just performed
	now := r.Records[0]
	summaries := Summaries{}
	// keep a reference of which days the user has requested
	sikeyref := make(map[int][]int)
	for k, si := range cfg.SummaryIntervals {
		if _, ok := sikeyref[si.DistanceDays]; !ok {
			sikeyref[si.DistanceDays] = []int{k}
		} else {
			// supports the user having more than one entry for a given distance.
			// not sure this has any value, but will at least give the user an
			// expected result
			sikeyref[si.DistanceDays] = append(sikeyref[si.DistanceDays], k)
		}
	}
	// loop the rest of our records
	for _, rec := range r.Records[1:] {
		if distance_hours, err := distanceHours(now.Timestamp, rec.Timestamp); err == nil {
			// has this distance been requested by the user?
			if _, ok := sikeyref[distance_hours/24]; ok {
				// loop through each iteration of this distance that the user
				// has provided. this will likely just be the 1 item
				for _, k := range sikeyref[distance_hours/24] {
					log.Debug(k)
					summaries[cfg.SummaryIntervals[k].Title] = CompareRecords(now.Results, rec.Results)
				}
			}
		}
	}

	return summaries
}

// CompareRecords takes two sets of results and compares them
func CompareRecords(master, other []SiteDefinition) Analysis {
	a := Analysis{}
	// store urls
	mstrByUrl := make(map[string]SiteDefinition)
	for _, sd := range master {
		mstrByUrl[sd.Url] = sd
	}
	for _, sd := range other {
		k := "same" // default state
		if mstr, ok := mstrByUrl[sd.Url]; ok {
			// url exists in master
			if sd.IsLockedDown != mstr.IsLockedDown {
				// IsLockedDown has changed
				k = "updated"
			}
		} else {
			// url does not exist in master (removed)
			k = "removed"
		}
		a[StateKeys[k]] = append(a[StateKeys[k]], sd)
		// delete the key from mstr, so we know whats left at the end
		delete(mstrByUrl, sd.Url)
	}
	if len(mstrByUrl) != 0 {
		// some sitedefinitions where left. these must be new
		for _, sd := range mstrByUrl {
			a[StateKeys["created"]] = append(a[StateKeys["created"]], sd)
		}
	}
	return a
}

// distanceHours takes two timestamp strings and returns the difference in days
func distanceHours(a, b string) (int, error) {
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
