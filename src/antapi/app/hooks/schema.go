package hooks

import "antapi/app/logic"

// RegisterSchemaHooks : 注册Schema的所有勾子
func RegisterSchemaHooks() {
	collectionName := "schema"
	registerBeforeSaveHooks(
		collectionName,
		logic.Schema.CheckFields,
	)
	registerAfterSaveHooks(
		collectionName,
		logic.Schema.MigrateCollectionSchema,
		logic.Schema.ReloadGlobalSchemas,
	)
	registerAfterDeleteHooks(
		collectionName,
		logic.Schema.ReloadGlobalSchemas,
	)
}
