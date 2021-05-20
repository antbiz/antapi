package global

import (
	"context"
	"sync"

	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/os/glog"
)

// permission 常驻内存数据
var (
	permissionLocker sync.RWMutex
	permissionsMap   map[string]*dto.Permission
)

// LoadPermissions 将所有 权限 加载到内存
func LoadPermissions() error {
	permissions := ([]*dto.Permission)(nil)

	if err := db.DB().Collection("permission").Find(context.Background(), nil).All(&permissions); err != nil {
		glog.Error("LoadPermissions permission read fail:", err)
		return err
	}

	permissionLocker.Lock()
	defer permissionLocker.Unlock()

	permissionsMap = make(map[string]*dto.Permission)
	for _, perm := range permissions {
		permissionsMap[perm.CollectionName] = perm
	}

	glog.Info("LoadPermissions successfully!")
	return nil
}

// GetPermission 从内存中获取某个Collection的权限
func GetPermission(collectionName string) *dto.Permission {
	return permissionsMap[collectionName]
}
