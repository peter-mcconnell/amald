package defs

type Config struct {
	Loaders map[string]map[string]string `json:loaders,omitempty`
	Reports map[string]map[string]string `json:reports,omitempty`
	Storage map[string]map[string]string `json:storage,omitempty`
	Tests   map[string]bool
}

type SiteDefinition struct {
	Url          string `json:"url"`
	IsLockedDown bool   `json:"islockeddown"`
}

type Results struct {
	Timestamp string           `json:"timestamp"`
	Results   []SiteDefinition `json:"results"`
}

type Records struct {
	Records []Results `json:"records"`
}
