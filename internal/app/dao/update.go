package dao

import (
	"context"
	"time"

	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

// Update 更新单个数据
func Update(ctx context.Context, collectionName string, doc interface{}, opts ...*UpdateOptions) error {
	opt := &UpdateOptions{
		Filter: bson.M{},
	}
	if len(opts) > 0 {
		opt = opts[0]
	}
	jsonDoc := transToGJson(doc)

	// 执行 BeforeInsertHooks, BeforeSaveHooks 勾子
	for _, hook := range global.GetBeforeUpdateHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return err
		}
	}
	for _, hook := range global.GetBeforeSaveHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return err
		}
	}

	newDoc := make(map[string]interface{})
	schema := global.GetSchema(collectionName)
	if schema == nil {
		jsonDoc.Remove("_id")
		newDoc = jsonDoc.Map()
	} else {
		for _, field := range schema.GetFields(opt.IncludeHiddenField, opt.IncludePrivateField) {
			if field.IsSysField {
				continue
			}
			if !jsonDoc.Contains(field.Name) {
				continue
			}
			val := jsonDoc.Get(field.Name)
			if !opt.IgnoreFieldValueCheck {
				if validErr := field.CheckFieldValue(val); validErr != nil {
					return validErr
				}
			}
			newDoc[field.Name] = val
		}
	}
	newDoc["updatedAt"] = time.Now().Unix()
	newDoc["updatedBy"] = opt.CtxUser.ID

	if err := db.DB().Collection(collectionName).UpdateOne(ctx, opt.Filter, bson.M{"$set": doc}); err != nil {
		return err
	}

	// 执行 AfterUpdateHooks, AfterSaveHooks 勾子
	for _, hook := range global.GetAfterUpdateHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return err
		}
	}
	for _, hook := range global.GetAfterSaveHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return err
		}
	}

	return nil
}

// UpdateMany 更新多个数据
func UpdateMany(ctx context.Context, collectionName string, doc interface{}, filter interface{}) error {
	_, err := db.DB().Collection(collectionName).UpdateAll(ctx, filter, doc)
	return err
}
