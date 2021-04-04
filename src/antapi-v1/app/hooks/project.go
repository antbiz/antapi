package hooks

import (
	"antapi/app/global"
	"antapi/app/logic"
)

// registerProjectHooks 注册Project的所有勾子
func registerProjectHooks() {
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
