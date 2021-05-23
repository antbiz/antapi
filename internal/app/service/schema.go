package service

import (
	"context"
	"fmt"

	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"go.mongodb.org/mongo-driver/bson"
)

var Schema = &schemaSrv{}

type schemaSrv struct {
	baseSysSrv
}

// CollectionName .
func (srv *schemaSrv) CollectionName() string {
	return "schema"
}

// CheckJSONSchema 检验 JSONSchema
func (srv *schemaSrv) CheckJSONSchema(ctx context.Context, data *gjson.Json) error {
	const propPath = "properties"
	var (
		hasIDField        = false
		hasCreatedAtField = false
		hasUpdatedAtField = false
		hasCreatedByField = false
		hasUpdatedByField = false
	)

	for fieldName := range data.GetMap(propPath) {
		switch fieldName {
		case "_id":
			hasIDField = true
		case "createdAt":
			hasCreatedAtField = true
		case "updatedAt":
			hasUpdatedAtField = true
		case "createdBy":
			hasCreatedByField = true
		case "updatedBy":
			hasUpdatedByField = true
		}
	}

	if !hasIDField {
		data.Set(fmt.Sprintf("%s.%s", propPath, "_id"), g.Map{
			"title":  "ID",
			"type":   "string",
			"hidden": true,
		})
	}
	if !hasCreatedAtField {
		data.Set(fmt.Sprintf("%s.%s", propPath, "createdAt"), g.Map{
			"title":  "创建时间",
			"type":   "string",
			"index":  true,
			"hidden": true,
		})
	}
	if !hasUpdatedAtField {
		data.Set(fmt.Sprintf("%s.%s", propPath, "updatedAt"), g.Map{
			"title":  "更新时间",
			"type":   "string",
			"hidden": true,
		})
	}
	if !hasCreatedByField {
		data.Set(fmt.Sprintf("%s.%s", propPath, "createdBy"), g.Map{
			"title":  "创建者",
			"type":   "string",
			"index":  true,
			"hidden": true,
		})
	}
	if !hasUpdatedByField {
		data.Set(fmt.Sprintf("%s.%s", propPath, "updatedBy"), g.Map{
			"title":  "修改者",
			"type":   "string",
			"hidden": true,
		})
	}

	return nil
}

// ReloadGlobalSchemas 当某个Collection的Schema插入/更新/删除后，重新加载数据到内存
func (srv *schemaSrv) ReloadGlobalSchemas(ctx context.Context, data *gjson.Json) error {
	global.SchemaChan <- struct{}{}
	return nil
}

// AutoCreateCollectionPermission 新建模型后初始化权限设置
func (srv *schemaSrv) AutoCreateCollectionPermission(ctx context.Context, data *gjson.Json) error {
	_, err := db.
		DB().
		Collection(Permission.CollectionName()).
		Upsert(
			ctx,
			bson.M{"collectionName": data.GetString("name")},
			g.Map{
				"title":          data.GetString("title"),
				"projectName":    data.GetString("projectName"),
				"collectionName": data.GetString("name"),
				"createLevel":    0,
				"readLevel":      0,
				"updateLevel":    0,
				"deleteLevel":    0,
			},
		)
	return err
}

// AutoDeleteCollectionPermission 删除模型后移除对应的权限设置
func (srv *schemaSrv) AutoDeleteCollectionPermission(ctx context.Context, data *gjson.Json) error {
	_, err := db.
		DB().
		Collection(Permission.CollectionName()).
		RemoveAll(ctx, bson.M{"collectionName": data.GetString("name")})
	return err
}

// GetJSONFilePath 获取文件备份导出路径
func (srv *schemaSrv) GetJSONFilePath(collectionName string) string {
	return gfile.Join(gfile.Pwd(), "internal", "data", "schemas", "biz", fmt.Sprintf("%s.json", collectionName))
}

// AutoExportJSONFile .
func (srv *schemaSrv) AutoExportJSONFile(ctx context.Context, data *gjson.Json) error {
	return srv.ExportJSONFile(data, srv.GetJSONFilePath(data.GetString("name")))
}

// AutoDeleteJSONFile .
func (srv *schemaSrv) AutoDeleteJSONFile(ctx context.Context, data *gjson.Json) error {
	return srv.DeleteJSONFile(srv.GetJSONFilePath(data.GetString("name")))
}
