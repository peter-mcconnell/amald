package loaders

import (
	"fmt"
	"strings"
	"testing"
)

func TestTextfileLoaderAvailable(t *testing.T) {
	var testfile = make(map[string]string)
	testfile["path"] = "./example.urls.txt" // target known file
	if !textfileLoaderAvailable(testfile) {
		t.Fatal("textfileLoaderAvailable failed to find known file")
	}
}

func TestGcloudCLISetup(t *testing.T) {
	if _, err := execGcloudComponentRequirements(); err != nil {
		t.Fatalf("gcloud components install failed: %s", err)
	}
}

func TestProjectCLIOutput(t *testing.T) {
	if _, err := execGcloudProjects(); err != nil {
		t.Fatalf("gcloud projects command failed: %s", err)
	}
}

func TestGcloudComponentsCmd(t *testing.T) {
	if _, err := execGcloudModules(""); err != nil {
		t.Fatalf("gcloud modules command failed: %s", err)
	}
}

func TestGcloudFetchUrls(t *testing.T) {
	var l = LoaderGcloudCLI{}
	if _, err := l.FetchUrls(); err != nil {
		t.Fatalf("FetchUrls failed: %s", err)
	}
}

func TestTextfileFetchUrls(t *testing.T) {
	var l = LoaderTextfile{}
	if _, err := l.FetchUrls(); err != nil {
		t.Fatalf("FetchUrls failed: %s", err)
	}
}

// Ensure the CLI returns in the expected format. If this changes, the
// regex will break
func ExampleProjectCLIOutput() {
	if out, err := execGcloudProjects(); err == nil {
		fline := strings.Split(out, "\n")
		fmt.Print(fline[0])
	}
	// Output:
	//PROJECT_ID         NAME               PROJECT_NUMBER
}

// Ensure the CLI returns in the expected format. If this changes, the
// regex will break
func ExampleGcloudModulesOutput() {
	if out, err := execGcloudModules(""); err == nil {
		fline := strings.Split(out, "\n")
		fmt.Print(fline[0])
	}
	// Output:
	//MODULE   VERSION  IS_DEFAULT
}

// Ensure the project list regex doesn't break
func ExampleProjectParseOutput() {
	ex_in := "PROJECT_ID         NAME               PROJECT_NUMBER\n" +
		"something  somethingelse  123456780981\n" +
		"else       blah           901284101982"
	fmt.Print(parseProjectsOutput(ex_in))
	// Output:
	//PROJECT_ID
	//something
	//else
}

// Ensure the modules regex doesn't break
func ExampleModulesParseOutput() {
	ex_in := "MODULE   VERSION  IS_DEFAULT\n" +
		"default  1        *"

	fmt.Print(parseModulesOutput(ex_in))
	// Output:
	//VERSION
	//1
}
