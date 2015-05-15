// TODO: create a proper template for these. The strings are nasty
package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"html/template"
)

type ReportHTML struct {
	Templatepath string
}

// GenerateHTML creates an HTML string with all the required data
// in place
func (r *ReportHTML) Generate(data []defs.JsonData) (string, error) {
	diff, err := FindDiffs(data)
	if err != nil {
		log.Errorf("failed to diffData: %s", err)
	}
	type TemplateData struct {
		Summary map[string]DataDiff
		List    defs.JsonData
	}
	tmpldata := TemplateData{}
	tmpldata.Summary = diff
	tmpldata.List = data[0]
	var doc bytes.Buffer
	tmpl := template.Must(template.ParseGlob(r.Templatepath + "email/*"))
	err = tmpl.ExecuteTemplate(&doc, "email", tmpldata)
	if err != nil {
		log.Fatalf("Error running ParseFiles: %s", err)
	}
	return doc.String(), nil
}
