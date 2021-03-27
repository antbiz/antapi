package hooks

// RegisterAllHooks : 注册全部的勾子
func RegisterAllHooks() {
	RegisterSchemaHooks()
	RegisterProjectHooks()
}
