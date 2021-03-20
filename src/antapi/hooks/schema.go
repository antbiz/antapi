package hooks

import "antapi/logic"

// RegisterSchemaHooks : 注册Schema的所有勾子
func RegisterSchemaHooks() {
	registerBeforeSaveHooks("schema", logic.DefaultSchemaLogic.CheckFields)
}
