package logic

import (
	"antapi/app/dao"
	"antapi/app/global"
	"antapi/app/model"
	"antapi/common/errcode"
	"fmt"

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
	permCollectionName := data.GetString("collection_name")
	permRoleName := data.GetString("role_name")
	if permCollectionName == "" || permRoleName == "" {
		return nil
	}
	permID := data.GetString("id")
	arg := &dao.ExistsAndCountFuncArg{}

	if permID == "" {
		arg.Where = "collection_name=?"
		arg.WhereArgs = []string{permCollectionName}
	} else {
		arg.Where = "id<>? collection_name=?"
		arg.WhereArgs = []string{permID, permCollectionName}
	}

	if exists, err := dao.Exists(p.collectionName, arg); err != nil {
		return err
	} else if exists {
		return gerror.NewCodef(errcode.DuplicateError, errcode.DuplicateErrorMsg, fmt.Sprintf("%s 模型的权限", permCollectionName))
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
func (p *permissionLogic) getPermission(permName model.PermissionName, collectionName string) (*canDo, error) {
	perm := global.GetPermission(collectionName)
	return &canDo{PermissionVal: perm.GetPermissionLevel(permName)}, nil
}

// GetCreatePermission 增
func (p *permissionLogic) GetCreatePermission(collectionName string) (*canDo, error) {
	return p.getPermission(model.CreateLevel, collectionName)
}

// GetReadPermission 查
func (p *permissionLogic) GetReadPermission(collectionName string) (*canDo, error) {
	return p.getPermission(model.ReadLevel, collectionName)
}

// GetUpdatePermission 改
func (p *permissionLogic) GetUpdatePermission(collectionName string) (*canDo, error) {
	return p.getPermission(model.UpdateLevel, collectionName)
}

// GetDeletePermission 删
func (p *permissionLogic) GetDeletePermission(collectionName string) (*canDo, error) {
	return p.getPermission(model.DeleteLevel, collectionName)
}
