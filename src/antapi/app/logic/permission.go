package logic

import (
	"antapi/app/dao"
	"antapi/app/global"
	"antapi/app/model"
	"antapi/common/errcode"
	"fmt"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
)

var Permission = &permissionLogic{
	collectionName: "permission",
}

type permissionLogic struct {
	collectionName string
}

// CheckDuplicatePermission 不允许创建 collection&role 名称均相同的权限
func (p *permissionLogic) CheckDuplicatePermission(data *gjson.Json) error {
	permID := data.GetString("id")
	permCollectionName := data.GetString("collection_name")
	permRoleName := data.GetString("role_name")
	arg := &dao.ExistsAndCountFuncArg{}

	if permID == "" {
		arg.Where = "collection_name=? AND role_name=?"
		arg.WhereArgs = []string{permCollectionName, permRoleName}
	} else {
		arg.Where = "id<>? collection_name=? AND role_name=?"
		arg.WhereArgs = []string{permID, permCollectionName, permRoleName}
	}

	if exists, err := dao.Exists(p.collectionName, arg); err != nil {
		return err
	} else if exists {
		return gerror.NewCodef(errcode.DuplicateError, errcode.DuplicateErrorMsg, fmt.Sprintf("%s 角色的权限", permRoleName))
	}
	return nil
}

// ReloadGlobalPermissions 重新加载所有 权限 到内存
func (permissionLogic) ReloadGlobalPermissions(_ *gjson.Json) error {
	global.PermissionChan <- struct{}{}
	return nil
}

type canDo struct {
	PermissionVal int
}

func (c *canDo) CanDo() bool {
	return c.PermissionVal > 0
}

func (c *canDo) CanNot() bool {
	return c.PermissionVal == 0
}

func (c *canDo) CanDoOnlyOwner() bool {
	return c.PermissionVal == 1
}

// getPermission 从内存中获取权限等级。0-没有权限，1-仅创建者，2-有完全的权限
func (p *permissionLogic) getPermission(permName model.PermissionName, collectionName string, roleNames ...string) (*canDo, error) {
	var (
		permissions = global.GetPermissions(collectionName)
		levelsArr   = garray.NewIntArray(true)
		rolesArr    = garray.NewStrArrayFrom(roleNames, true)
	)

	for _, perm := range permissions {
		if rolesArr.Contains(perm.RoleName) {
			levelsArr.Append(perm.GetPermissionLevel(permName))
		}
	}
	// 取最大权限值
	maxLevel, _ := levelsArr.Sort().PopRight()
	return &canDo{PermissionVal: maxLevel}, nil
}

// GetCreatePermission 增
func (p *permissionLogic) GetCreatePermission(collectionName string, roleNames ...string) (*canDo, error) {
	return p.getPermission(model.CreateLevel, collectionName, roleNames...)
}

// GetReadPermission 查
func (p *permissionLogic) GetReadPermission(collectionName string, roleNames ...string) (*canDo, error) {
	return p.getPermission(model.ReadLevel, collectionName, roleNames...)
}

// GetUpdatePermission 改
func (p *permissionLogic) GetUpdatePermission(collectionName string, roleNames ...string) (*canDo, error) {
	return p.getPermission(model.UpdateLevel, collectionName, roleNames...)
}

// GetDeletePermission 删
func (p *permissionLogic) GetDeletePermission(collectionName string, roleNames ...string) (*canDo, error) {
	return p.getPermission(model.DeleteLevel, collectionName, roleNames...)
}
