package loaders

import (
	"testing"
)

func TestTextfileLoaderAvailable(t *testing.T) {
	var testfile = make(map[string]string)
	testfile["path"] = "./example.urls.txt" // target known file
	if !textfileLoaderAvailable(testfile) {
		t.Fatal("textfileLoaderAvailable failed to find known file")
	}
}

func TestParseProjectsOutput(t *testing.T) {
	if parseProjectsOutput("test-name   other-thing  a") != "test-name" {
		t.Error("Was expecting a different output")
	}
}

func TestParseModulesOutput(t *testing.T) {
	if parseModulesOutput("test-name   other-thing  a") != "other-thing" {
		t.Error("Was expecting a different output")
	}
}
