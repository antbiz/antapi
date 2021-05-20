package hooks

import (
	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/app/service"
)

// registerSchemaHooks 注册Schema的所有勾子
func registerSchemaHooks() {
	collectionName := service.Schema.CollectionName()
	global.RegisterBeforeSaveHooks(
		collectionName,
		service.Schema.CheckJSONSchema,
	)
	global.RegisterAfterSaveHooks(
		collectionName,
		service.Schema.ReloadGlobalSchemas,
		service.Schema.AutoExportJSONFile,
		service.Schema.AutoCreateCollectionPermission,
	)
	global.RegisterAfterDeleteHooks(
		collectionName,
		service.Schema.ReloadGlobalSchemas,
		service.Schema.AutoDeleteJSONFile,
		service.Schema.AutoDeleteCollectionPermission,
	)
}
