package notifiers

import (
	"github.com/pemcconnell/amald/defs"
	"testing"
)

var (
	AsciiN   NotifierAscii
	MailgunN NotifierMailgun
)

func init() {
	TestScanResults := []defs.SiteDefinition{
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
	}
	MailgunN = NotifierMailgun{}
}

func TestAsciiFire(t *testing.T) {
	if err := AsciiN.Fire(); err != nil {
		t.Errorf("Got unexpected error: %s")
	}
}
