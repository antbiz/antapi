package boot

import (
	_ "github.com/antbiz/antapi/internal/app/hooks"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/frame/g"
)

func init() {
	if err := db.Init(); err != nil {
		g.Log().Fatalf("init mongo failed: %v", err)
	}
	ServerBackground()
}
