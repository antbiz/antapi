package model

import (
	schemahook "antapi/model/hooks/schema"

	"github.com/gogf/gf/encoding/gjson"
)

// 勾子不对子表支持，批量操作时个别勾子不对子表数据起作用
var (
	afterFindHooks map[string][]func(data *gjson.Json) error

	beforeSaveHooks map[string][]func(data *gjson.Json) error
	afterSaveHooks  map[string][]func(data *gjson.Json) error

	beforeUpdateHooks map[string][]func(data *gjson.Json) error
	afterUpdateHooks  map[string][]func(data *gjson.Json) error

	beforeInsertHooks map[string][]func(data *gjson.Json) error
	afterInsertHooks  map[string][]func(data *gjson.Json) error

	// beforeDeleteHooks 不关心子表数据
	beforeDeleteHooks map[string][]func(data *gjson.Json) error
	// afterDeleteHooks 不关心子表数据
	afterDeleteHooks map[string][]func(data *gjson.Json) error
)

func registerAfterFindHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterFindHooks[collectionName] = append(afterFindHooks[collectionName], hook...)
}
func GetAfterFindHooks() map[string][]func(data *gjson.Json) error {
	return afterFindHooks
}
func GetAfterFindHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterFindHooks[collectionName]
}

func registerBeforeSaveHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	beforeSaveHooks[collectionName] = append(beforeSaveHooks[collectionName], hook...)
}
func GetBeforeSaveHooks() map[string][]func(data *gjson.Json) error {
	return beforeSaveHooks
}
func GetBeforeSaveHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return beforeSaveHooks[collectionName]
}

func registerAfterSaveHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterSaveHooks[collectionName] = append(afterSaveHooks[collectionName], hook...)
}
func GetAfterSaveHooks() map[string][]func(data *gjson.Json) error {
	return afterSaveHooks
}
func GetAfterSaveHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterSaveHooks[collectionName]
}

func registerBeforeUpdateHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	beforeUpdateHooks[collectionName] = append(beforeUpdateHooks[collectionName], hook...)
}
func GetBeforeUpdateHooks() map[string][]func(data *gjson.Json) error {
	return beforeUpdateHooks
}
func GetBeforeUpdateHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return beforeUpdateHooks[collectionName]
}

func registerAfterUpdateHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterUpdateHooks[collectionName] = append(afterUpdateHooks[collectionName], hook...)
}
func GetAfterUpdateHooks() map[string][]func(data *gjson.Json) error {
	return afterUpdateHooks
}
func GetAfterUpdateHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterUpdateHooks[collectionName]
}

func registerBeforeInsertHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	beforeInsertHooks[collectionName] = append(beforeInsertHooks[collectionName], hook...)
}
func GetBeforeInsertHooks() map[string][]func(data *gjson.Json) error {
	return beforeInsertHooks
}
func GetBeforeInsertHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return beforeInsertHooks[collectionName]
}

func registerAfterInsertHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterInsertHooks[collectionName] = append(afterInsertHooks[collectionName], hook...)
}
func GetAfterInsertHooks() map[string][]func(data *gjson.Json) error {
	return afterInsertHooks
}
func GetAfterInsertHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterInsertHooks[collectionName]
}

func registerBeforeDeleteHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	beforeDeleteHooks[collectionName] = append(beforeDeleteHooks[collectionName], hook...)
}
func GetBeforeDeleteHooks() map[string][]func(data *gjson.Json) error {
	return beforeDeleteHooks
}
func GetBeforeDeleteHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return beforeDeleteHooks[collectionName]
}

func registerAfterDeleteHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	afterDeleteHooks[collectionName] = append(afterDeleteHooks[collectionName], hook...)
}
func GetAfterDeleteHooks() map[string][]func(data *gjson.Json) error {
	return afterDeleteHooks
}
func GetAfterDeleteHooksByCollectionName(collectionName string) []func(data *gjson.Json) error {
	return afterDeleteHooks[collectionName]
}

// RegisterAllHooks : 注册所有collection的勾子
func RegisterAllHooks() {
	registerBeforeSaveHooks("schema", schemahook.CheckFields)
}
