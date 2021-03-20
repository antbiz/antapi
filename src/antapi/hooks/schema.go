package hooks

import "antapi/logic"

// RegisterSchemaHooks : 注册Schema的所有勾子
func RegisterSchemaHooks() {
	collectionName := "schema"
	registerBeforeSaveHooks(collectionName, logic.DefaultSchemaLogic.CheckFields)
	registerAfterSaveHooks(collectionName, logic.DefaultSchemaLogic.ReloadGlobalSchemas)
	registerAfterDeleteHooks(collectionName, logic.DefaultSchemaLogic.ReloadGlobalSchemas)
}
