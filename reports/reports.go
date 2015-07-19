package reports

import (
	"github.com/pemcconnell/amald/defs"
)

type Report struct {
	Cfg              defs.Config
	ScanResults      []defs.SiteDefinition
	AnsiColorEnabled bool
}
