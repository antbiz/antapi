package boot

import (
	"antapi/model"
	_ "antapi/packed"
)

func init() {
	model.Sync()
	model.RegisterAllHooks()
}
