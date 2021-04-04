package model

// Permission .
type Permission struct {
	CollectionName string `orm:"collection_name"`
	CreateLevel    int    `orm:"create_level"`
	ReadLevel      int    `orm:"read_level"`
	UpdateLevel    int    `orm:"update_level"`
	DeleteLevel    int    `orm:"delete_level"`
}

// PermissionName 权限等级字段
type PermissionName string

// 权限等级分为 4 个等级
// 0: 无权限
// 1: 仅创建者
// 2: 仅登录者
// 3: 所有人
const (
	CreateLevel PermissionName = "create_level"
	ReadLevel   PermissionName = "read_level"
	UpdateLevel PermissionName = "update_level"
	DeleteLevel PermissionName = "delete_level"
)

// GetPermissionLevel 获取权限等级
func (p *Permission) GetPermissionLevel(name PermissionName) int {
	switch name {
	case CreateLevel:
		return p.CreateLevel
	case ReadLevel:
		return p.ReadLevel
	case UpdateLevel:
		return p.UpdateLevel
	case DeleteLevel:
		return p.DeleteLevel
	default:
		return 0
	}
}
