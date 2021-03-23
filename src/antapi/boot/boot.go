package boot

import (
	"antapi/cmd"
	_ "antapi/packed"
)

func init() {
	cmd.Sync()
	// hooks.RegisterAllHooks()
	ServerBackground()
}
