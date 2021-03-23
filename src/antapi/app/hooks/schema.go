package hooks

import (
	"antapi/app/global"
	"antapi/app/logic"
)

// RegisterSchemaHooks : 注册Schema的所有勾子
func RegisterSchemaHooks() {
	collectionName := "schema"
	global.RegisterBeforeSaveHooks(
		collectionName,
		logic.Schema.CheckFields,
	)
	global.RegisterAfterSaveHooks(
		collectionName,
		logic.Schema.MigrateCollectionSchema,
		logic.Schema.ReloadGlobalSchemas,
	)
	global.RegisterAfterDeleteHooks(
		collectionName,
		logic.Schema.ReloadGlobalSchemas,
	)
}
