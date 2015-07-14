package defs

import (
	"testing"
)

var (
	config         Config
	scanResults    []SiteDefinition
	oldScanResults []SiteDefinition
	records        Records
)

func init() {
	config = Config{
		SummaryIntervals: []IntervalSettings{
			IntervalSettings{
				Title:         "yesterday",
				DistanceHours: 24,
			},
			IntervalSettings{
				Title:         "last week",
				DistanceHours: 168,
			},
			IntervalSettings{
				Title:         "last month",
				DistanceHours: 720,
			},
		},
	}
	scanResults = []SiteDefinition{ // this is the same
		SiteDefinition{
			Url:            "http://a",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		SiteDefinition{ // this has changed
			Url:            "http://b",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		SiteDefinition{ // this has changed
			Url:            "http://c",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		SiteDefinition{ // this is new
			Url:            "http://d",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	oldScanResults = []SiteDefinition{
		SiteDefinition{
			Url:            "http://a",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		SiteDefinition{
			Url:            "http://b",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		SiteDefinition{
			Url:            "http://c",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		SiteDefinition{ // this is removed in scanResults
			Url:            "http://e",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
	}
	records.Records = []Results{
		Results{
			Timestamp: "2015-07-14T11:00:00Z",
			Results:   scanResults,
		},
		Results{ // 1 day old
			Timestamp: "2015-07-13T11:00:00Z",
			Results:   oldScanResults,
		},
		Results{ // 7 days old
			Timestamp: "2015-07-07T11:00:00Z",
			Results:   oldScanResults,
		},
		Results{ // 30 days old
			Timestamp: "2015-06-14T11:00:00Z",
			Results:   oldScanResults,
		},
		Results{ // this shouldn't be picked up in our report
			Timestamp: "2015-07-10T11:00:00Z",
			Results:   oldScanResults,
		},
	}
}

func TestDistanceHours(t *testing.T) {
	if d, err := DistanceHours("2015-07-14T11:00:00Z", "2015-06-14T11:00:00Z"); err != nil {
		t.Errorf("DistanceHours returned error: %s", err)
	} else if d != 720 {
		t.Errorf("DistanceHours returned an unexpected value: %d", d)
	}
}

func TestSiteDefinitionsToRecords(t *testing.T) {
	sitedefs := []SiteDefinition{
		SiteDefinition{
			Url:          "https://google.com",
			IsLockedDown: false,
		},
		SiteDefinition{
			Url:          "https://test.com/",
			IsLockedDown: true,
		},
	}
	records := SiteDefinitionsToRecords(sitedefs)
	if len(records.Records[0].Results) != 2 {
		t.Errorf("Got unexpected number of records: ", len(records.Records))
	}
}

func TestAnalyseRecords(t *testing.T) {

	analysis := AnalyseRecords(config, records)

	// running tests against 'yesterday'
	if len(analysis["yesterday"]) != 4 {
		t.Error("Didn't get expected number of analysis results for 'yesterday'")
	} else if len(analysis["yesterday"][0]) != 1 {
		t.Errorf("Got unexpected number of analysis results for 'yesterday', removed: %s", len(analysis["yesterday"][0]))
	} else if analysis["yesterday"][0][0].Url != "http://e" {
		t.Error("Didnt get expected result for 'yesterday'[0][0]")
	} else if len(analysis["yesterday"][1]) != 1 {
		t.Errorf("Got unexpected number of analysis results for 'yesterday', created: %s", len(analysis["yesterday"][1]))
	} else if analysis["yesterday"][1][0].Url != "http://d" {
		t.Error("Didnt get expected result for 'yesterday'[1][0]")
	} else if len(analysis["yesterday"][2]) != 2 {
		t.Errorf("Got unexpected number of analysis results for 'yesterday', updated: %s", len(analysis["yesterday"][2]))
	} else if analysis["yesterday"][2][0].Url != "http://b" {
		t.Error("Didnt get expected result for 'yesterday'[2][0]")
	} else if analysis["yesterday"][2][1].Url != "http://c" {
		t.Error("Didnt get expected result for 'yesterday'[2][1]")
	} else if len(analysis["yesterday"][3]) != 1 {
		t.Errorf("Got unexpected number of analysis results for 'yesterday', same: %s", len(analysis["yesterday"][3]))
	} else if analysis["yesterday"][3][0].Url != "http://a" {
		t.Error("Didnt get expected result for 'yesterday'[3][0]")
	}
}
