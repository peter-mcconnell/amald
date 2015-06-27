package loaders

import (
	"fmt"
	"strings"
	"testing"
)

var testfile = "./example.urls.txt"

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

func ExampleProjectCLIOutput() {
	if out, err := execGcloudProjects(); err == nil {
		fline := strings.Split(out, "\n")
		fmt.Print(fline[0])
	}
	// Output:
	//PROJECT_ID         NAME               PROJECT_NUMBER
}

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

func TestTextfile(t *testing.T) {

}
