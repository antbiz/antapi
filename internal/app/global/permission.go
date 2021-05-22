package global

import (
	"context"
	"sync"

	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/frame/g"
	"go.mongodb.org/mongo-driver/bson"
)

// permission 常驻内存数据
var (
	permissionLocker sync.RWMutex
	permissionsMap   map[string]*dto.Permission
)

// LoadPermissions 将所有 权限 加载到内存
func LoadPermissions() {
	permissions := ([]*dto.Permission)(nil)
	if err := db.DB().Collection("permission").Find(context.Background(), bson.M{}).All(&permissions); err != nil {
		g.Log().Errorf("LoadPermissions permission read fail: %v", err)
		return
	}

	permissionLocker.Lock()
	defer permissionLocker.Unlock()

	permissionsMap = make(map[string]*dto.Permission)
	for _, perm := range permissions {
		permissionsMap[perm.CollectionName] = perm
	}

	g.Log().Debug("LoadPermissions successfully!")
}

// GetPermission 从内存中获取某个Collection的权限
func GetPermission(collectionName string) *dto.Permission {
	return permissionsMap[collectionName]
}
