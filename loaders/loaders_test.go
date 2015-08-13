package loaders

import (
	"testing"
)

var (
	LoaderT LoaderTextfile
)

func init() {
	LoaderT = LoaderTextfile{}
}

func TestTextfileFetchUrls(t *testing.T) {
	if _, err := LoaderT.FetchUrls(); err == nil {
		t.Error("FetchUrls should have failed as no path has been set yet")
	}

	textfile_path = "example.urls.txt"
	if urls, err := LoaderT.FetchUrls(); err != nil {
		t.Error("Shouldnt have returned an error - a valid path has been set")
	} else {
		if len(urls) != 1 {
			t.Error("Didn't get an expected number of results")
		} else {
			if urls[0] != "https://www.google.com/" {
				t.Error("Didn't get the expected result")
			}
		}
	}
}

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
