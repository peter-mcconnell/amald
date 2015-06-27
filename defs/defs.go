package defs

type Config struct {
	Loaders map[string]map[string]string `json:loaders,omitempty`
	Reports map[string]map[string]string `json:reports,omitempty`
	Storage map[string]map[string]string `json:storage,omitempty`
}

type SiteDefinition struct {
	Url          string `json:"url"`
	IsLockedDown bool   `json:"islockeddown"`
}

type JsonData struct {
	Meta map[string]string         `json:"Meta"`
	Data map[string]SiteDefinition `json:"Data"`
}
