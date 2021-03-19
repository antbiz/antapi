package crud

import (
	"antapi/model"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/guid"
)

// InsertOne : 插入单个数据，返回插入的主体id
func InsertOne(collectionName string, data interface{}) (string, error) {
	if res, err := InsertList(collectionName, data); err != nil {
		return "", nil
	} else {
		return res[0], nil
	}
}

// InsertList : 插入多个数据，返回一组插入的主体id
// TODO: 需要考虑子表数据校验的提示信息
// TODO: 存在勾子的情况下应该循环插入，所以需要提前判断是批量插入还是循环插入
func InsertList(collectionName string, data ...interface{}) ([]string, error) {
	dataLen := len(data)
	if dataLen == 0 {
		return nil, nil
	}
	db := g.DB()
	schema, err := model.GetSchema(collectionName)
	if err != nil {
		return nil, nil
	}
	ids := make([]string, 0, dataLen)
	contents := make([]map[string]interface{}, 0, dataLen)

	// 批量插入主体数据
	for i := 0; i < dataLen; i++ {
		oriObj, err := gjson.LoadJson(data)
		if err != nil {
			return nil, nil
		}
		id := guid.S()
		oriObj.Set("id", id)
		ids = append(ids, id)

		var content map[string]interface{}
		for _, field := range schema.GetPublicFields() {
			val := oriObj.Get(field.Name)
			if validErr := field.CheckFieldValue(val); validErr != nil {
				return nil, validErr.Current()
			}
			content[field.Name] = val
		}

		contents = append(contents, content)
	}
	if _, err := db.Table(collectionName).Insert(contents); err != nil {
		return nil, err
	}

	// 批量插入子表数据
	for _, field := range schema.GetTableFields() {
		tableContent := make([]map[string]interface{}, 0)
		for i := 0; i < dataLen; i++ {
			oriObj, err := gjson.LoadJson(data)
			if err != nil {
				return nil, err
			}
			tableRowsLen := len(oriObj.GetArray(field.Name))
			if tableRowsLen == 0 {
				continue
			}
			tableSchema, err := model.GetSchema(field.RelatedCollection)
			if err != nil {
				return nil, err
			}

			for j := 0; j < tableRowsLen; j++ {
				var tableRowContent map[string]interface{}
				for _, tableField := range tableSchema.GetPublicFields() {
					val := oriObj.Get(fmt.Sprintf("%s.%d.%s", field.Name, j, tableField.Name))
					if validErr := field.CheckFieldValue(val); validErr != nil {
						return nil, validErr.Current()
					}
					tableRowContent[tableField.Name] = val
				}
				tableRowContent["pcn"] = collectionName
				tableRowContent["id"] = guid.S()
				tableRowContent["idx"] = j
				tableRowContent["pid"] = ids[i]
				tableRowContent["pfd"] = field.Name
				tableContent = append(tableContent, tableRowContent)
			}
		}

		if _, err := db.Table(field.RelatedCollection).Insert(tableContent); err != nil {
			return nil, err
		}
	}
	return ids, nil
}
