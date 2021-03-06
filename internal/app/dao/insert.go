package dao

import (
	"context"
	"time"

	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/encoding/gjson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert 新建单个数据
func Insert(ctx context.Context, collectionName string, doc interface{}, opts ...*InsertOptions) (string, error) {
	var opt *InsertOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	jsonDoc := transToGJson(doc)

	// 执行 BeforeInsertHooks, BeforeSaveHooks 勾子
	for _, hook := range global.GetBeforeInsertHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return "", err
		}
	}
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
			if field.IsSysField || !jsonDoc.Contains(field.Name) {
				continue
			}
			if field.HasDefault && jsonDoc.GetString(field.Name) == "" {
				jsonDoc.Set(field.Name, field.Default)
				newDoc[field.Name] = field.Default
			} else {
				val := jsonDoc.Get(field.Name)
				if !opt.IgnoreFieldValueCheck {
					if validErr := field.CheckFieldValue(val); validErr != nil {
						return "", validErr
					}
				}
				newDoc[field.Name] = val
			}
		}
	}
	newDoc["createdAt"] = time.Now().Unix()
	newDoc["updatedAt"] = newDoc["createdAt"]
	newDoc["createdBy"] = opt.CtxUser.ID
	newDoc["updatedBy"] = opt.CtxUser.ID

	result, err := db.DB().Collection(collectionName).InsertOne(ctx, newDoc)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	newDoc["_id"] = id

	// 执行 AfterInsertHooks, AfterSaveHooks 勾子
	var jsonNewDoc *gjson.Json
	for _, hook := range global.GetAfterInsertHooksByCollectionName(collectionName) {
		if jsonNewDoc == nil {
			jsonNewDoc = gjson.New(newDoc)
		}
		if err := hook(ctx, jsonNewDoc); err != nil {
			return "", err
		}
	}
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

// InsertMany 新建多个数据
func InsertMany(ctx context.Context, collectionName string, docs ...interface{}) ([]string, error) {
	if len(docs) == 0 {
		return nil, nil
	}

	result, err := db.DB().Collection(collectionName).InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(result.InsertedIDs))
	for _, id := range result.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID).Hex())
	}
	return ids, nil
}
