package config

/*
Target holds information relating to a desired target for amald
*/
type Target struct {
	URL string `json:"url"`
}

/*
Storage stores information for all of amalds storage engine
*/
type Storage struct {
	Type string            `json:"type"`
	Opts map[string]string `json:"opts"`
}

/*
Template stores information for all of amalds template config
*/
type Template struct {
	Type string            `json:"type"`
	Opts map[string]string `json:"opts"`
}

/*
Loader tells amald how/where to load urls from
*/
type Loader struct {
	Type string            `json:"type"`
	Opts map[string]string `json:"opts"`
}

/*
Interval specifies a status update on amalds reports
*/
type Interval struct {
	Title         string `json:"title"`
	DistanceHours string `json:"distancehours"`
	Color         string `json:"color"`
}

/*
Report should detail the attributes of a given report type
*/
type Report struct {
	Type string            `json:"type"`
	Opts map[string]string `json:"opts"`
}

/*
Config contains the configuration loaded from the assigned config file during run
*/
type Config struct {
	IgnoreUnchanged bool       `json:"ignoreunchanged"`
	Loaders         []Loader   `json:"loaders"`
	Reports         []Report   `json:"reports"`
	Intervals       []Interval `json:"intervals"`
	Templates       []Template `json:"templates"`
	Storage         []Storage  `json:"storage"`
}
