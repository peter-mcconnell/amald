package urltest

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

// TestUrlIsLockedDown checks a URL to see if it returns a 401 or has
// X-Auto-Login headers
func TestUrlIsLockedDown(url string) (bool, error) {

	lockeddown := false
	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Failed to get URL")
	}

	// HTTP 401, or User Service login
	if (resp.StatusCode == 401) || (len(resp.Header["X-Auto-Login"]) > 0) {
		lockeddown = true
	}

	log.Debug("url: ", url, ", lockeddown: ", lockeddown)

	return lockeddown, err
}
