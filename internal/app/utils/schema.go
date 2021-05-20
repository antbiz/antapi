package utils

import (
	"fmt"

	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/gogf/gf/encoding/gjson"
)

// ParseFormRenderSchema 将 FormRender 的 schema 结构化
func ParseFormRenderSchema(data *gjson.Json) *dto.Schema {
	projectName := data.GetString("projectName")
	collectionName := data.GetString("name")
	schema := &dto.Schema{
		Title:          data.GetString("title"),
		CollectionName: collectionName,
		ProjectName:    projectName,
		Description:    data.GetString("description"),
		Fields:         make([]*dto.SchemaField, 0),
	}
	for fieldName := range data.GetMap("properties") {
		field := &dto.SchemaField{
			Name:         fieldName,
			IsRequired:   data.GetBool(fmt.Sprintf("properties.%s.required", fieldName)),
			DisplayName:  data.GetString(fmt.Sprintf("properties.%s.title", fieldName)),
			Description:  data.GetString(fmt.Sprintf("properties.%s.description", fieldName)),
			IsHidden:     data.GetBool(fmt.Sprintf("properties.%s.hidden", fieldName)),
			IsReadOnly:   data.GetBool(fmt.Sprintf("properties.%s.readOnly", fieldName)),
			IsUnique:     data.GetBool(fmt.Sprintf("properties.%s.unique", fieldName)),
			IsPrivate:    data.GetBool(fmt.Sprintf("properties.%s.private", fieldName)),
			IsIndexField: data.GetBool(fmt.Sprintf("properties.%s.index", fieldName)),
			Default:      data.GetString(fmt.Sprintf("properties.%s.default", fieldName)),
		}
		schema.Fields = append(schema.Fields, field)
	}

	return schema
}
