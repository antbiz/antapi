package boot

import (
	_ "antapi/app/hooks"
	"antapi/cmd"
	_ "antapi/packed"
)

func init() {
	cmd.Sync()
	ServerBackground()
}
