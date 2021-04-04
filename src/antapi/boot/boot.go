package boot

import (
	_ "antapi/app/hooks"
	"antapi/cmd"
)

func init() {
	cmd.Sync()
	ServerBackground()
}
