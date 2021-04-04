package global

import (
	"antapi/app/model"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

// permission 常驻内存数据
var (
	permissionLocker sync.RWMutex
	permissionsMap   map[string]*model.Permission
)

// LoadPermissions 将所有 权限 加载到内存
func LoadPermissions() error {
	db := g.DB()
	permissions := ([]*model.Permission)(nil)

	if err := db.Table("permission").Structs(&permissions); err != nil {
		glog.Error("LoadPermissions permission read fail:", err)
		return err
	}

	permissionsMap = make(map[string]*model.Permission)
	for _, perm := range permissions {
		permissionsMap[perm.CollectionName] = perm
	}

	glog.Info("LoadPermissions successfully!")
	return nil
}

// GetPermission 从内存中获取某个Collection的权限
func GetPermission(collectionName string) *model.Permission {
	return permissionsMap[collectionName]
}
