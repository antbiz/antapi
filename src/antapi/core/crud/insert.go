package crud

import (
	"github.com/gogf/gf/frame/g"
)

// InsertOne : 插入单个数据
// TODO: 返回插入数据的uuid，过滤hidden字段和private字段，校验必填，校验数据是否合法，
func InsertOne(collectionName string, data interface{}) error {
	db := g.DB()
	// obj, err := gjson.LoadJson(data)
	// if err != nil {
	// 	return err
	// }
	// schema, err := model.GetSchema(collectionName)
	// if err != nil {
	// 	return err
	// }

	_, err := db.Table(collectionName).Data(data).Insert()
	if err != nil {
		return err
	}
	return nil

}

// InsertList : 插入多个数据
func InsertList() {

}
