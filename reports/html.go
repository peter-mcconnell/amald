package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"html/template"
)

// Generate creates an HTML string with all the required data
// in place
func (r *Report) GenerateHtml(summaries defs.Summaries) (string, error) {
	// format our data
	type TemplateData struct {
		Summaries      defs.Summaries
		Cfg            defs.Config
		StateKeysByInt map[int]string
	}
	tmpldata := TemplateData{
		Summaries:      summaries,
		Cfg:            r.Cfg,
		StateKeysByInt: map[int]string{},
	}
	for s, i := range defs.StateKeys {
		tmpldata.StateKeysByInt[i] = s
	}
	// now parse our template
	var doc bytes.Buffer
	dir := r.Cfg.Global["templatesdir"]
	if dir == "" {
		log.Fatal("global > templatesdir not set in config")
	}
	tmpl := template.Must(template.ParseGlob(dir + "email/*.html"))
	if err := tmpl.ExecuteTemplate(&doc, "email", tmpldata); err != nil {
		log.Fatalf("Error running ParseFiles: %s", err)
	}
	return doc.String(), nil
}
