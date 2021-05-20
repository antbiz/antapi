package dto

// Permission .
type Permission struct {
	Title          string
	ProjectName    string
	CollectionName string
	CreateLevel    int
	ReadLevel      int
	UpdateLevel    int
	DeleteLevel    int
}

// PermissionName 权限等级字段
type PermissionName string

// 权限等级分为 4 个等级
// 0: 无权限
// 1: 仅创建者
// 2: 仅登录者
// 3: 所有人
const (
	CreateLevel PermissionName = "createLevel"
	ReadLevel   PermissionName = "readLevel"
	UpdateLevel PermissionName = "updateLevel"
	DeleteLevel PermissionName = "deleteLevel"
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
