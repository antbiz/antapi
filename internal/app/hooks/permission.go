package hooks

import (
	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/app/service"
)

// registerPermissionHooks 注册Permission的所有勾子
func registerPermissionHooks() {
	collectionName := service.Permission.CollectionName()
	global.RegisterBeforeSaveHooks(
		collectionName,
		service.Permission.CheckDuplicatePermission,
	)
	global.RegisterAfterSaveHooks(
		collectionName,
		service.Permission.ReloadGlobalPermissions,
	)
	global.RegisterAfterDeleteHooks(
		collectionName,
		service.Permission.ReloadGlobalPermissions,
	)
}
