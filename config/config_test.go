package config

import (
	"github.com/pemcconnell/amald/defs"
	"testing"
)

var (
	cfg defs.Config
	err error
)

func TestLoad(t *testing.T) {
	// real sample file
	cfg, err = Load("example.config.yaml")
	if err != nil {
		t.Errorf("encountered error: %s", err)
	}
	if _, ok := cfg.Global["templatesdir"]; !ok {
		t.Error("templatesdir not set")
	}

	// file does not exist
	_, err = Load("......")
	if err == nil {
		t.Error("Should have returned an error for fake file")
	}

	// golang file for path (expecting yaml)
	_, err = Load("config_test.go")
	if err == nil {
		t.Error("Should have returned an error for fake file")
	}
}

func TestLoadDefaults(t *testing.T) {
	testcfg := defs.Config{}
	cfg := loadDefaults(testcfg)
	if cfg.Reports == nil {
		t.Error("Reports not being set by LoadDefaults")
	}
}

func TestLoadersExist(t *testing.T) {
	if len(cfg.Loaders) == 0 {
		t.Fatal("Couldn't find any loaders")
	}
}

func TestLoaders(t *testing.T) {
	if _, ok := cfg.Loaders["gcloudcli"]; !ok {
		t.Fatal("Couldn't find the gcloudcli loader")
	}

	if _, ok := cfg.Loaders["textfile"]; !ok {
		t.Fatal("Couldn't find the textfile loader")
	}
}

func TestReportsExist(t *testing.T) {
	if len(cfg.Reports) == 0 {
		t.Fatal("Couldn't find any reports")
	}
}

func TestSummaryIntervals(t *testing.T) {
	if len(cfg.SummaryIntervals) == 0 {
		t.Error("Can't test any SummaryIntervals (none found)")
	} else {
		// ensure our data reflects what we've put in the example file
		if cfg.SummaryIntervals[0].Title != "yesterday" {
			t.Errorf("SummaryIntervals[0].Title wasn't expected: %s",
				cfg.SummaryIntervals[0].Title)
		}
		if cfg.SummaryIntervals[0].DistanceHours != 24 {
			t.Errorf("SummaryIntervals[0].DistanceHours wasn't expected: %s",
				cfg.SummaryIntervals[0].DistanceHours)
		}
		if cfg.SummaryIntervals[1].Title != "last week" {
			t.Errorf("SummaryIntervals[1].Title wasn't expected: %s",
				cfg.SummaryIntervals[1].Title)
		}
		if cfg.SummaryIntervals[1].DistanceHours != 168 {
			t.Errorf("SummaryIntervals[1].DistanceHours wasn't expected: %s",
				cfg.SummaryIntervals[1].DistanceHours)
		}
		if cfg.SummaryIntervals[2].Title != "last month" {
			t.Errorf("SummaryIntervals[2].Title wasn't expected: %s",
				cfg.SummaryIntervals[2].Title)
		}
		if cfg.SummaryIntervals[2].DistanceHours != 720 {
			t.Errorf("SummaryIntervals[2].DistanceHours wasn't expected: %s",
				cfg.SummaryIntervals[2].DistanceHours)
		}
	}

}

func TestSummaryIntervalsExist(t *testing.T) {
	if len(cfg.SummaryIntervals) == 0 {
		t.Fatal("Couldn't find any summaryintervals")
	}
	// 3 is the known number of summaryintervals in the example file
	if len(cfg.SummaryIntervals) != 3 {
		t.Fatal("Found an unexpected number of summaryintervals")
	}

}

func TestReports(t *testing.T) {
	if _, ok := cfg.Reports["ascii"]; !ok {
		t.Fatal("Couldn't find the ascii report")
	}

	if _, ok := cfg.Reports["mailgun"]; !ok {
		t.Fatal("Couldn't find the mailgun report")
	}
}

func TestStorageExist(t *testing.T) {
	if len(cfg.Storage["json"]) == 0 {
		t.Fatal("Couldn't find any storage")
	}
}

func TestStorage(t *testing.T) {
	if _, ok := cfg.Storage["json"]; !ok {
		t.Fatal("Couldn't find the json storage")
	}
}

func TestValidateStorageSettings(t *testing.T) {
	tmpcfg := cfg
	tmpcfg.Storage["json"]["path"] = "....."
	valid, err := validateStorageSettings(tmpcfg)
	if valid && (err != nil) {
		t.Errorf("validateStorageSettings should not be valid with an "+
			"error: %+v", err)
	}
	if valid {
		t.Error("Should have failed with no path")
	}
}
