package crud

import (
	"antapi/hooks"
	"antapi/logic"
	"errors"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/guid"
)

// Update : 更新单个数据
func Update(collectionName string, data interface{}) error {
	db := g.DB()
	dataGJson, err := gjson.LoadJson(data)
	if err != nil {
		return err
	}
	id := dataGJson.GetString("id")
	if id == "" {
		return errors.New("id is required")
	}
	schema := logic.GetSchema(collectionName)

	// 执行 BeforeInsertHooks, BeforeSaveHooks 勾子
	for _, hook := range hooks.GetBeforeUpdateHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return err
		}
	}
	for _, hook := range hooks.GetBeforeSaveHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return err
		}
	}

	// 更新主体数据
	var content map[string]interface{}
	for _, field := range schema.GetPublicFields() {
		val := dataGJson.Get(field.Name)
		if validErr := field.CheckFieldValue(val); validErr != nil {
			return validErr.Current()
		}
		content[field.Name] = val
	}
	if _, err := db.Table(collectionName).Where("id", id).Update(content); err != nil {
		return err
	}

	// 更新子表数据
	for _, field := range schema.GetTableFields() {
		tableRowsLen := len(dataGJson.GetArray(field.Name))
		if tableRowsLen == 0 {
			continue
		}
		tableContent := make([]map[string]interface{}, 0)
		tableSchema := logic.GetSchema(field.RelatedCollection)

		tableIds := make([]string, tableRowsLen)
		for i := 0; i < tableRowsLen; i++ {
			dataGJson.Set(fmt.Sprintf("%s.%d.pcn", field.Name, i), collectionName)
			dataGJson.Set(fmt.Sprintf("%s.%d.idx", field.Name, i), i)
			dataGJson.Set(fmt.Sprintf("%s.%d.pid", field.Name, i), id)
			dataGJson.Set(fmt.Sprintf("%s.%d.pfd", field.Name, i), field.Name)
			tableRowId := dataGJson.GetString(fmt.Sprintf("%s.%d.%s", field.Name, i, "id"))
			if tableRowId == "" {
				tableRowId = guid.S()
				dataGJson.Set(fmt.Sprintf("%s.%d.id", field.Name, i), tableRowId)
			}
			tableIds = append(tableIds, tableRowId)

			var tableRowContent map[string]interface{}
			for _, tableField := range tableSchema.GetPublicFields() {
				val := dataGJson.Get(fmt.Sprintf("%s.%d.%s", field.Name, i, tableField.Name))
				if validErr := field.CheckFieldValue(val); validErr != nil {
					return validErr.Current()
				}
				tableRowContent[tableField.Name] = val
			}
			tableContent = append(tableContent, tableRowContent)
		}

		// 执行save操作，如果存在则更新，否则插入
		if _, err := db.Table(field.RelatedCollection).Save(tableContent); err != nil {
			return err
		}

		// 自动处理需要删除的子表数据
		if _, err := db.Table(field.RelatedCollection).
			Where("id not in (?)", tableIds).
			Where("pcn", collectionName).
			Where("pid", id).
			Where("pfd", field.Name).
			Delete(); err != nil {
			return err
		}
	}

	// 执行 AfterUpdateHooks, AfterSaveHooks 勾子
	for _, hook := range hooks.GetAfterUpdateHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return err
		}
	}
	for _, hook := range hooks.GetAfterSaveHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return err
		}
	}

	return nil
}
