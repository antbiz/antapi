package crud

import (
	"antapi/model"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// InsertOne : 插入单个数据
// TODO: 返回插入数据的uuid
func InsertOne(collectionName string, data interface{}) error {
	return InsertList(collectionName, data)
}

// InsertList : 插入多个数据
func InsertList(collectionName string, data ...interface{}) error {
	db := g.DB()
	allData := make([]map[string]interface{}, 0)
	for i := 0; i < len(data); i++ {
		oriObj, err := gjson.LoadJson(data)
		if err != nil {
			return err
		}
		schema, err := model.GetSchema(collectionName)
		if err != nil {
			return err
		}

		var newObj map[string]interface{}
		for _, field := range schema.Fields {
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
	return nil
}
