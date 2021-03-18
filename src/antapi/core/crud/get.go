package crud

import (
	"antapi/model"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// GetOne : 获取单个数据
func GetOne(collectionName string, where interface{}, args ...interface{}) (map[string]interface{}, error) {
	db := g.DB()
	schema, err := model.GetSchema(collectionName)
	if err != nil {
		return nil, err
	}

	record, err := db.Table(collectionName).Fields(schema.GetPublicFieldNames()).Where(where, args...).One()
	if err != nil {
		return nil, err
	}
	obj := record.GMap()

	for _, field := range schema.GetLinkFields() {
		relatedSchema, err := model.GetSchema(field.RelatedCollection)
		if err != nil {
			return nil, err
		}
		relatedRecord, err := db.Table(field.RelatedCollection).Fields(relatedSchema.GetPublicFieldNames()).Where("id", obj.Get(field.Name)).One()
		if err != nil {
			return nil, err
		}
		obj.Set(field.RelatedCollection, relatedRecord.Map())
	}

	return obj.MapStrAny(), nil
}

// GetList : 获取列表数据
func GetList(collectionName string, pageNum, pageSize int, where interface{}, args ...interface{}) ([]map[string]interface{}, error) {
	db := g.DB()
	schema, err := model.GetSchema(collectionName)
	if err != nil {
		return nil, err
	}

	orm := db.Table(collectionName).Fields(schema.GetPublicFieldNames()).Where(where, args...)
	if pageNum > 0 && pageSize > 0 {
		orm = orm.Limit((pageNum-1)*pageSize, pageSize)
	}
	records, err := orm.All()
	if err != nil {
		return nil, err
	}
	recordsLen := records.Len()
	objs := gjson.New(records.Json())

	var objIds []string
	for i := 0; i < recordsLen; i++ {
		objIds = append(objIds, objs.GetString(fmt.Sprintf("%d.id", i)))
	}

	for _, field := range schema.GetLinkFields() {
		relatedSchema, err := model.GetSchema(field.RelatedCollection)
		if err != nil {
			return nil, err
		}

		relatedRecords, err := db.Table(field.RelatedCollection).Fields(relatedSchema.GetPublicFieldNames()).Where("id", objIds).All()
		if err != nil {
			return nil, err
		}
		relatedObjs := relatedRecords.MapKeyValue("id")
		for i := 0; i < recordsLen; i++ {
			relatedId := objs.GetString(fmt.Sprintf("%d.%s", i, field.Name))
			if err := objs.Set(fmt.Sprintf("%d.%s", i, field.RelatedCollection), relatedObjs[relatedId]); err != nil {
				return nil, err
			}
		}
	}

	return objs.Var().Maps(), nil
}