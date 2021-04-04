package utils

import (
	"antapi/app/model"
	"antapi/app/model/fieldtype"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
)

// ParseFormRenderSchema 将 FormRender 的 schema 结构化
func ParseFormRenderSchema(jsonSchema *gjson.Json) (schemas []*model.Schema) {
	schemas = make([]*model.Schema, 0)

	processObject := func(data *gjson.Json, projectName, collectionName string) *model.Schema {
		schema := &model.Schema{
			Title:          data.GetString("title"),
			CollectionName: collectionName,
			ProjectName:    projectName,
			Description:    data.GetString("description"),
			Fields:         make([]*model.SchemaField, 0),
		}
		for fieldName := range data.GetMap("properties") {
			field := &model.SchemaField{
				Name:        fieldName,
				IsRequired:  data.GetBool(fmt.Sprintf("properties.%s.required", fieldName)),
				Title:       data.GetString(fmt.Sprintf("properties.%s.title", fieldName)),
				Description: data.GetString(fmt.Sprintf("properties.%s.description", fieldName)),
				IsHidden:    data.GetBool(fmt.Sprintf("properties.%s.ui:hidden", fieldName)),
				IsReadOnly:  data.GetBool(fmt.Sprintf("properties.%s.ui:readonly", fieldName)),
				IsUnique:    data.GetBool(fmt.Sprintf("properties.%s.unique", fieldName)),
				IsPrivate:   data.GetBool(fmt.Sprintf("properties.%s.private", fieldName)),
				CanIndex:    data.GetBool(fmt.Sprintf("properties.%s.index", fieldName)),
				Default:     data.GetString(fmt.Sprintf("properties.%s.default", fieldName)),
			}

			fieldType := data.GetString(fmt.Sprintf("properties.%s.type", fieldName))
			fieldFormat := data.GetString(fmt.Sprintf("properties.%s.format", fieldName))
			switch fmt.Sprintf("%s:%s", fieldType, fieldFormat) {
			case "string:":
				field.Type = fieldtype.String
			case "string:textarea":
				field.Type = fieldtype.Text
			case "string:richtext":
				field.Type = fieldtype.RichText
			case "string:markdown":
				field.Type = fieldtype.Markdown
			case "string:image":
				field.Type = fieldtype.Image
			case "string:email":
				field.Type = fieldtype.Email
			case "string:phone":
				field.Type = fieldtype.Phone
			case "string:password":
				field.Type = fieldtype.Password
			case "string:url":
				field.Type = fieldtype.URL
			case "string:dateTime":
				field.Type = fieldtype.DateTime
			case "string:date":
				field.Type = fieldtype.Date
			case "string:time":
				field.Type = fieldtype.Time
			case "string:year":
				field.Type = fieldtype.Year
			case "string:upload":
				field.Type = fieldtype.File
			case "string:auto":
				field.Type = fieldtype.AutoComplete
			case "string:link":
				field.Type = fieldtype.Link
			case "number:int":
				field.Type = fieldtype.Int
			case "number:float":
				field.Type = fieldtype.Float
			case "number:money":
				field.Type = fieldtype.Money
			case "boolean:":
				field.Type = fieldtype.Bool
			case "object:", "array:", "range:", "range:dateTime":
				field.Type = fieldtype.JSON
			case "table:":
				field.Type = fieldtype.Table
			case "html:":
				field.Type = fieldtype.HTML
			default:
				field.Type = fieldtype.Text
			}

			if fieldType == "table" {
				field.RelatedCollection = fmt.Sprintf("%s_%s", collectionName, fieldName)
			}

			schema.Fields = append(schema.Fields, field)
		}

		return schema
	}

	projectName := jsonSchema.GetString("project_name")
	collectionName := jsonSchema.GetString("name")
	schemas = append(schemas, processObject(jsonSchema.GetJson("schema"), projectName, collectionName))
	for fieldName := range jsonSchema.GetMap("schema.properties") {
		fieldType := jsonSchema.GetString(fmt.Sprintf("schema.properties.%s.type", fieldName))
		if fieldType == "table" {
			schemas = append(schemas, processObject(
				jsonSchema.GetJson(fmt.Sprintf("schema.properties.%s.items", fieldName)),
				projectName,
				fmt.Sprintf("%s_%s", collectionName, fieldName),
			))
		}
	}

	return schemas
}
