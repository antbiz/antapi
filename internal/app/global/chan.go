package global

var (
	SchemaChan     = make(chan struct{}, 1)
	PermissionChan = make(chan struct{}, 1)
)
