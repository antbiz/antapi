package boot

import (
	"antapi/model"
	_ "antapi/packed"
)

func init() {
	setUpDatabase()
}

func setUpDatabase() {
	model.SyncCollections()
	model.SyncDefaultsData()
}
