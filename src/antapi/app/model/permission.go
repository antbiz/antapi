package model

// Permission .
type Permission struct {
	CollectionName string `orm:"collection_name"`
	RoleName       string `orm:"role_name"`
	CreateLevel    int    `orm:"create_level"`
	ReadLevel      int    `orm:"read_level"`
	UpdateLevel    int    `orm:"update_level"`
	DeleteLevel    int    `orm:"delete_level"`
}

type PermissionName string

const (
	CreateLevel PermissionName = "create_level"
	ReadLevel   PermissionName = "read_level"
	UpdateLevel PermissionName = "update_level"
	DeleteLevel PermissionName = "delete_level"
)

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
