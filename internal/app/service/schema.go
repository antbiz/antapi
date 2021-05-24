package service

import (
	"context"
	"fmt"

	"github.com/antbiz/antapi/internal/app/global"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/qiniu/qmgo/options"
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
	return nil
}

// CreateIndexes 创建索引
func (srv *schemaSrv) CreateIndexes(ctx context.Context, data *gjson.Json) error {
	uniqueIndexes := make([]options.IndexModel, 0)
	normalIndexes := make([]options.IndexModel, 0)
	for fieldName := range data.GetMap("properties") {
		isUnique := data.GetBool(fmt.Sprintf("properties.%s.unique", fieldName))
		isIndexField := data.GetBool(fmt.Sprintf("properties.%s.index", fieldName))
		if isUnique {
			uniqueIndexes = append(uniqueIndexes, options.IndexModel{
				Key: []string{fieldName}, Unique: true,
			})
		}
		if isIndexField {
			normalIndexes = append(normalIndexes, options.IndexModel{
				Key: []string{fieldName},
			})
		}
	}

	uniqueIndexesLen := len(uniqueIndexes)
	normalIndexesLen := len(normalIndexes)
	if uniqueIndexesLen == 0 && normalIndexesLen == 0 {
		return nil
	}
	cli := db.DB().Collection(data.GetString("name"))

	if uniqueIndexesLen > 0 {
		if err := cli.CreateIndexes(ctx, uniqueIndexes); err != nil {
			return err
		}
	}
	if normalIndexesLen > 0 {
		if err := cli.CreateIndexes(ctx, normalIndexes); err != nil {
			return err
		}
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
				"title":          data.GetString("displayName"),
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
