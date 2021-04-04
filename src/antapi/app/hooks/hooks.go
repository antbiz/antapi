package hooks

// 注册全部的勾子
func init() {
	registerSchemaHooks()
	registerPermissionHooks()
	registerProjectHooks()
}
