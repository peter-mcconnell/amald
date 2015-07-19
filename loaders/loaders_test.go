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
