package schema

import (
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// GetValue : 获取单个数据
func GetValue(collectionName string, where interface{}, args ...interface{}) (map[string]interface{}, error) {
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

	for _, field := range schema.GetBelongsToFields() {
		relatedSchema, err := GetSchema(field.RelatedCollection)
		if err != nil {
			return nil, err
		}
		relatedRecord, err := db.Table(field.RelatedCollection).Fields(relatedSchema.GetFieldNames()).Where("id", obj.Get(field.Name)).One()
		if err != nil {
			return nil, err
		}
		obj.Set(field.RelatedCollection, relatedRecord.Map())
	}

	for _, field := range schema.GetHasManyFields() {
		relatedSchema, err := GetSchema(field.RelatedCollection)
		if err != nil {
			return nil, err
		}
		relatedRecords, err := db.Table(field.RelatedCollection).Fields(relatedSchema.GetFieldNames()).Order("idx asc").Where("pid", obj.Get("id")).All()
		if err != nil {
			return nil, err
		}
		obj.Set(field.RelatedCollection, relatedRecords.List())
	}

	return obj.MapStrAny(), nil
}

// GetList : 获取列表数据
func GetList(collectionName string, pageNum, pageSize int, where interface{}, args ...interface{}) ([]map[string]interface{}, error) {
	db := g.DB()
	schema, err := GetSchema(collectionName)
	if err != nil {
		return nil, err
	}

	model := db.Table(collectionName).Fields(schema.GetFieldNames()).Where(where, args...)
	if pageNum > 0 && pageSize > 0 {
		model = model.Limit((pageNum-1)*pageSize, pageSize)
	}
	records, err := model.All()
	if err != nil {
		return nil, err
	}
	recordsLen := records.Len()
	objs := gjson.New(records.Json())

	var objIds []string
	for i := 0; i < recordsLen; i++ {
		objIds = append(objIds, objs.GetString(fmt.Sprintf("%d.id", i)))
	}

	for _, field := range schema.GetBelongsToFields() {
		relatedSchema, err := GetSchema(field.RelatedCollection)
		if err != nil {
			return nil, err
		}

		relatedRecords, err := db.Table(field.RelatedCollection).Fields(relatedSchema.GetFieldNames()).Where("id", objIds).All()
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

	for _, field := range schema.GetHasManyFields() {
		relatedSchema, err := GetSchema(field.RelatedCollection)
		if err != nil {
			return nil, err
		}
		relatedRecords, err := db.Table(field.RelatedCollection).Fields(relatedSchema.GetFieldNames()).Order("idx asc").Where("pid", objIds).All()
		if err != nil {
			return nil, err
		}

		var mapObjRelatedRecords map[string][]string
		for _, relatedRecord := range relatedRecords {
			recordObj := relatedRecord.GMap()
			pid := recordObj.GetVar("pid").String()
			mapObjRelatedRecords[pid] = append(mapObjRelatedRecords[pid], relatedRecord.Json())
		}

		for i := 0; i < recordsLen; i++ {
			pid := objs.GetString(fmt.Sprintf("%d.id", i))
			if err := objs.Set(fmt.Sprintf("%d.%s", i, field.Name), mapObjRelatedRecords[pid]); err != nil {
				return nil, err
			}
		}
	}

	return objs.Var().Maps(), nil
}
