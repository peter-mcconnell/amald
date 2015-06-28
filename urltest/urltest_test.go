package urltest

import (
	"testing"
)

func TestIsUrlLockedDown(t *testing.T) {
	if isLockedDown, err := IsUrlLockedDown("https://www.google.com"); err == nil {
		if isLockedDown == true {
			t.Error("Reporting that google is locked down? We got a problem")
		}
	} else {
		t.Errorf("IsUrlLockedDown failed: %s", err)
	}
}

func TestBatch(t *testing.T) {
	urls := []string{"https://www.google.com", "https://bbc.co.uk"}
	if sites, err := Batch(urls); err == nil {
		// verify response
		if len(sites) != 2 {
			t.Errorf("Didnt return expected number of results: %s\n%s", len(sites), sites)
		}
		if sites[0].Url != urls[0] {
			t.Errorf("[0] URL appears to have changed: %s, %s", sites[0].Url, urls[0])
		}
		if sites[0].IsLockedDown == true {
			t.Errorf("Reporting that google is locked down? we got a problem. Tested %s", sites[0].Url)
		}
	} else {
		t.Errorf("Failed to batch urls: %s", err)
	}
}
