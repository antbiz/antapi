package crud

import (
	"antapi/model"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// InsertOne : 插入单个数据
// TODO: 返回插入数据的uuid
func InsertOne(collectionName string, data interface{}) error {
	db := g.DB()
	oriObj, err := gjson.LoadJson(data)
	if err != nil {
		return err
	}
	schema, err := model.GetSchema(collectionName)
	if err != nil {
		return err
	}

	newObj := gjson.New(`{}`)
	for _, field := range schema.Fields {
		val := oriObj.Get(field.Name)
		if validErr := field.CheckFieldValue(val); validErr != nil {
			return validErr.Current()
		}
		newObj.Set(field.Name, val)
	}

	_, err = db.Table(collectionName).Data(newObj.Map()).Insert()
	if err != nil {
		return err
	}
	return nil

}

// InsertList : 插入多个数据
func InsertList() {

}
