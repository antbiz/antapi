package global

import "github.com/gogf/gf/encoding/gjson"

// TODO: 目前为了避免循环引用只能开放hooks的注册方法，但是其实只允许初始化时执行，需要做优化
var (
	afterFindHooks = map[string][]func(data *gjson.Json) error{}

	beforeSaveHooks = map[string][]func(data *gjson.Json) error{}
	afterSaveHooks  = map[string][]func(data *gjson.Json) error{}

	beforeUpdateHooks = map[string][]func(data *gjson.Json) error{}
	afterUpdateHooks  = map[string][]func(data *gjson.Json) error{}

	beforeInsertHooks = map[string][]func(data *gjson.Json) error{}
	afterInsertHooks  = map[string][]func(data *gjson.Json) error{}

	beforeDeleteHooks = map[string][]func(data *gjson.Json) error{}
	afterDeleteHooks  = map[string][]func(data *gjson.Json) error{}
)

func RegisterAfterFindHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterFindHooks[collectionName] = append(afterFindHooks[collectionName], hook...)
}
func GetAfterFindHooks() map[string][]func(data *gjson.Json) error {
	return afterFindHooks
}
func GetAfterFindHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterFindHooks[collectionName]
}

func RegisterBeforeSaveHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	beforeSaveHooks[collectionName] = append(beforeSaveHooks[collectionName], hook...)
}
func GetBeforeSaveHooks() map[string][]func(data *gjson.Json) error {
	return beforeSaveHooks
}
func GetBeforeSaveHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return beforeSaveHooks[collectionName]
}

func RegisterAfterSaveHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterSaveHooks[collectionName] = append(afterSaveHooks[collectionName], hook...)
}
func GetAfterSaveHooks() map[string][]func(data *gjson.Json) error {
	return afterSaveHooks
}
func GetAfterSaveHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterSaveHooks[collectionName]
}

func RegisterBeforeUpdateHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	beforeUpdateHooks[collectionName] = append(beforeUpdateHooks[collectionName], hook...)
}
func GetBeforeUpdateHooks() map[string][]func(data *gjson.Json) error {
	return beforeUpdateHooks
}
func GetBeforeUpdateHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return beforeUpdateHooks[collectionName]
}

func RegisterAfterUpdateHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterUpdateHooks[collectionName] = append(afterUpdateHooks[collectionName], hook...)
}
func GetAfterUpdateHooks() map[string][]func(data *gjson.Json) error {
	return afterUpdateHooks
}
func GetAfterUpdateHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterUpdateHooks[collectionName]
}

func RegisterBeforeInsertHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	beforeInsertHooks[collectionName] = append(beforeInsertHooks[collectionName], hook...)
}
func GetBeforeInsertHooks() map[string][]func(data *gjson.Json) error {
	return beforeInsertHooks
}
func GetBeforeInsertHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return beforeInsertHooks[collectionName]
}

func RegisterAfterInsertHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterInsertHooks[collectionName] = append(afterInsertHooks[collectionName], hook...)
}
func GetAfterInsertHooks() map[string][]func(data *gjson.Json) error {
	return afterInsertHooks
}
func GetAfterInsertHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterInsertHooks[collectionName]
}

func RegisterBeforeDeleteHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	beforeDeleteHooks[collectionName] = append(beforeDeleteHooks[collectionName], hook...)
}
func GetBeforeDeleteHooks() map[string][]func(data *gjson.Json) error {
	return beforeDeleteHooks
}
func GetBeforeDeleteHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return beforeDeleteHooks[collectionName]
}

func RegisterAfterDeleteHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterDeleteHooks[collectionName] = append(afterDeleteHooks[collectionName], hook...)
}
func GetAfterDeleteHooks() map[string][]func(data *gjson.Json) error {
	return afterDeleteHooks
}
func GetAfterDeleteHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterDeleteHooks[collectionName]
}
