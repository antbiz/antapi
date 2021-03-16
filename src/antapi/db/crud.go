package db

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
)

func buildOneToOneDataKey(fieldName string) string {
	return fmt.Sprintf("_%s", fieldName)
}

// GetOne : 获取单个数据
func GetOne(collectionName string, where interface{}, args ...interface{}) (map[string]interface{}, error) {
	db := g.DB()
	schema, err := GetSchema(collectionName)
	if err != nil {
		return nil, err
	}

	record, err := db.Table(collectionName).Fields(schema.GetFieldNames()).Where(where, args...).One()
	if err != nil {
		return nil, err
	}
	obj := record.GMap()

	for _, field := range schema.GetOneToOneFields() {
		connectSchema, err := GetSchema(field.ConnectCollection)
		if err != nil {
			return nil, err
		}
		oneToOneRecord, err := db.Table(field.ConnectCollection).Fields(connectSchema.GetFieldNames()).Where("id", obj.Get(field.Name)).One()
		if err != nil {
			return nil, err
		}
		obj.Set(buildOneToOneDataKey(field.Name), oneToOneRecord.Map())
	}

	return obj.MapStrAny(), nil
}
