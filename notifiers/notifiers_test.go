package notifiers

import (
	"github.com/pemcconnell/amald/defs"
	"testing"
)

var (
	AsciiN          NotifierAscii
	MailgunN        NotifierMailgun
	TestScanResults []defs.SiteDefinition
)

func init() {
	TestScanResults = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "blah",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		defs.SiteDefinition{
			Url:            "bah",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		defs.SiteDefinition{
			Url:            "black",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		defs.SiteDefinition{
			Url:            "sheep",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
	}
	AsciiN = NotifierAscii{
		ScanResults: TestScanResults,
		TestMode:    true,
	}
	MailgunN = NotifierMailgun{
		ScanResults: TestScanResults,
		TestMode:    true,
	}
}

func TestAsciiFire(t *testing.T) {
	if err := AsciiN.Fire(); err != nil {
		t.Errorf("Got unexpected error: %s")
	}
}

func TestMailgunFire(t *testing.T) {
	if err := AsciiN.Fire(); err != nil {
		t.Errorf("Got unexpected error: %s")
	}
}

func TestNotifiers(t *testing.T) {
	cfg := defs.Config{
		Reports: map[string]map[string]string{
			"ascii": map[string]string{
				"a": "b",
			},
			"mailgun": map[string]string{
				"a": "b",
			},
		},
	}
	summaries := make(defs.Summaries, 0)
	fired := FireNotifiers(cfg, summaries, TestScanResults, true)
	if _, ok := fired["ascii"]; !ok {
		t.Error("FireNotifiers didnt fire ascii")
	}
	if _, ok := fired["mailgun"]; !ok {
		t.Error("FireNotifiers didnt fire mailgun")
	}
}
