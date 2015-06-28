package urltest

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"net/http"
)

// Batch will take a series of urls and check to see if they are locked down
func Batch(urls []string) ([]defs.SiteDefinition, error) {
	var r []defs.SiteDefinition
	if len(urls) != 0 {
		for _, url := range urls {
			if ld, err := IsUrlLockedDown(url); err == nil {
				r = append(r, defs.SiteDefinition{
					Url:          url,
					IsLockedDown: ld,
				})
			} else {
				log.Fatalf("TestUrlIsLockedDown failed: %s", err)
				return r, err
			}
		}
	}
	return r, nil
}

// IsUrlLockedDown checks a URL to see if it returns a 401 or has
// X-Auto-Login headers
func IsUrlLockedDown(url string) (bool, error) {

	lockeddown := false
	resp, err := http.Get(url)

	if err != nil {
		log.Warnf("Failed to get URL:\n%s", err)
		return lockeddown, err
	}

	// HTTP 401, or User Service login
	if (resp.StatusCode == 401) || (len(resp.Header["X-Auto-Login"]) > 0) {
		lockeddown = true
	}

	log.Debug("url: ", url, ", lockeddown: ", lockeddown)

	return lockeddown, err
}
