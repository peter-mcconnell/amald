package config

import "testing"

var sampleConfigContent = []byte(`ignoreunchanged: false
templates:
  - type: html
    opts:
      dir: "./reports/tmpl/"
loaders:
  - type: testa
  - type: testb
  - type: gcloudcli
  - type: textfile
    opts: 
      path: "urls.txt"
intervals:
  - title: "yesterday"
    distancehours: 24
    color: "green"
  - title: "last week"
    distancehours: 168
    color: "red"
  - title: "last month"
    distancehours: 720
    color: "blue"
reports:
  - type: ascii
  - type: mailgun
    opts:
      privatekey: "api:something"
      domain: "something.com"
      from: "Lockdown Tracker <postmaster@something.com>"
      to: "Sysadmins <sysadmin@something.com>"
      subj: "Lockdown status report"
storage:
  - type: json
    opts:
      path: "../tmp/data.json" # folder must exist
      recordlimit: 100 # careful: this could truncate an existing file
`)

func TestParse(t *testing.T) {
	if err := parse(sampleConfigContent); err != nil {
		t.Errorf("Failed to parse config file data: %s", err)
	}

	conf := Conf
	if len(conf.Intervals) != 3 {
		t.Errorf("Expected 3 intervals, got %d", len(conf.Intervals))
	}

	if len(conf.Reports) != 2 {
		t.Errorf("Expected 2 reports, got %d", len(conf.Reports))
	}

	if len(conf.Loaders) != 4 {
		t.Errorf("Expected 4 loaders, got %d", len(conf.Loaders))
	}
}
