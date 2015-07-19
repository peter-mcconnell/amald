package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"github.com/pemcconnell/amald/reports"
	"net/http"
	"net/url"
	"strings"
)

type NotifierMailgun struct {
	ScanResults []defs.SiteDefinition
	Summaries   defs.Summaries
	Cfg         defs.Config
}

// Send the message via mailgun
func (n *NotifierMailgun) Fire() {
	log.Debug("Firing mailgun notifier")
	r := reports.Report{
		Cfg:         n.Cfg,
		ScanResults: n.ScanResults,
	}
	if message, err := r.GenerateHtml(n.Summaries); err == nil {
		config := n.Cfg.Reports["mailgun"]
		client := &http.Client{}
		data := url.Values{}
		data.Add("from", config["from"])
		data.Add("to", config["to"])
		data.Add("subject", config["subj"])
		data.Add("html", message)
		req, err := http.NewRequest("POST", config["domain"]+"/messages", strings.NewReader(data.Encode()))
		req.SetBasicAuth("api", config["privatekey"])
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != 200 {
			log.Fatal(resp)
		} else {
			log.Info("Mailgun notification sent!")
		}
	}
}
