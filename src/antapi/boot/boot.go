package boot

import (
	"antapi/hooks"
	"antapi/model"
	_ "antapi/packed"
)

func init() {
	model.Sync()
	hooks.RegisterAllHooks()
	ServerBackground()
}
