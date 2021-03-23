package boot

import (
	"antapi/cmd"
	"antapi/app/hooks"
	_ "antapi/packed"
)

func init() {
	cmd.Sync()
	hooks.RegisterAllHooks()
	ServerBackground()
}
