package storage

import (
	"github.com/pemcconnell/amald/defs"
	"testing"
)

var (
	scanResults []defs.SiteDefinition
	olddata     defs.Records
)

func init() {
	scanResults = append(scanResults, defs.SiteDefinition{
		Url:          "https://google.com",
		IsLockedDown: false,
	}, defs.SiteDefinition{
		Url:          "https://test.com/",
		IsLockedDown: true,
	})
	olddata.Records = append(olddata.Records, defs.SiteDefinitionsToRecords(scanResults).Records...)
}

func TestMergeData(t *testing.T) {
	merged := MergeData(scanResults, olddata)
	if len(merged.Records) != 2 {
		t.Error("Didn't get expected number of results from MergeData")
	}
}

func TestLoadSiteDefsFromStorage(t *testing.T) {
	_, err := LoadSiteDefsFromStorage("example.data.json")
	if err != nil {
		t.Errorf("Failed to LoadSiteDefsFromStorage: %s", err)
	}
	_, err = LoadSiteDefsFromStorage("")
	if err == nil {
		t.Error("Should have returned an error for a blank path")
	}
}

func TestStoreScan(t *testing.T) {
	err := StoreScan("", olddata)
	if err == nil {
		t.Error("Should have returned an error for a blank path")
	}
}
