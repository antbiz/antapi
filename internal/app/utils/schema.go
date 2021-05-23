package utils

import (
	"fmt"

	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/encoding/gjson"
)

// ParseFormRenderSchema 将 FormRender 的 schema 结构化
func ParseFormRenderSchema(data *gjson.Json) *dto.Schema {
	defaultSysFieldNames := gset.NewStrSetFrom(
		[]string{
			"_id", "createdAt", "updatedAt",
			"createdBy", "updatedBy",
		},
	)
	existsSysFieldNames := gset.NewStrSet()

	schema := &dto.Schema{
		Name:        data.GetString("name"),
		Title:       data.GetString("title"),
		ProjectName: data.GetString("projectName"),
		Description: data.GetString("description"),
		Fields:      make([]*dto.SchemaField, 0),
	}
	for fieldName := range data.GetMap("properties") {
		field := &dto.SchemaField{
			Name:              fieldName,
			Title:             data.GetString(fmt.Sprintf("properties.%s.title", fieldName)),
			IsRequired:        data.GetBool(fmt.Sprintf("properties.%s.required", fieldName)),
			IsHidden:          data.GetBool(fmt.Sprintf("properties.%s.hidden", fieldName)),
			IsReadOnly:        data.GetBool(fmt.Sprintf("properties.%s.readOnly", fieldName)),
			IsUnique:          data.GetBool(fmt.Sprintf("properties.%s.unique", fieldName)),
			IsPrivate:         data.GetBool(fmt.Sprintf("properties.%s.private", fieldName)),
			IsIndexField:      data.GetBool(fmt.Sprintf("properties.%s.index", fieldName)),
			Default:           data.GetString(fmt.Sprintf("properties.%s.default", fieldName)),
			ConnectCollection: data.GetString(fmt.Sprintf("properties.%s.connectCollection", fieldName)),
			ConnectField:      data.GetString(fmt.Sprintf("properties.%s.connectField", fieldName)),
			ConnectMany:       data.GetBool(fmt.Sprintf("properties.%s.connectMany", fieldName)),
			Validator:         data.GetString(fmt.Sprintf("properties.%s.validator", fieldName)),
			Description:       data.GetString(fmt.Sprintf("properties.%s.description", fieldName)),
		}
		field.Enum = data.GetStrings(fmt.Sprintf("properties.%s.enum", fieldName))
		field.EnumNames = data.GetStrings(fmt.Sprintf("properties.%s.enumNames", fieldName), field.Enum)
		if defaultSysFieldNames.Contains(fieldName) {
			field.IsSysField = true
			existsSysFieldNames.AddIfNotExist(fieldName)
		}
		schema.Fields = append(schema.Fields, field)
	}

	for _, fieldName := range defaultSysFieldNames.Diff(existsSysFieldNames).Slice() {
		schema.Fields = append(schema.Fields, &dto.SchemaField{
			Name:       fieldName,
			Title:      fieldName,
			IsSysField: true,
		})
	}

	return schema
}
