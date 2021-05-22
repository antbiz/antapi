package dao

import (
	"context"

	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

// Delete 删除单个数据
func Delete(ctx context.Context, collectionName string, opts ...*DeleteOptions) error {
	opt := &DeleteOptions{
		Filter: bson.M{},
	}
	if len(opts) > 0 {
		opt = opts[0]
	}

	jsonDoc, err := Get(ctx, collectionName, &GetOptions{
		Filter: opt.Filter,
	})
	if err != nil {
		return err
	}

	// 执行 BeforeDelete 勾子
	for _, hook := range global.GetBeforeDeleteHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return err
		}
	}

	if err := db.DB().Collection(collectionName).Remove(ctx, opt.Filter); err != nil {
		return err
	}

	// 执行 AfterDelete 勾子
	for _, hook := range global.GetAfterDeleteHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return err
		}
	}

	return nil
}

// DeleteMany 删除多个数据
func DeleteMany(ctx context.Context, collectionName string, filter interface{}) error {
	_, err := db.DB().Collection(collectionName).RemoveAll(ctx, filter)
	return err
}
