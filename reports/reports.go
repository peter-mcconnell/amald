package reports

import (
	"github.com/pemcconnell/amald/defs"
)

type Report struct {
	Cfg               defs.Config
	AnsiColorEnabled bool
}

func (r *Report) SetCfg(cfg defs.Config) {
	r.Cfg = cfg
}
