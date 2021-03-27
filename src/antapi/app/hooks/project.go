package hooks

import (
	"antapi/app/global"
	"antapi/app/logic"
)

// RegisterProjectHooks : 注册Project的所有勾子
func RegisterProjectHooks() {
	collectionName := "project"
	global.RegisterAfterSaveHooks(
		collectionName,
		logic.Project.AutoExportProjectData,
	)
	global.RegisterAfterDeleteHooks(
		collectionName,
		logic.Project.AutoDeleteExportedJsonFile,
	)
}
