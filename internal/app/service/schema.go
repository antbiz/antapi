package service

import (
	"errors"
	"fmt"

	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/app/global"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

var Schema = &schemaSrv{
	collectionName: "schema",
}

type schemaSrv struct {
	collectionName string
}

// CollectionName .
func (srv *schemaSrv) CollectionName() string {
	return srv.collectionName
}

// CheckJSONSchema 检验 JSONSchema 并填充系统必要字段
func (srv *schemaSrv) CheckJSONSchema(data *gjson.Json) error {
	checkJSONSchema := func(propPath string) error {
		var (
			hasField          = false
			hasIDField        = false
			hasCreatedAtField = false
			hasUpdatedAtField = false
			hasDeletedAtField = false
			hasCreatedByField = false
			hasUpdatedByField = false
		)

		for fieldName := range data.GetMap(propPath) {
			if !hasField {
				hasField = true
			}
			switch fieldName {
			case "id":
				hasIDField = true
			case "created_at":
				hasCreatedAtField = true
			case "updated_at":
				hasUpdatedAtField = true
			case "deleted_at":
				hasDeletedAtField = true
			case "created_by":
				hasCreatedByField = true
			case "updated_by":
				hasUpdatedByField = true
			}

			if !hasField {
				return errors.New("无效的schema")
			}

			if !hasIDField {
				data.Set(fmt.Sprintf("%s.%s", propPath, "id"), g.Map{
					"title":  "编号",
					"type":   "string",
					"hidden": true,
				})
			}
			if !hasCreatedAtField {
				data.Set(fmt.Sprintf("%s.%s", propPath, "created_at"), g.Map{
					"title":  "创建时间",
					"type":   "string",
					"index":  true,
					"hidden": true,
				})
			}
			if !hasUpdatedAtField {
				data.Set(fmt.Sprintf("%s.%s", propPath, "updated_at"), g.Map{
					"title":  "更新时间",
					"type":   "string",
					"hidden": true,
				})
			}
			if !hasDeletedAtField {
				data.Set(fmt.Sprintf("%s.%s", propPath, "deleted_at"), g.Map{
					"title":  "删除时间",
					"type":   "string",
					"hidden": true,
				})
			}
			if !hasCreatedByField {
				data.Set(fmt.Sprintf("%s.%s", propPath, "created_by"), g.Map{
					"title":  "创建者",
					"type":   "string",
					"index":  true,
					"hidden": true,
				})
			}
			if !hasUpdatedByField {
				data.Set(fmt.Sprintf("%s.%s", propPath, "updated_by"), g.Map{
					"title":  "修改者",
					"type":   "string",
					"hidden": true,
				})
			}
		}

		return nil
	}

	if err := checkJSONSchema("schema.properties"); err != nil {
		return err
	}

	return nil
}

// ReloadGlobalSchemas 当某个Collection的Schema插入/更新/删除后，重新加载数据到内存
func (srv *schemaSrv) ReloadGlobalSchemas(_ *gjson.Json) error {
	global.SchemaChan <- struct{}{}
	return nil
}

// AutoCreateCollectionPermission 新建模型后初始化权限设置
func (srv *schemaSrv) AutoCreateCollectionPermission(data *gjson.Json) error {
	return nil
}

// AutoDeleteCollectionPermission 删除模型后移除对应的权限设置
func (srv *schemaSrv) AutoDeleteCollectionPermission(data *gjson.Json) error {
	return nil
}

// GetJSONFilePath 获取文件备份导出路径
func (srv *schemaSrv) GetJSONFilePath(collectionName string) string {
	return gfile.Join(gfile.Pwd(), "internal", "data", "schemas", "biz", fmt.Sprintf("%s.json", collectionName))
}

// AutoExportSchemaData 保存数据到 boot/schemas/biz 以便项目初始化
func (srv *schemaSrv) AutoExportJSONFile(data *gjson.Json) error {
	g.Log().Info("Auto Export Schema Data To internal/data/schemas/biz")
	jsonFilePath := srv.GetJSONFilePath(data.GetString("collectionName"))

	// 将data复制一份
	_data := new(gjson.Json)
	*_data = *data

	fieldsLen := len(_data.GetArray("fields"))
	for _, fieldName := range dto.DefaultFieldNames {
		_data.Remove(fieldName)
		for i := 0; i < fieldsLen; i++ {
			_data.Remove(fmt.Sprintf("fields.%d.%s", i, fieldName))
		}
	}
	if err := gfile.PutContents(jsonFilePath, _data.MustToJsonIndentString()); err != nil {
		g.Log().Fatal(err)
	}

	return nil
}

// AutoDeleteJSONFile 删除导出的json文件
func (srv *schemaSrv) AutoDeleteJSONFile(data *gjson.Json) error {
	g.Log().Info("Auto Delete Exported Json File From internal/data/schemas/biz")
	jsonFilePath := srv.GetJSONFilePath(data.GetString("collectionName"))

	if err := gfile.Remove(jsonFilePath); err != nil {
		g.Log().Fatal(err)
	}

	return nil
}
