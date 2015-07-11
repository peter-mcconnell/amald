package defs

import (
	"testing"
)

var (
	scanResults []SiteDefinition
	olddata     Records
)

func TestSiteDefinitionsToRecords(t *testing.T) {
	scanResults = append(scanResults, SiteDefinition{
		Url:          "https://google.com",
		IsLockedDown: false,
	}, SiteDefinition{
		Url:          "https://test.com/",
		IsLockedDown: true,
	})
	records := SiteDefinitionsToRecords(scanResults)
	if len(records.Records[0].Results) != 2 {
		t.Errorf("Got unexpected number of records: ", len(records.Records))
	}
}
