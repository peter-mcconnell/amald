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
	Summaries defs.Summaries
	Cfg       defs.Config
}

// Send the message via mailgun
func (n *NotifierMailgun) Fire(config map[string]string) {
	log.Debug("Firing mailgun notifier")
	r := reports.ReportHTML{
		Templatepath: n.Cfg.Reports["mailgun"]["templatepath"],
	}
	if message, err := r.Generate(n.Summaries); err == nil {
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
