package loaders

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
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
	_, err := os.Stat(textfile_path)
	if err != nil {
		log.Warnf("textfile file not found: %s", textfile_path)
		return false
	}
	return true

}

// ScanUrls calls some Gcloud CLI commands, parses the output
func (l *LoaderTextfile) FetchUrls() []string {
	m := []string{}
	file, err := os.Open(textfile_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}
