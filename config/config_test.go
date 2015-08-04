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
	cfg, err = Load("example.config.yaml")
	if err != nil {
		t.Fatalf("encountered error: %s", err)
	}
	if _, ok := cfg.Global["templatesdir"]; !ok {
		t.Fatal("templatesdir not set")
	}
}

func TestLoadDefaults(t *testing.T) {
	testcfg := defs.Config{}
	cfg, err := LoadDefaults(testcfg)
	if err == nil {
		if cfg.Reports == nil {
			t.Error("Reports not being set by LoadDefaults")
		}
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
	// ensure our data reflects what we've put in the example file
	if cfg.SummaryIntervals[0].Title != "yesterday" {
		t.Fatalf("SummaryIntervals[0].Title wasn't expected: %s", cfg.SummaryIntervals[0].Title)
	}
	if cfg.SummaryIntervals[0].DistanceHours != 24 {
		t.Fatalf("SummaryIntervals[0].DistanceHours wasn't expected: %s", cfg.SummaryIntervals[0].DistanceHours)
	}
	if cfg.SummaryIntervals[1].Title != "last week" {
		t.Fatalf("SummaryIntervals[1].Title wasn't expected: %s", cfg.SummaryIntervals[1].Title)
	}
	if cfg.SummaryIntervals[1].DistanceHours != 168 {
		t.Fatalf("SummaryIntervals[1].DistanceHours wasn't expected: %s", cfg.SummaryIntervals[1].DistanceHours)
	}
	if cfg.SummaryIntervals[2].Title != "last month" {
		t.Fatalf("SummaryIntervals[2].Title wasn't expected: %s", cfg.SummaryIntervals[2].Title)
	}
	if cfg.SummaryIntervals[2].DistanceHours != 720 {
		t.Fatalf("SummaryIntervals[2].DistanceHours wasn't expected: %s", cfg.SummaryIntervals[2].DistanceHours)
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
