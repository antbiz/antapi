package model

import (
	schemahook "antapi/model/hooks/schema"

	"github.com/gogf/gf/encoding/gjson"
)

// 勾子不对子表支持，批量操作时个别勾子不对子表数据起作用
var (
	AfterFindHooks map[string][]func(data *gjson.Json) error

	BeforeSaveHooks map[string][]func(data *gjson.Json) error
	AfterSaveHooks  map[string][]func(data *gjson.Json) error

	BeforeUpdateHooks map[string][]func(data *gjson.Json) error
	AfterUpdateHooks  map[string][]func(data *gjson.Json) error

	BeforeInsertHooks map[string][]func(data *gjson.Json) error
	AfterInsertHooks  map[string][]func(data *gjson.Json) error

	BeforeDeleteHooks map[string][]func(data *gjson.Json) error
	AfterDeleteHooks  map[string][]func(data *gjson.Json) error
)

func registerAfterFindHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	AfterFindHooks[collectionName] = append(AfterFindHooks[collectionName], hook...)
}

func registerBeforeSaveHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	BeforeSaveHooks[collectionName] = append(BeforeSaveHooks[collectionName], hook...)
}

func registerAfterSaveHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	AfterSaveHooks[collectionName] = append(AfterSaveHooks[collectionName], hook...)
}

func registerBeforeUpdateHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	BeforeUpdateHooks[collectionName] = append(BeforeUpdateHooks[collectionName], hook...)
}

func registerAfterUpdateHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	AfterUpdateHooks[collectionName] = append(AfterUpdateHooks[collectionName], hook...)
}

func registerBeforeInsertHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	BeforeInsertHooks[collectionName] = append(BeforeInsertHooks[collectionName], hook...)
}

func registerAfterInsertHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	AfterInsertHooks[collectionName] = append(AfterInsertHooks[collectionName], hook...)
}

func registerBeforeDeleteHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	BeforeDeleteHooks[collectionName] = append(BeforeDeleteHooks[collectionName], hook...)
}

func registerAfterDeleteHooks(collectionName string, hook ...func(data *gjson.Json) error) {
	AfterDeleteHooks[collectionName] = append(AfterDeleteHooks[collectionName], hook...)
}

// RegisterAllHooks : 注册所有collection的勾子
func RegisterAllHooks() {
	registerBeforeSaveHooks("schema", schemahook.CheckFields)
}
