package reports

type ReportLoader interface {
	Generate() (string, error)
}
