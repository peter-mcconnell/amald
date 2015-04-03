package loaders

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"github.com/pemcconnell/amald/urltest"
	"os"
	"path/filepath"
)

type LoaderTextfile struct{}

var (
	textfile_path string
)

func textfileLoaderAvailable(settings map[string]string) bool {
	// does file exist?
	textfile_path, _ = filepath.Abs(settings["path"])
	log.Debugf("textfile: %s", textfile_path)
	_, err := os.Stat(textfile_path)
	if err != nil {
		log.Warnf("textfile file not found: %s", textfile_path)
		return false
	}
	return true

}

// ScanUrls calls some Gcloud CLI commands, parses the output & then checks
// the url using authtest
func (l *LoaderTextfile) FetchUrls() map[string]defs.SiteDefinition {
	m := map[string]defs.SiteDefinition{}
	file, err := os.Open(textfile_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		lockeddown, err := urltest.TestUrlIsLockedDown(url)
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Fatal(err)
		}
		m[url] = defs.SiteDefinition{Url: url, IsLockedDown: lockeddown}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}
