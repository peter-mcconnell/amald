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
	Current []SiteDefinition            // []{Url, IsLockedDown}
	Summary map[string][]SiteDefinition // ["lastweek"][]{Url. IsLockedDown}
}

type JsonFormat struct {
	Data map[string][]SubJsonFormat
}

type SubJsonFormat struct {
	Meta map[string]string `json:"Meta"`
	Data []SiteDefinition  `json:"Data"`
}
