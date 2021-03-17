package boot

import (
	"antapi/core/schema"
	_ "antapi/packed"
)

func init() {
	setUpDatabase()
}

func setUpDatabase() {
	schema.SyncCollections()
	schema.SyncDefaultsData()
}
