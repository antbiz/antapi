package dao

import (
	"github.com/antbiz/antapi/internal/app/global"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
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

	if schema == nil {
		doc["_id"] = doc["_id"].(primitive.ObjectID).Hex()
		return gjson.New(doc)
	}

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

	if schema == nil {
		for _, doc := range docs {
			doc["_id"] = doc["_id"].(primitive.ObjectID).Hex()
		}
		return gjson.New(docs)
	}

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
