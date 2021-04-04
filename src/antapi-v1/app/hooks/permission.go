package hooks

import (
	"antapi/app/global"
	"antapi/app/logic"
)

// registerPermissionHooks 注册Permission的所有勾子
func registerPermissionHooks() {
	collectionName := "permission"
	global.RegisterBeforeSaveHooks(
		collectionName,
		logic.Permission.CheckDuplicatePermission,
	)
	global.RegisterAfterSaveHooks(
		collectionName,
		logic.Permission.ReloadGlobalPermissions,
	)
	global.RegisterAfterDeleteHooks(
		collectionName,
		logic.Permission.ReloadGlobalPermissions,
	)
}
