package storage

import (
	"github.com/pemcconnell/amald/defs"
)

// MergeData simply takes the scanResults and merges it into the existing data
func MergeData(scanResults []defs.SiteDefinition, olddata defs.Records) defs.Records {

	// add newdata to olddata
	merged := olddata
	merged.Records = append(merged.Records, defs.SiteDefinitionsToRecords(scanResults).Records...)

	return merged
}
