package storage

import (
	"github.com/pemcconnell/amald/defs"
	"regexp"
	"testing"
)

var (
	scanResults []defs.SiteDefinition
)

func init() {
	scanResults = append(scanResults, defs.SiteDefinition{
		Url:          "https://google.com",
		IsLockedDown: false,
	}, defs.SiteDefinition{
		Url:          "https://test.com/",
		IsLockedDown: true,
	})
}

func TestFormData(t *testing.T) {
	if reg, err := regexp.Compile("\\{"); err == nil {
		if r, err := formData(scanResults); err != nil {
			t.Errorf("FormData failed: %s", err)
		} else if reg.FindString(r) == "" {
			t.Error("Failed to get expected format")
		}
	} else {
		t.Errorf("Problem creating the regex: %s", err)
	}
}
