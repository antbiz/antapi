package model

var (
	AfterFindHooks map[string][]func(map[string]interface{}) error

	BeforeSaveHooks map[string][]func(map[string]interface{}) error
	AfterSaveHooks  map[string][]func(map[string]interface{}) error

	BeforeUpdateHooks map[string][]func(map[string]interface{}) error
	AfterUpdateHooks  map[string][]func(map[string]interface{}) error

	BeforeInsertHooks map[string][]func(map[string]interface{}) error
	AfterInsertHooks  map[string][]func(map[string]interface{}) error

	BeforeDeleteHooks    map[string][]func(map[string]interface{}) error
	AfterDeleteSaveHooks map[string][]func(map[string]interface{}) error
)

// RegisterAllHooks : 注册所有collection的勾子
func RegisterAllHooks() {
	registerAfterFindHooks()
	registerBeforeSaveHooks()
	registerAfterSaveHooks()
	registerBeforeUpdateHooks()
	registerAfterUpdateHooks()
	registerBeforeInsertHooks()
	registerAfterInsertHooks()
	registerBeforeDeleteHooks()
	registerAfterDeleteHooks()
}

func registerAfterFindHooks() {

}

func registerBeforeSaveHooks() {

}

func registerAfterSaveHooks() {

}

func registerBeforeUpdateHooks() {

}

func registerAfterUpdateHooks() {

}

func registerBeforeInsertHooks() {

}

func registerAfterInsertHooks() {

}

func registerBeforeDeleteHooks() {

}

func registerAfterDeleteHooks() {

}
