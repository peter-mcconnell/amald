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
	TestMode    bool
}

// Fire the Mailgun report (HTML email)
func (n *NotifierMailgun) Fire() error {
	log.Debug("Firing mailgun notifier")
	r := reports.Report{
		Cfg:         n.Cfg,
		ScanResults: n.ScanResults,
	}
	if message, err := r.GenerateHtml(n.Summaries); err != nil {
		log.Error(err)
	} else {
		config := n.Cfg.Reports["mailgun"]
		client := &http.Client{}
		to := config["to"]
		if n.TestMode {
			to += "&o:testmode=true"
		}
		data := url.Values{}
		data.Add("from", config["from"])
		data.Add("to", to)
		data.Add("subject", config["subj"])
		data.Add("html", message)
		req, err := http.NewRequest(
			"POST",
			config["domain"]+"/messages",
			strings.NewReader(data.Encode()))
		req.SetBasicAuth("api", config["privatekey"])
		req.Header.Set("Content-Type",
			"application/x-www-form-urlencoded")
		if err != nil {
			log.Error(err)
			return err
		}
		if resp, err := client.Do(req); err != nil {
			log.Error(err)
			return err
		} else if resp.StatusCode != 200 {
			log.Error(resp)
			return err
		} else {
			log.Info("Mailgun notification sent!")
		}
	}
	return nil
}
