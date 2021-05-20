package service

import (
	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/app/global"
	"github.com/gogf/gf/encoding/gjson"
)

var Permission = &permissionSrv{
	collectionName: "permission",
}

type permissionSrv struct {
	collectionName string
}

// CollectionName .
func (srv *permissionSrv) CollectionName() string {
	return srv.collectionName
}

// CheckDuplicatePermission 不允许创建 collection&role 名称均相同的权限
func (srv *permissionSrv) CheckDuplicatePermission(data *gjson.Json) error {
	return nil
}

// ReloadGlobalPermissions 重新加载所有 权限 到内存
func (srv *permissionSrv) ReloadGlobalPermissions(_ *gjson.Json) error {
	global.PermissionChan <- struct{}{}
	return nil
}

type canDo struct {
	PermissionVal int
}

func (c *canDo) CanNot() bool {
	return c.PermissionVal == 0
}

func (c *canDo) CanDoOnlyOwner() bool {
	return c.PermissionVal == 1
}

func (c *canDo) CanDoOnlyLogin() bool {
	return c.PermissionVal == 2
}

func (c *canDo) CanDo() bool {
	return c.PermissionVal == 3
}

// getPermission 从内存中获取权限等级。0-没有权限，1-仅创建者，2-仅登录者, 3-所有人
func (srv *permissionSrv) getPermission(permName dto.PermissionName, collectionName string) (*canDo, error) {
	perm := global.GetPermission(collectionName)
	return &canDo{PermissionVal: perm.GetPermissionLevel(permName)}, nil
}

// GetCreatePermission 增
func (srv *permissionSrv) GetCreatePermission(collectionName string) (*canDo, error) {
	return srv.getPermission(dto.CreateLevel, collectionName)
}

// GetReadPermission 查
func (srv *permissionSrv) GetReadPermission(collectionName string) (*canDo, error) {
	return srv.getPermission(dto.ReadLevel, collectionName)
}

// GetUpdatePermission 改
func (srv *permissionSrv) GetUpdatePermission(collectionName string) (*canDo, error) {
	return srv.getPermission(dto.UpdateLevel, collectionName)
}

// GetDeletePermission 删
func (srv *permissionSrv) GetDeletePermission(collectionName string) (*canDo, error) {
	return srv.getPermission(dto.DeleteLevel, collectionName)
}
