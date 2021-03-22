package boot

import (
	"antapi/cmd"
	"antapi/hooks"
	_ "antapi/packed"
)

func init() {
	cmd.Sync()
	hooks.RegisterAllHooks()
	ServerBackground()
}
