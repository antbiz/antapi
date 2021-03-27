package logic

import (
	"antapi/app/global"
	"antapi/app/model"
	"antapi/app/model/fieldtype"
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

var Schema = new(schemaLogic)

type schemaLogic struct{}

// CheckFields : 校验collection的字段，并填充系统必要字段
func (schemaLogic) CheckFields(data *gjson.Json) error {
	fieldsLen := len(data.GetArray("fields"))
	if fieldsLen == 0 {
		return gerror.NewCodef(errcode.MissRequiredParameter, errcode.MissRequiredParameterMsg, "fields")
	}

	isChildTable := data.GetBool("is_child")

	var (
		hasIdField        bool
		hasCreatedAtField bool
		hasUpdatedAtField bool
		hasDeletedAtField bool
		hasCreatedByField bool
		hasUpdatedByField bool
		hasPcnField       bool
		hasIdxField       bool
		hasPidField       bool
		hasPfdField       bool
	)

	for i := 0; i < fieldsLen; i++ {
		fieldName := data.GetString(fmt.Sprintf("fields.%d.name", i))
		switch fieldName {
		case "id":
			hasIdField = true
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
		case "pcn":
			hasPcnField = true
		case "idx":
			hasIdxField = true
		case "pid":
			hasPidField = true
		case "pfd":
			hasPfdField = true
		}

		if data.GetBool(fmt.Sprintf("fields.%d.is_unique", i)) {
			data.Set(fmt.Sprintf("fields.%d.is_required", i), true)
		}
	}

	addField := func(v interface{}) {
		data.Append("fields", v)
	}

	if !hasIdField {
		addField(g.Map{
			"type":      "UUID",
			"title":     "ID",
			"name":      "id",
			"is_unique": false,
			"can_index": false,
		})
	}
	if !hasCreatedAtField {
		addField(g.Map{
			"type":      "DateTime",
			"title":     "Created At",
			"name":      "created_at",
			"can_index": true,
		})
	}
	if !hasUpdatedAtField {
		addField(g.Map{
			"type":  "DateTime",
			"title": "Updated At",
			"name":  "updated_at",
		})
	}
	if !hasDeletedAtField {
		addField(g.Map{
			"type":  "DateTime",
			"title": "Deleted At",
			"name":  "deleted_at",
		})
	}
	if !hasCreatedByField {
		addField(g.Map{
			"type":      "String",
			"title":     "Created By",
			"name":      "created_by",
			"can_index": true,
		})
	}
	if !hasUpdatedByField {
		addField(g.Map{
			"type":  "String",
			"title": "Updated By",
			"name":  "updated_by",
		})
	}
	if isChildTable {
		if !hasPcnField {
			addField(g.Map{
				"type":      "String",
				"title":     "Parent Co",
				"name":      "pcn",
				"can_index": true,
			})
		}
		if !hasIdxField {
			addField(g.Map{
				"type":  "Int",
				"title": "Index",
				"name":  "idx",
			})
		}
		if !hasPidField {
			addField(g.Map{
				"type":      "String",
				"title":     "Parent ID",
				"name":      "pid",
				"can_index": true,
			})
		}
		if !hasPfdField {
			addField(g.Map{
				"type":      "String",
				"title":     "Parent Field",
				"name":      "pfd",
				"can_index": true,
			})
		}
	}

	return nil
}

// ReloadGlobalSchemas : 当某个Collection的Schema插入/更新/删除后，重新加载数据到内存
func (schemaLogic) ReloadGlobalSchemas(_ *gjson.Json) error {
	global.SchemaChan <- struct{}{}
	return nil
}

// MigrateCollectionSchema : 迁移collection，同步collection的schema和collection的数据库表结构
func (schemaLogic) MigrateCollectionSchema(collection *gjson.Json) error {
	tableName := collection.GetString("name")
	defaultFieldNames := garray.NewStrArrayFrom(model.DefaultFieldNames)
	baseColumns := make([]*dbsm.Column, 0)
	bizColumns := make([]*dbsm.Column, 0)
	for _, field := range collection.GetJsons("fields") {
		fieldType := fieldtype.FieldType(field.GetString("type"))
		if fieldType == fieldtype.Table {
			continue
		}
		col := &dbsm.Column{
			Name:     field.GetString("name"),
			Default:  field.GetString("default"),
			IsUnique: field.GetBool("is_unique"),
			Comment:  field.GetString("title"),
		}

		if field.GetBool("is_multiple") {
			col.Type = dbsmtyp.JSON
		} else {
			switch fieldType {
			case fieldtype.String, fieldtype.Enum:
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
			case fieldtype.URL, fieldtype.SmallText, fieldtype.Media:
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
		}

		col.Nullable = col.Name != "id"
		col.IsPrimaryKey = col.Name == "id"
		if col.Name == "id" {
			col.IndexName = ""
		} else if field.GetBool("can_index") {
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
	db := g.DB()
	tx, err := db.Begin()
	if err != nil {
		return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}
	dialect, err := getDialect()
	if err != nil {
		return err
	}
	if err := table.Sync(tx, dialect); err != nil {
		return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}

	return nil
}

// AutoExportSchemaData 保存数据到 app/model/collection 以便项目初始化
func (schemaLogic) AutoExportSchemaData(data *gjson.Json) error {
	glog.Info("Auto Export Schema Data To app/model/collection")
	exportPath := gfile.Join(gfile.Pwd(), "app", "model", "collection", fmt.Sprintf("%s.json", data.GetString("name")))

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
	glog.Info("Auto Delete Exported Json File From app/model/collection")
	jsonFilePath := gfile.Join(gfile.Pwd(), "app", "model", "collection", fmt.Sprintf("%s.json", data.GetString("name")))

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
