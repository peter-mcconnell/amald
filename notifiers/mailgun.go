package notifiers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"github.com/pemcconnell/amald/reports"
	"os/exec"
)

type NotifierMailgun struct {
	data []defs.JsonData
}

// Send the message via mailgun
func (n *NotifierMailgun) Send(config map[string]string) {
	r := reports.ReportHTML{}
	if message, err := r.Generate(n.data); err == nil {
		out, err := exec.Command("curl", "-s", "--user", config["privatekey"], config["domain"]+"/messages", "-F", "from="+config["from"], "-F", "to="+config["to"], "-F", "subject="+config["subj"], "-F", "html="+message).Output()
		if err != nil {
			log.Fatalf("failed to send using mailgun:%s\n%s", out, err)
		}
		log.Debugf("\nmailgun response:\n%s", out)
	}
}
