package boot

import (
	"antapi/app/global"
)

// ServerBackground : 后台常驻任务
func ServerBackground() {
	go loadData()
}

// loadData : 将数据加载到内存，当数据变化时重新加载
func loadData() {
	global.LoadSchemas()

	for {
		select {
		case <-global.SchemaChan:
			global.LoadSchemas()
		}
	}
}
