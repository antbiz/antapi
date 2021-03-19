package crud

import (
	"antapi/model"
	"errors"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/guid"
)

// UpdateOne : 更新单个数据
func UpdateOne(collectionName string, data interface{}) error {
	db := g.DB()
	oriObj, err := gjson.LoadJson(data)
	if err != nil {
		return err
	}
	id := oriObj.GetString("id")
	if len(id) == 0 {
		return errors.New("id is required")
	}
	schema, err := model.GetSchema(collectionName)
	if err != nil {
		return err
	}

	// 更新主体数据
	var content map[string]interface{}
	for _, field := range schema.GetPublicFields() {
		val := oriObj.Get(field.Name)
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
		tableRowsLen := len(oriObj.GetArray(field.Name))
		if tableRowsLen == 0 {
			continue
		}
		tableContent := make([]map[string]interface{}, 0)
		tableSchema, err := model.GetSchema(field.RelatedCollection)
		if err != nil {
			return err
		}

		for i := 0; i < tableRowsLen; i++ {
			oriObj.Set(fmt.Sprintf("%s.%d.pcn", field.Name, i), collectionName)
			oriObj.Set(fmt.Sprintf("%s.%d.idx", field.Name, i), i)
			oriObj.Set(fmt.Sprintf("%s.%d.pid", field.Name, i), id)
			oriObj.Set(fmt.Sprintf("%s.%d.pfd", field.Name, i), field.Name)
			if len(oriObj.GetString(fmt.Sprintf("%s.%d.%s", field.Name, i, "id"))) == 0 {
				oriObj.Set(fmt.Sprintf("%s.%d.id", field.Name, i), guid.S())
			}

			var tableRowContent map[string]interface{}
			for _, tableField := range tableSchema.GetPublicFields() {
				val := oriObj.Get(fmt.Sprintf("%s.%d.%s", field.Name, i, tableField.Name))
				if validErr := field.CheckFieldValue(val); validErr != nil {
					return validErr.Current()
				}
				tableRowContent[tableField.Name] = val
			}
			tableContent = append(tableContent, tableRowContent)
		}

		if _, err := db.Table(field.RelatedCollection).Save(tableContent); err != nil {
			return err
		}
	}

	return nil
}
