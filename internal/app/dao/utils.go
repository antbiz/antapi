package dao

import (
	"context"

	"github.com/antbiz/antapi/internal/app/global"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// transToGJson 任意类型数据转gjson对象
func transToGJson(data interface{}) *gjson.Json {
	if val, ok := data.(*gjson.Json); ok {
		return val
	}
	return gjson.New(data)
}

// Doc2GJson 将数据查询结果转换成 gjson
func Doc2GJson(collectionName string, doc map[string]interface{}, includeHiddenField, includePrivateField bool, fieldNames ...string) *gjson.Json {
	schema := global.GetSchema(collectionName)
	targetFieldsTmp := garray.NewStrArrayFrom(fieldNames)
	newDoc := make(map[string]interface{})

	for _, field := range schema.GetFields(includeHiddenField, includePrivateField) {
		if len(fieldNames) > 0 && !targetFieldsTmp.ContainsI(field.Name) {
			continue
		}
		if field.Name == "_id" {
			newDoc[field.Name] = doc[field.Name].(primitive.ObjectID).Hex()
		} else {
			newDoc[field.Name] = doc[field.Name]
		}
	}
	return gjson.New(newDoc)
}

// Docs2GJson 将多个数据查询结果转换成 gjson
func Docs2GJson(collectionName string, docs []map[string]interface{}, includeHiddenField, includePrivateField bool, fieldNames ...string) *gjson.Json {
	schema := global.GetSchema(collectionName)
	targetFieldsTmp := garray.NewStrArrayFrom(fieldNames)
	newDocs := make([]map[string]interface{}, 0)

	for _, doc := range docs {
		newDoc := make(map[string]interface{})
		for _, field := range schema.GetFields(includeHiddenField, includePrivateField) {
			if len(fieldNames) > 0 && !targetFieldsTmp.ContainsI(field.Name) {
				continue
			}
			if field.Name == "_id" {
				newDoc[field.Name] = doc[field.Name].(primitive.ObjectID).Hex()
			} else {
				newDoc[field.Name] = doc[field.Name]
			}
		}
		newDocs = append(newDocs, newDoc)
	}

	return gjson.New(newDocs)
}

// IsDuplicate .
func IsDuplicate(ctx context.Context, collectionName string, doc *gjson.Json, excludeID ...string) (bool, error) {
	schema := global.GetSchema(collectionName)
	fields := schema.GetUniqueFieldNames()
	if len(fields) == 0 {
		return false, nil
	}
	fields = append(fields, "_id")
	filters := make([]bson.M, 0)

	fieldNameMap := map[string]string{}
	for _, field := range schema.GetUniqueFields() {
		fieldNameMap[field.Name] = field.DisplayName
		filters = append(filters, bson.M{field.Name: doc.Get(field.Name)})
	}

	return false, nil
	// results, err := db.DB().Collection(collectionName).Find(ctx, bson.M{"$or": filters}).One()
}
