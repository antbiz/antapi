package logic

import (
	"antapi/app/global"
	"antapi/app/model"
	"antapi/app/model/fieldtype"
	"antapi/app/utils"
	"antapi/common/errcode"
	"antapi/pkg/dbsm"
	dbsmtyp "antapi/pkg/dbsm/types"
	"fmt"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

var Schema = schemaLogic{}

type schemaLogic struct{}

// CheckJSONSchema : 校验 jsonSchema 并填充系统必要字段
func (schemaLogic) CheckJSONSchema(data *gjson.Json) error {
	checkJSONSchema := func(propPath string, isTable bool) error {
		var (
			hasField          = false
			hasIDField        = false
			hasCreatedAtField = false
			hasUpdatedAtField = false
			hasDeletedAtField = false
			hasCreatedByField = false
			hasUpdatedByField = false
			hasIdxField       = false
			hasPidField       = false
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
			case "idx":
				hasIdxField = true
			case "pid":
				hasPidField = true
			}
		}

		if !hasField {
			return gerror.NewCode(errcode.JSONSchemaBadFormat, errcode.JSONSchemaBadFormatMsg)
		}

		if !hasIDField {
			data.Set(fmt.Sprintf("%s.%s", propPath, "id"), g.Map{
				"title":     "编号",
				"type":      "string",
				"ui:hidden": true,
			})
		}
		if !hasCreatedAtField {
			data.Set(fmt.Sprintf("%s.%s", propPath, "created_at"), g.Map{
				"title":     "创建时间",
				"type":      "string",
				"ui:hidden": true,
			})
		}
		if !hasUpdatedAtField {
			data.Set(fmt.Sprintf("%s.%s", propPath, "updated_at"), g.Map{
				"title":     "更新时间",
				"type":      "string",
				"ui:hidden": true,
			})
		}
		if !hasDeletedAtField {
			data.Set(fmt.Sprintf("%s.%s", propPath, "deleted_at"), g.Map{
				"title":     "删除时间",
				"type":      "string",
				"ui:hidden": true,
			})
		}
		if !hasCreatedByField {
			data.Set(fmt.Sprintf("%s.%s", propPath, "created_by"), g.Map{
				"title":     "创建者",
				"type":      "string",
				"ui:hidden": true,
			})
		}
		if !hasUpdatedByField {
			data.Set(fmt.Sprintf("%s.%s", propPath, "updated_by"), g.Map{
				"title":     "修改者",
				"type":      "string",
				"ui:hidden": true,
			})
		}
		if isTable {
			if !hasIdxField {
				data.Set(fmt.Sprintf("%s.%s", propPath, "idx"), g.Map{
					"title":     "序号",
					"type":      "string",
					"ui:hidden": true,
				})
			}
			if !hasPidField {
				data.Set(fmt.Sprintf("%s.%s", propPath, "pid"), g.Map{
					"title":     "父级编号",
					"type":      "string",
					"ui:hidden": true,
				})
			}
		}

		return nil
	}

	if err := checkJSONSchema("schema.properties", false); err != nil {
		return err
	}
	for fieldName := range data.GetMap("schema.properties") {
		fieldType := data.GetString(fmt.Sprintf("schema.properties.%s.type", fieldName))
		if fieldType == "table" {
			if err := checkJSONSchema(fmt.Sprintf("schema.properties.%s.items", fieldName), true); err != nil {
				return err
			}
		}
	}
	return nil
}

// ReloadGlobalSchemas : 当某个Collection的Schema插入/更新/删除后，重新加载数据到内存
func (schemaLogic) ReloadGlobalSchemas(_ *gjson.Json) error {
	global.SchemaChan <- struct{}{}
	return nil
}

// MigrateSchema : 迁移 Schema，同步collection的schema和collection的数据库表结构
func (schemaLogic) MigrateSchema(data *gjson.Json) error {
	tx, err := g.DB().Begin()
	if err != nil {
		return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}

	for _, schema := range utils.ParseFormRenderSchema(data) {
		tableName := schema.CollectionName
		defaultFieldNames := garray.NewStrArrayFrom(model.DefaultFieldNames)
		baseColumns := make([]*dbsm.Column, 0)
		bizColumns := make([]*dbsm.Column, 0)

		for _, field := range schema.Fields {
			if field.Type == fieldtype.Table {
				continue
			}
			col := &dbsm.Column{
				Name:     field.Name,
				Default:  field.Default,
				IsUnique: field.IsUnique,
				Comment:  field.Title,
			}

			switch field.Type {
			case fieldtype.String, fieldtype.Enum, fieldtype.AutoComplete:
				col.Type = dbsmtyp.VARCHAR
			case fieldtype.UUID, fieldtype.Link:
				col.Type = dbsmtyp.VARCHAR
				col.IndexName = col.Name
			case fieldtype.Email, fieldtype.Phone:
				col.Type = dbsmtyp.VARCHAR
				col.Size = 100
			case fieldtype.Color, fieldtype.Password:
				col.Type = dbsmtyp.VARCHAR
				col.Size = 100
				col.IndexName = ""
				col.IsUnique = false
			case fieldtype.URL, fieldtype.SmallText, fieldtype.File:
				col.Type = dbsmtyp.SMALLTEXT
				col.IndexName = ""
			case fieldtype.Text, fieldtype.RichText, fieldtype.Markdown, fieldtype.Code, fieldtype.HTML:
				col.Type = dbsmtyp.TEXT
				col.IndexName = ""
				col.IsUnique = false
			case fieldtype.Signature:
				col.Type = dbsmtyp.BLOB
				col.IndexName = ""
				col.IsUnique = false
			case fieldtype.JSON:
				col.Type = dbsmtyp.JSON
				col.IndexName = ""
				col.IsUnique = false
			case fieldtype.Int, fieldtype.Money:
				col.Type = dbsmtyp.INT
			case fieldtype.BigInt:
				col.Type = dbsmtyp.BIGINT
			case fieldtype.Float:
				col.Type = dbsmtyp.FLOAT
			case fieldtype.Date:
				col.Type = dbsmtyp.DATE
			case fieldtype.DateTime:
				col.Type = dbsmtyp.DATETIME
			case fieldtype.Time:
				col.Type = dbsmtyp.TIME
			case fieldtype.TimeStamp:
				col.Type = dbsmtyp.TIMESTAMP
			case fieldtype.Year:
				col.Type = dbsmtyp.YEAR
			case fieldtype.Bool:
				col.Type = dbsmtyp.BOOL
			}

			col.Nullable = col.Name != "id"
			col.IsPrimaryKey = col.Name == "id"
			if col.Name == "id" {
				col.IndexName = ""
			} else if field.CanIndex {
				col.IndexName = col.Name
			}

			if defaultFieldNames.Contains(col.Name) {
				baseColumns = append(baseColumns, col)
			} else {
				bizColumns = append(bizColumns, col)
			}
		}
		baseColumns = append(baseColumns, bizColumns...)

		table := dbsm.NewTable(tableName, baseColumns)
		dialect, err := getDialect()
		if err != nil {
			return err
		}
		if err := table.Sync(tx, dialect); err != nil {
			return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
		}

		return nil
	}
	return nil
}

// AutoExportSchemaData 保存数据到 boot/schemas/biz 以便项目初始化
func (schemaLogic) AutoExportSchemaData(data *gjson.Json) error {
	glog.Info("Auto Export Schema Data To boot/schemas/biz")
	exportPath := gfile.Join(gfile.Pwd(), "boot", "schemas", "biz", fmt.Sprintf("%s.json", data.GetString("collection_name")))

	// 将data复制一份
	_data := new(gjson.Json)
	*_data = *data

	fieldsLen := len(_data.GetArray("fields"))
	for _, fieldName := range model.DefaultFieldNames {
		_data.Remove(fieldName)
		for i := 0; i < fieldsLen; i++ {
			_data.Remove(fmt.Sprintf("fields.%d.%s", i, fieldName))
		}
	}
	if err := gfile.PutContents(exportPath, _data.MustToJsonIndentString()); err != nil {
		glog.Fatal(err)
	}

	return nil
}

// AutoDeleteExportedJsonFile 删除导出的json文件
func (schemaLogic) AutoDeleteExportedJsonFile(data *gjson.Json) error {
	glog.Info("Auto Delete Exported Json File From boot/schemas/biz")
	jsonFilePath := gfile.Join(gfile.Pwd(), "boot", "schemas", "biz", fmt.Sprintf("%s.json", data.GetString("collection_name")))

	if err := gfile.Remove(jsonFilePath); err != nil {
		glog.Fatal(err)
	}

	return nil
}

func getDialect() (dbsm.Dialect, error) {
	db := g.DB()
	dbType := db.GetConfig().Type

	switch dbsmtyp.DBType(dbType) {
	case dbsmtyp.MYSQL:
		dialect := &dbsm.MySQLDialect{
			DBName:  db.GetConfig().Name,
			Charset: "utf8mb4",
		}
		return dialect, nil
	default:
		return nil, gerror.Newf("Not implemented %s", dbType)
	}
}
