package boot

import (
	"antapi/db"
	_ "antapi/packed"
)

func init() {
	setUpDatabase()
}

func setUpDatabase() {
	db.SyncCollections()
	db.SyncDefaultsData()
}
