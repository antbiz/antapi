package crud

import (
	"antapi/model"
	"errors"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
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

	var newObj map[string]interface{}
	for _, field := range schema.Fields {
		val := oriObj.Get(field.Name)
		if validErr := field.CheckFieldValue(val); validErr != nil {
			return validErr.Current()
		}
		newObj[field.Name] = val
	}

	if _, err := db.Table(collectionName).Where("id", id).Update(newObj); err != nil {
		return err
	}
	return nil
}
