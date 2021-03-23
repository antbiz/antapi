package hooks

import "antapi/app/logic"

// RegisterSchemaHooks : 注册Schema的所有勾子
func RegisterSchemaHooks() {
	collectionName := "schema"
	registerBeforeSaveHooks(
		collectionName,
		logic.DefaultSchemaLogic.CheckFields,
	)
	registerAfterSaveHooks(
		collectionName,
		logic.DefaultSchemaLogic.MigrateCollectionSchema,
		logic.DefaultSchemaLogic.ReloadGlobalSchemas,
	)
	registerAfterDeleteHooks(
		collectionName,
		logic.DefaultSchemaLogic.ReloadGlobalSchemas,
	)
}
