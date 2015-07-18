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
			if sd, err := IsUrlLockedDown(url); err == nil {
				r = append(r, sd)
			} else {
				log.Warnf("Failed to test %s for lockdown. Is the URL correct?", url)
			}
		}
	}
	return r, nil
}

// IsUrlLockedDown returns a SiteDefinition,err for a given url
func IsUrlLockedDown(url string) (defs.SiteDefinition, error) {
	// init return struct
	sd := defs.SiteDefinition{
		Url:            url,
		IsLockedDown:   false,
		HttpStatusCode: 0,
	}
	resp, err := http.Get(url)

	if err != nil {
		log.Debug(err)
		return sd, err
	}

	// HTTP 401, HTTP 403, HTTP 407, or User Service login
	if (resp.StatusCode == 401) ||  (resp.StatusCode == 403) ||  (resp.StatusCode == 407) ||  (resp.StatusCode == 550) || (len(resp.Header["X-Auto-Login"]) > 0) {
		sd.IsLockedDown = true
	}

	return sd, err
}
