package logic

import (
	"antapi/app/dao"
	"fmt"

	"github.com/gogf/gf/frame/g"
)

var Permission = &permissionLogic{
	collectionName: "permission",
}

type permissionLogic struct {
	collectionName string
}

// TODO: 数据缓存
func (p *permissionLogic) canDo(action, collectionName string, roleNames ...string) (bool, error) {
	arg := &dao.ExistsAndCountFuncArg{
		Where:     fmt.Sprintf("%s=? AND collection_name=? AND role_name IN(?)", action),
		WhereArgs: g.Slice{true, roleNames},
	}
	return dao.Exists(collectionName, arg)
}

// CanCreate 增
func (p *permissionLogic) CanCreate(collectionName string, roleNames ...string) (bool, error) {
	return p.canDo("can_create", collectionName, roleNames...)
}

// CanRead 查
func (p *permissionLogic) CanRead(collectionName string, roleNames ...string) (bool, error) {
	return p.canDo("can_read", collectionName, roleNames...)
}

// CanUpdate 改
func (p *permissionLogic) CanUpdate(collectionName string, roleNames ...string) (bool, error) {
	return p.canDo("can_update", collectionName, roleNames...)
}

// CanDelete 删
func (p *permissionLogic) CanDelete(collectionName string, roleNames ...string) (bool, error) {
	return p.canDo("can_delete", collectionName, roleNames...)
}
