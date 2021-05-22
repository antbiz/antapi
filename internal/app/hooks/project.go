package hooks

import (
	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/app/service"
)

// registerProjectHooks 注册Project的所有勾子
func registerProjectHooks() {
	collectionName := service.Project.CollectionName()
	global.RegisterAfterSaveHooks(
		collectionName,
		service.Project.AutoExportJSONFile,
	)
	global.RegisterAfterDeleteHooks(
		collectionName,
		service.Project.AutoDeleteJSONFile,
	)
}
