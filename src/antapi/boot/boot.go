package boot

import (
	"antapi/app/hooks"
	"antapi/cmd"
	_ "antapi/packed"
)

func init() {
	hooks.RegisterAllHooks()
	cmd.Sync()
	ServerBackground()
}
