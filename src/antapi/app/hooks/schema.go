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
		logic.Schema.CheckFields,
	)
	global.RegisterAfterSaveHooks(
		collectionName,
		logic.Schema.MigrateCollectionSchema,
		logic.Schema.ReloadGlobalSchemas,
		logic.Schema.AutoExportSchemaData,
	)
	global.RegisterAfterDeleteHooks(
		collectionName,
		logic.Schema.ReloadGlobalSchemas,
		logic.Schema.AutoDeleteExportedJsonFile,
	)
}
