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
	allData := make([]map[string]interface{}, 0)

	// 批量插入主体数据
	for i := 0; i < len(data); i++ {
		oriObj, err := gjson.LoadJson(data)
		if err != nil {
			return err
		}

		var newObj map[string]interface{}
		for _, field := range schema.GetPublicFields() {
			val := oriObj.Get(field.Name)
			if validErr := field.CheckFieldValue(val); validErr != nil {
				return validErr.Current()
			}
			newObj[field.Name] = val
		}

		allData = append(allData, newObj)
	}
	if _, err := db.Table(collectionName).Insert(allData); err != nil {
		return err
	}

	// 批量插入子表数据
	for _, field := range schema.GetTableFields() {
		allTableData := make([]map[string]interface{}, 0)
		for i := 0; i < len(data); i++ {
			oriObj, err := gjson.LoadJson(data)
			if err != nil {
				return err
			}
			tableSchema, err := model.GetSchema(field.RelatedCollection)
			if err != nil {
				return err
			}
			childsLen := len(oriObj.GetArray(field.Name))
			if childsLen == 0 {
				continue
			}

			for i := 0; i < childsLen; i++ {
				var newChildObj map[string]interface{}
				for _, tableField := range tableSchema.GetPublicFields() {
					val := oriObj.Get(fmt.Sprintf("%s.%d.%s", field.Name, i, tableField.Name))
					if validErr := field.CheckFieldValue(val); validErr != nil {
						return validErr.Current()
					}
					newChildObj[tableField.Name] = val
				}
				allTableData = append(allTableData, newChildObj)
			}
		}

		if _, err := db.Table(field.RelatedCollection).Insert(allTableData); err != nil {
			return err
		}
	}
	return nil
}
