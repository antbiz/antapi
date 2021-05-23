package service

import (
	"context"

	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"go.mongodb.org/mongo-driver/bson"
)

var Permission = &permissionSrv{}

type permissionSrv struct{}

// CollectionName .
func (srv *permissionSrv) CollectionName() string {
	return "permission"
}

// CheckDuplicatePermission 不允许创建 collection&role 名称均相同的权限
func (srv *permissionSrv) CheckDuplicatePermission(ctx context.Context, data *gjson.Json) error {
	collectionName := data.GetString("collectionName")
	roleName := data.GetString("roleName")
	if collectionName == "" || roleName == "" {
		return nil
	}
	id := data.GetString("_id")

	filter := bson.M{}
	if id == "" {
		filter["collectionName"] = collectionName
	} else {
		filter["collectionName"] = collectionName
		filter["$ne"] = bson.M{
			"_id": id,
		}
	}

	if total, err := db.DB().Collection(srv.CollectionName()).Find(context.Background(), filter).Count(); err != nil {
		return err
	} else if total > 0 {
		return gerror.New("Duplicate Error")
	}

	return nil
}

// ReloadGlobalPermissions 重新加载所有 权限 到内存
func (srv *permissionSrv) ReloadGlobalPermissions(ctx context.Context, data *gjson.Json) error {
	global.PermissionChan <- struct{}{}
	return nil
}

type canDo struct {
	PermissionVal int
}

func (c *canDo) CanDoOnlySysUser() bool {
	return c.PermissionVal == 0
}

func (c *canDo) CanDoOnlyOwner() bool {
	return c.PermissionVal == 1
}

func (c *canDo) CanDoOnlyLogin() bool {
	return c.PermissionVal == 2
}

func (c *canDo) CanDoAll() bool {
	return c.PermissionVal == 3
}

// getPermission 从内存中获取权限等级。0-没有权限，1-仅创建者，2-仅登录者, 3-所有人
func (srv *permissionSrv) getPermission(permName string, collectionName string) *canDo {
	perm := global.GetPermission(collectionName)
	if perm == nil {
		return &canDo{}
	}
	return &canDo{PermissionVal: perm.GetPermissionLevel(permName)}
}

// GetCreatePermission 增
func (srv *permissionSrv) GetCreatePermission(collectionName string) *canDo {
	return srv.getPermission(dto.CreateLevel, collectionName)
}

// GetReadPermission 查
func (srv *permissionSrv) GetReadPermission(collectionName string) *canDo {
	return srv.getPermission(dto.ReadLevel, collectionName)
}

// GetUpdatePermission 改
func (srv *permissionSrv) GetUpdatePermission(collectionName string) *canDo {
	return srv.getPermission(dto.UpdateLevel, collectionName)
}

// GetDeletePermission 删
func (srv *permissionSrv) GetDeletePermission(collectionName string) *canDo {
	return srv.getPermission(dto.DeleteLevel, collectionName)
}
