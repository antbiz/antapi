package boot

import (
	"antapi/global"
	"antapi/logic"
)

// ServerBackground : 后台常驻任务
func ServerBackground() {
	go loadData()
}

// loadData : 将数据加载到内存，当数据变化时重新加载
func loadData() {
	logic.LoadSchemas()

	for {
		select {
		case <-global.SchemaChan:
			logic.LoadSchemas()
		}
	}
}
