package reports

type DataDiff struct {
	AddedApps   map[string]bool
	RemovedApps map[string]bool
	UpdatedApps map[string]bool
}

type ReportLoader interface {
	Generate() (string, error)
}
