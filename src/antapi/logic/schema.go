package logic

import (
	"antapi/common/errcode"
	"antapi/global"
	"antapi/model"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
)

type SchemaLogic struct{}

var DefaultSchemaLogic = SchemaLogic{}

// CheckFields : 校验collection的字段，并填充系统必要字段
func (SchemaLogic) CheckFields(data *gjson.Json) error {
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

	getDataPathForField := func(i int, name string) string {
		return fmt.Sprintf("fields.%d.%s", i, name)
	}

	for i := 0; i < fieldsLen; i++ {
		fieldName := getDataPathForField(i, "name")
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
	}

	updateField := func(i int, name string, v interface{}) {
		data.Set(fmt.Sprintf("fields.%d.%s", i, name), v)
	}

	if !hasIdField {
		fieldsLen++
		updateField(fieldsLen, "type", "UUID")
		updateField(fieldsLen, "title", "ID")
		updateField(fieldsLen, "name", "id")
		updateField(fieldsLen, "is_unique", true)
		updateField(fieldsLen, "can_index", true)
	}
	if !hasCreatedAtField {
		fieldsLen++
		updateField(fieldsLen, "type", "DateTime")
		updateField(fieldsLen, "title", "Created At")
		updateField(fieldsLen, "name", "created_at")
		updateField(fieldsLen, "can_index", true)
	}
	if !hasUpdatedAtField {
		fieldsLen++
		updateField(fieldsLen, "type", "DateTime")
		updateField(fieldsLen, "title", "Updated At")
		updateField(fieldsLen, "name", "updated_at")
	}
	if !hasDeletedAtField {
		fieldsLen++
		updateField(fieldsLen, "type", "DateTime")
		updateField(fieldsLen, "title", "Deleted At")
		updateField(fieldsLen, "name", "deleted_at")
	}
	if !hasCreatedByField {
		fieldsLen++
		updateField(fieldsLen, "type", "String")
		updateField(fieldsLen, "title", "Created By")
		updateField(fieldsLen, "name", "created_by")
		updateField(fieldsLen, "can_index", true)
	}
	if !hasUpdatedByField {
		fieldsLen++
		updateField(fieldsLen, "type", "String")
		updateField(fieldsLen, "title", "Updated By")
		updateField(fieldsLen, "name", "updated_by")
	}
	if isChildTable {
		if !hasPcnField {
			fieldsLen++
			updateField(fieldsLen, "type", "String")
			updateField(fieldsLen, "title", "Parent Collection")
			updateField(fieldsLen, "name", "pcn")
			updateField(fieldsLen, "can_index", true)
		}
		if !hasIdxField {
			fieldsLen++
			updateField(fieldsLen, "type", "Int")
			updateField(fieldsLen, "title", "Index")
			updateField(fieldsLen, "name", "idx")
		}
		if !hasPidField {
			fieldsLen++
			updateField(fieldsLen, "type", "String")
			updateField(fieldsLen, "title", "Parent ID")
			updateField(fieldsLen, "name", "pid")
			updateField(fieldsLen, "can_index", true)
		}
		if !hasPfdField {
			fieldsLen++
			updateField(fieldsLen, "type", "String")
			updateField(fieldsLen, "title", "Parent Field")
			updateField(fieldsLen, "name", "pfd")
			updateField(fieldsLen, "can_index", true)
		}
	}

	return nil
}

// ReloadGlobalSchemas : 当某个Collection的Schema插入/更新/删除后，重新加载数据到内存
func (SchemaLogic) ReloadGlobalSchemas(_ *gjson.Json) error {
	global.SchemaChan <- struct{}{}
	return nil
}

// GetLinkPathIncludeTableInner : 获取所有link字段的路径，包括子表
func (SchemaLogic) GetLinkPathIncludeTableInner(schema *model.Schema) (paths map[string][]string) {
	for _, linkField := range schema.GetLinkFields() {
		paths[linkField.RelatedCollection] = append(paths[linkField.RelatedCollection], linkField.Name)
	}
	for _, tableField := range schema.GetTableFields() {
		tableSchema := GetSchema(tableField.RelatedCollection)
		for _, tableLinkField := range tableSchema.GetLinkFields() {
			paths[tableLinkField.RelatedCollection] = append(paths[tableLinkField.RelatedCollection], fmt.Sprintf("%s.%s", tableField.Name, tableLinkField.Name))
		}
	}
	return
}
