package loaders

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"os/exec"
	"regexp"
	"strings"
)

type LoaderGcloudCLI struct{}

func gcloudcliLoaderAvailable() bool {

	_, err := execGcloudComponentRequirements()
	if err != nil {
		log.Fatalf("Unable to execute gcloud CLI projects cmd: %s", err)
		return false
	}

	return true
}

// execGcloudComponentRequirements Calls `gcloud components update app` and
// returns the output
func execGcloudComponentRequirements() (string, error) {

	log.Info("Attempting to install / update the gcloud app component")
	cmd := exec.Command("gcloud", "components", "update", "app", "alpha")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// execGcloudProjects Calls `gcloud preview projects list` and returns the
// output
func execGcloudProjects() (string, error) {
	cmd := exec.Command("gcloud", "alpha", "projects", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// Calls `gcloud preview app modules list` with a specified project and
// returns the output
func execGcloudModules(project string) (string, error) {
	cmd := exec.Command("gcloud", "preview", "app", "modules", "list", "--project", project)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// Parse the gcloud projects output to get the project names on their own
func parseProjectsOutput(data string) string {
	re := regexp.MustCompile("([^\\s]+)\\s.*")
	return re.ReplaceAllString(data, "$1")
}

// Parse the gcloud modules output to get the project names on their own
func parseModulesOutput(data string) string {
	re := regexp.MustCompile("[^\\s]+\\s+([^\\s]+)\\s.*")
	return re.ReplaceAllString(data, "$1")
}

// ScanUrls calls some Gcloud CLI commands, parses the output & then checks
// the url using authtest
func (l *LoaderGcloudCLI) FetchUrls() ([]string, error) {
	m := []string{}
	projectstring, err := execGcloudProjects()
	if err != nil {
		log.Fatalf("gcloud projects command failed: %s", err)
		return m, err
	}
	data := parseProjectsOutput(projectstring)
	projectsraw := strings.Split(data, "\n")
	projects := projectsraw[1 : len(projectsraw)-1]
	for _, project := range projects {
		if modules, err := execGcloudModules(project); err == nil {
			if modules == "" {
				log.Debugf("skipping FetchUrl loop for %s", project)
				continue
			}
			versionsraw := strings.Split(parseModulesOutput(modules), "\n")
			l := len(versionsraw)
			if l > 1 {
				versions := versionsraw[1 : len(versionsraw)-1]
				// versionscache ensures that we're only testing once per version
				// (gcloud modules can return multiple results per version)
				versionscache := make(map[string]bool)
				for _, version := range versions {
					if !versionscache[version] {
						m = append(m, "https://"+version+"-dot-"+project+".appspot.com")
						versionscache[version] = true
					}
				}
			}
		} else {
			// execGcloudModules failed
			log.Fatalf("execGcloudModules command failed for: %s", project)
			return m, err
		}
	}

	return m, nil
}
