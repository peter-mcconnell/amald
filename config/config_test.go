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
	if _, ok := cfg.Reports["templates"]["path"]; !ok {
		t.Fatal("templates path not set")
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
