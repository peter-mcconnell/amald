package defs

import (
	log "github.com/Sirupsen/logrus"
	"sort"
	"time"
)

var (
	StateKeys = map[string]int{
		"removed": 0,
		"created": 1,
		"updated": 2,
		"same":    3,
	}
	AllowedDistanceHoursOffset = 0.1
	TS                         string
)

type Config struct {
	Global           map[string]string            `json:"global"`
	Loaders          map[string]map[string]string `json:"loaders",omitempty`
	Reports          map[string]map[string]string `json:"reports",omitempty`
	Storage          map[string]map[string]string `json:"storage",omitempty`
	SummaryIntervals []IntervalSettings           `json:"summaryintervals"`
	ShowSameState    bool                         `json:"showsamestate"`
	Tests            map[string]bool
}

type IntervalSettings struct {
	Title         string `json:"title"`
	DistanceHours int    `json:"distancehours"`
	Ansii         string `json:"ansii"`
}

// SiteDefinition is the information amald holds about a single URL
type SiteDefinition struct {
	Url            string `json:"url"`
	IsLockedDown   bool   `json:"islockeddown"`
	HttpStatusCode int    `json:"httpstatuscode",omitempty`
}

// Results holds a timestamp for when the scan was ran and the associated
// SiteDefinitions it found for that scan
type Results struct {
	Timestamp string           `json:"timestamp"`
	Results   []SiteDefinition `json:"results"`
}

// Records is simply a list of results. This is a reflection of the storage
type Records struct {
	Records []Results `json:"records"`
}

// Analysis should have a key range of 0-3 (deleted, created, updated, same)
type Analysis map[int][]SiteDefinition

type Summaries map[int]Analysis

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

// SetTS simply sets the timestamp which defs uses to stamp records with
func SetTS(ts string) {
	TS = ts
}

// SiteDefinitionsToRecords takes a series of SiteDefinition and turns them
// into a format we can use for our storage
func SiteDefinitionsToRecords(scanResults []SiteDefinition) Records {
	// combine metadata and scan results, then convert to json
	records := Records{}
	records.Records = append(records.Records, Results{
		Timestamp: TS,
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
	log.Debugf("Analysing data. Found %d records", len(r.Records))
	// lets sort the Records (newest first)
	sort.Sort(r)
	// the first item is the scan which we just performed
	now := r.Records[0]
	summaries := Summaries{}
	// keep a reference of which hours the user has requested
	sikeyref := make(map[int][]int)
	for k, si := range cfg.SummaryIntervals {
		if _, ok := sikeyref[si.DistanceHours]; !ok {
			sikeyref[si.DistanceHours] = []int{k}
		} else {
			// supports the user having more than one entry for a
			// given distance. not sure this has any value, but
			// will at least give the user an expected result
			sikeyref[si.DistanceHours] = append(
				sikeyref[si.DistanceHours], k)
		}
	}
	log.Debugf("SummaryIntervalKeyRef: %+v", sikeyref)
	// loop the rest of our records
	for _, rec := range r.Records[1:] {
		if distance_hours, err := DistanceHours(now.Timestamp, rec.Timestamp); err == nil {
			log.Debugf("~disthrs: %f", distance_hours)
			dh := int(distance_hours)
			// if the calculated distance isn't matched, try
			// offsetting the value
			if _, ok := sikeyref[dh]; !ok {
				dh = FactorOffset(distance_hours)
			}
			log.Debugf("~dh: %d", dh)
			// has this distance been requested by the user?
			if _, ok := sikeyref[dh]; ok {
				log.Debugf("Found record with distance (%d) "+
					"specified by user: %s",
					distance_hours, rec.Timestamp)
				// loop through each iteration of this distance
				// that the user has provided. this will likely
				// just be the 1 item
				for _, k := range sikeyref[int(distance_hours)] {
					summaries[k] = CompareRecords(
						now.Results, rec.Results)
				}
			}
		}
	}

	return summaries
}

// FactorOffset is designed to allow for 'some give' when comparing timestamps.
// The offset here is defined by AllowDistanceHoursOffset
func FactorOffset(x float64) int {
	if int(x+AllowedDistanceHoursOffset) != int(x) {
		return int(x + AllowedDistanceHoursOffset)
	}
	return int(x)
}

// CompareRecords takes two sets of results and compares them
func CompareRecords(master, other []SiteDefinition) Analysis {
	log.Debug("ComparingRecords...")
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

// DistanceHours takes two timestamp strings and returns the difference in
// hours
func DistanceHours(a, b string) (float64, error) {
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
	return ta.Sub(tb).Hours(), nil
}
