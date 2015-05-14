package reports

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"math"
	"time"
)

type DataDiff struct {
	AddedApps   map[string]bool
	RemovedApps map[string]bool
	UpdatedApps map[string]bool
}

type ReportLoader interface {
	Generate() (string, error)
}

// findDiffs is designed to scan all of the available data to try and pull out
// some comparitive differences in the data sets.
// data is in order (newest first) which gives us some advantages for marking
// out which data to use for - the capture points will keep getting overwritten
// with the earliest available entry for that day
func FindDiffs(data []defs.JsonData) (map[string]DataDiff, error) {
	var err error
	diffdata := make(map[string]defs.JsonData)
	diffreturn := make(map[string]DataDiff)
	for _, d := range data {
		if then, err := time.Parse(time.RFC3339, d.Meta["timestamp"]); err == nil {
			distance := math.Ceil(time.Since(then).Hours()/24) - 1
			log.Debug(distance)
			// no need to inspect past 30 days
			if distance > 30 {
				break
			}
			switch distance {
			case 1:
				diffdata["yesterday"] = d
			case 7:
				diffdata["lastweek"] = d
			case 30:
				diffdata["thirtydays"] = d
			}
		}
	}
	// now we (might) have the relevant data selected, run comparisons against
	// the latest scan [0]
	if _, ok := diffdata["yesterday"]; ok {
		diffreturn["yesterday"] = compareData(diffdata["yesterday"], data[0])
	}
	if _, ok := diffdata["lastweek"]; ok {
		diffreturn["lastweek"] = compareData(diffdata["lastweek"], data[0])
	}
	if _, ok := diffdata["thirtydays"]; ok {
		diffreturn["thirtydays"] = compareData(diffdata["thirtydays"], data[0])
	}
	return diffreturn, err
}

// compareData is designed to inspect and compare two different maps. It
// returns a DataDiff struct detailing the differences that it finds
func compareData(oldscan defs.JsonData, newscan defs.JsonData) DataDiff {
	diffreturn := DataDiff{
		AddedApps:   make(map[string]bool),
		RemovedApps: make(map[string]bool),
		UpdatedApps: make(map[string]bool)}
	matchedurls := make(map[string]bool)
	// loop through the old data to assess what's changed
	for _, def := range oldscan.Data {
		// does this url appear in both datasets?
		if _, newdata := newscan.Data[def.Url]; newdata {
			// was the lockdown status different?
			if newscan.Data[def.Url].IsLockedDown != def.IsLockedDown {
				log.Debug("url: %s has changed lockdown status to %s",
					def.Url, newscan.Data[def.Url].IsLockedDown)
				diffreturn.UpdatedApps[def.Url] = newscan.Data[def.Url].IsLockedDown
			}
			// track matched urls so at the end we can assess what's left
			// (will leave new urls).
			matchedurls[def.Url] = true
		} else {
			// url is not in the newscan (removed)
			diffreturn.RemovedApps[def.Url] = newscan.Data[def.Url].IsLockedDown
			log.Debugf("url: %s appears to have been removed", def.Url)
		}
	}
	// now that we have checked all of the old data, we can see what's left
	// from the new scan, which we will determine as new urls
	for _, data := range newscan.Data {
		if _, ok := matchedurls[data.Url]; !ok {
			diffreturn.AddedApps[data.Url] = data.IsLockedDown
			log.Debugf("url: %s appears to be new", data.Url)
		}
	}
	return diffreturn
}
