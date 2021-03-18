package crud

import (
	"antapi/model"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// InsertOne : 插入单个数据
// TODO: 返回插入数据的uuid
func InsertOne(collectionName string, data interface{}) error {
	return InsertList(collectionName, data)
}

// InsertList : 插入多个数据
// TODO: 需要考虑子表数据校验的提示信息
// TODO: 提前生成每个数据的ID，主体的ID需要填充到子表的PID上，从而完成批量插入
// TODO: 存在勾子的情况下应该循环插入，所以需要提前判断是批量插入还是循环插入
func InsertList(collectionName string, data ...interface{}) error {
	db := g.DB()
	schema, err := model.GetSchema(collectionName)
	if err != nil {
		return err
	}
	contents := make([]map[string]interface{}, 0)

	// 批量插入主体数据
	for i := 0; i < len(data); i++ {
		oriObj, err := gjson.LoadJson(data)
		if err != nil {
			return err
		}

		var content map[string]interface{}
		for _, field := range schema.GetPublicFields() {
			val := oriObj.Get(field.Name)
			if validErr := field.CheckFieldValue(val); validErr != nil {
				return validErr.Current()
			}
			content[field.Name] = val
		}

		contents = append(contents, content)
	}
	if _, err := db.Table(collectionName).Insert(contents); err != nil {
		return err
	}

	// 批量插入子表数据
	for _, field := range schema.GetTableFields() {
		tableContent := make([]map[string]interface{}, 0)
		for i := 0; i < len(data); i++ {
			oriObj, err := gjson.LoadJson(data)
			if err != nil {
				return err
			}
			tableRowsLen := len(oriObj.GetArray(field.Name))
			if tableRowsLen == 0 {
				continue
			}
			tableSchema, err := model.GetSchema(field.RelatedCollection)
			if err != nil {
				return err
			}

			for i := 0; i < tableRowsLen; i++ {
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
		}

		if _, err := db.Table(field.RelatedCollection).Insert(tableContent); err != nil {
			return err
		}
	}
	return nil
}
