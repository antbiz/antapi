package dao

import (
	"context"
	"time"

	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/encoding/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Upsert 保存单个数据
func Upsert(ctx context.Context, collectionName string, doc interface{}, opts ...*UpsertOptions) (string, error) {
	opt := &UpsertOptions{
		Filter: bson.M{},
	}
	if len(opts) > 0 {
		opt = opts[0]
	}
	jsonDoc := transToGJson(doc)

	// 执行 BeforeSaveHooks 勾子
	for _, hook := range global.GetBeforeSaveHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return "", err
		}
	}

	newDoc := make(map[string]interface{})
	schema := global.GetSchema(collectionName)
	if schema == nil {
		newDoc = jsonDoc.Map()
	} else {
		for _, field := range schema.GetFields(opt.IncludeHiddenField, opt.IncludePrivateField) {
			val := jsonDoc.Get(field.Name)
			if !opt.IgnoreFieldValueCheck {
				if validErr := field.CheckFieldValue(val); validErr != nil {
					return "", validErr
				}
			}
			newDoc[field.Name] = val
		}
	}
	newDoc["updatedAt"] = time.Now().Unix()
	newDoc["updatedBy"] = opt.CtxUser.ID
	if jsonDoc.GetString("createdAt") == "" {
		newDoc["createdAt"] = newDoc["updatedAt"]
	}
	if jsonDoc.GetString("createdBy") == "" {
		newDoc["createdBy"] = opt.CtxUser.ID
	}

	res, err := db.DB().Collection(collectionName).Upsert(ctx, opt.Filter, newDoc)
	if err != nil {
		return "", err
	}
	id := res.UpsertedID.(primitive.ObjectID).Hex()
	newDoc["_id"] = id

	// 执行 AfterSaveHooks 勾子
	var jsonNewDoc *gjson.Json
	for _, hook := range global.GetAfterSaveHooksByCollectionName(collectionName) {
		if jsonNewDoc == nil {
			jsonNewDoc = gjson.New(newDoc)
		}
		if err := hook(ctx, jsonNewDoc); err != nil {
			return "", err
		}
	}

	return id, nil
}
