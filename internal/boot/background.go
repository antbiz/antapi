package boot

import "github.com/antbiz/antapi/internal/app/global"

// ServerBackground : 后台常驻任务
func ServerBackground() {
	go loadData()
}

// loadData : 将数据加载到内存，当数据变化时重新加载
func loadData() {
	global.LoadSchemas()
	global.LoadPermissions()

	for {
		select {
		case <-global.SchemaChan:
			global.LoadSchemas()
		case <-global.PermissionChan:
			global.LoadPermissions()
		}
	}
}
