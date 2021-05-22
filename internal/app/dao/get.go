package dao

import (
	"context"

	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/encoding/gjson"
	"go.mongodb.org/mongo-driver/bson"
)

// List 获取列表数据
func List(ctx context.Context, collectionName string, opts ...*ListOptions) (*gjson.Json, int, error) {
	opt := &ListOptions{
		Filter: bson.M{},
	}
	if len(opts) > 0 {
		opt = opts[0]
	}

	q := db.DB().Collection(collectionName).Find(ctx, opt.Filter)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	if opt.Limit > 0 && opt.Offset > 0 {
		q.Limit(opt.Limit).Skip(opt.Offset)
	}
	docs := make([]map[string]interface{}, 0)
	if err := q.Sort(opt.Sort...).All(&docs); err != nil {
		return nil, 0, err
	}
	jsonDoc := Docs2GJson(collectionName, docs, opt.IncludeHiddenField, opt.IncludePrivateField, opt.Fields...)

	// 执行 AfterFindHooks 勾子
	for _, hook := range global.GetAfterFindHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return nil, 0, err
		}
	}

	return jsonDoc, int(total), nil
}

// Get 获取单个数据
func Get(ctx context.Context, collectionName string, opts ...*GetOptions) (*gjson.Json, error) {
	opt := &GetOptions{
		Filter: bson.M{},
	}
	if len(opts) > 0 {
		opt = opts[0]
	}

	doc := make(map[string]interface{})
	if err := db.DB().Collection(collectionName).Find(ctx, opt.Filter).One(&doc); err != nil {
		return nil, err
	}
	jsonDoc := Doc2GJson(collectionName, doc, opt.IncludeHiddenField, opt.IncludePrivateField, opt.Fields...)

	// 执行 AfterFindHooks 勾子
	for _, hook := range global.GetAfterFindHooksByCollectionName(collectionName) {
		if err := hook(ctx, jsonDoc); err != nil {
			return nil, err
		}
	}

	return jsonDoc, nil
}
