package hooks

import (
	"antapi/app/global"
	"antapi/app/logic"
)

// registerSchemaHooks 注册Schema的所有勾子
func registerSchemaHooks() {
	collectionName := "schema"
	global.RegisterBeforeSaveHooks(
		collectionName,
		logic.Schema.CheckJSONSchema,
	)
	global.RegisterAfterSaveHooks(
		collectionName,
		logic.Schema.MigrateSchema,
		logic.Schema.ReloadGlobalSchemas,
		logic.Schema.AutoExportSchemaData,
	)
	global.RegisterAfterInsertHooks(
		collectionName,
		logic.Schema.AutoCreateCollectionPermission,
	)
	global.RegisterAfterDeleteHooks(
		collectionName,
		logic.Schema.ReloadGlobalSchemas,
		logic.Schema.AutoDeleteExportedJsonFile,
		logic.Schema.AutoDeleteCollectionPermission,
	)
}
