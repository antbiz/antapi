package logic

import (
	"antapi/app/dao"
	"fmt"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
)

var Permission = &permissionLogic{
	collectionName: "permission",
}

type permissionLogic struct {
	collectionName string
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

// TODO: 数据缓存
func (p *permissionLogic) getPermissionLevel(action, collectionName string, roleNames ...string) (*canDo, error) {
	arg := &dao.GetListFuncArg{
		Where:     fmt.Sprintf("%s=? AND collection_name=? AND role_name IN(?)", action),
		WhereArgs: g.Slice{true, collectionName, roleNames},
		Fields:    []string{action},
	}
	data, total, err := dao.GetList(p.collectionName, arg)
	if err != nil {
		return nil, err
	} else if data == nil {
		return &canDo{PermissionVal: 0}, nil
	}
	allLevels := garray.NewIntArray(true)
	for i := 0; i < total; i++ {
		allLevels.Append(data.GetInt(fmt.Sprintf("%d.%s", i, action)))
	}
	maxLevel, _ := allLevels.Sort().PopRight()
	return &canDo{PermissionVal: maxLevel}, nil
}

// GetCreatePermission 增
func (p *permissionLogic) GetCreatePermission(collectionName string, roleNames ...string) (*canDo, error) {
	return p.getPermissionLevel("can_create", collectionName, roleNames...)
}

// GetReadPermission 查
func (p *permissionLogic) GetReadPermission(collectionName string, roleNames ...string) (*canDo, error) {
	return p.getPermissionLevel("can_read", collectionName, roleNames...)
}

// GetUpdatePermission 改
func (p *permissionLogic) GetUpdatePermission(collectionName string, roleNames ...string) (*canDo, error) {
	return p.getPermissionLevel("can_update", collectionName, roleNames...)
}

// GetDeletePermission 删
func (p *permissionLogic) GetDeletePermission(collectionName string, roleNames ...string) (*canDo, error) {
	return p.getPermissionLevel("can_delete", collectionName, roleNames...)
}
