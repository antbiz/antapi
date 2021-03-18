package model

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gvalid"
)

// Schema .
type Schema struct {
	ID          string
	Title       string
	Name        string
	ProjectName string
	Description string
	Fields      []*SchemaField
}

// SchemaField .
type SchemaField struct {
	ID                  string
	Index               int
	Pid                 string
	Type                string
	Title               string
	Name                string
	Description         string
	IsRequired          bool
	IsHidden            bool
	IsUnique            bool
	IsPrivate           bool
	IsEncrypted         bool
	CanSort             bool
	CanIndex            bool
	IsMultiple          bool
	Default             string
	RelatedCollection   string
	RelatedDisplayField string
	Min                 int
	Max                 int
	Validator           string
	Style               string
	EnumType            string
	EnumOptions         string
}

// FieldType .
type FieldType string

// Supported field types
const (
	String    FieldType = "String"
	Color     FieldType = "Color"
	Email     FieldType = "Email"
	Phone     FieldType = "Phone"
	URL       FieldType = "Url"
	Password  FieldType = "Password"
	Text      FieldType = "Text"
	LongText  FieldType = "LongText"
	RichText  FieldType = "RichText"
	Markdown  FieldType = "Markdown"
	Code      FieldType = "Code"
	HTML      FieldType = "HTML"
	Signature FieldType = "Signature"
	Media     FieldType = "Media"
	Enum      FieldType = "Enum"
	JSON      FieldType = "JSON"
	UUID      FieldType = "UUID"
	Int       FieldType = "Int"
	BigInt    FieldType = "BigInt"
	Float     FieldType = "Float"
	Money     FieldType = "Money"
	Date      FieldType = "Date"
	DateTime  FieldType = "DateTime"
	Time      FieldType = "Time"
	TimeStamp FieldType = "TimeStamp"
	Year      FieldType = "Year"
	Bool      FieldType = "Bool"
	Link      FieldType = "Link"
	Table     FieldType = "Table"
)

// TODO: 缓存
func GetSchema(collectionName string) (*Schema, error) {
	db := g.DB()

	var schema *Schema
	if err := db.Table("schema").Scan(&schema, "name", collectionName); err != nil {
		return nil, err
	}
	if err := db.Table("schema_field").Order("idx asc").ScanList(&schema.Fields, "pid", schema.ID); err != nil {
		return nil, err
	}
	return schema, nil
}

// GetPublicFieldNames : 获取所有对外开放的字段名
func (schema *Schema) GetPublicFieldNames() []string {
	var fieldNames []string
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden || FieldType(field.Type) == Table {
			continue
		}
		fieldNames = append(fieldNames, field.Name)
	}
	return fieldNames
}

// GetPublicFields : 获取所有对外开放的字段
func (schema *Schema) GetPublicFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		fields = append(fields, field)
	}
	return fields
}

// GetRequiredFieldNames : 获取所有必填的字段名
func (schema *Schema) GetRequiredFieldNames() []string {
	var fieldNames []string
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if field.IsRequired {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

// GetRequiredFields : 获取所有必填的字段
func (schema *Schema) GetRequiredFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if field.IsRequired {
			fields = append(fields, field)
		}
	}
	return fields
}

// GetLinkFieldNames : 获取所有Link类型字段名
func (schema *Schema) GetLinkFieldNames() []string {
	var fieldNames []string
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if len(field.RelatedCollection) > 0 && FieldType(field.Type) == Link {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

// GetLinkFields : 获取所有Link类型字段
func (schema *Schema) GetLinkFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if len(field.RelatedCollection) > 0 && FieldType(field.Type) == Link {
			fields = append(fields, field)
		}
	}
	return fields
}

// GetLinkFieldNames : 获取所有Table类型字段名
func (schema *Schema) GetTableFieldNames() []string {
	var fieldNames []string
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if len(field.RelatedCollection) > 0 && FieldType(field.Type) == Table {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

// GetLinkFields : 获取所有Table类型字段
func (schema *Schema) GetTableFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if len(field.RelatedCollection) > 0 && FieldType(field.Type) == Table {
			fields = append(fields, field)
		}
	}
	return fields
}

// CheckFieldValue : 校验字段值的合法性
// TODO: 错误信息支持多语言
func (field *SchemaField) CheckFieldValue(value interface{}) *gvalid.Error {
	var (
		err  *gvalid.Error
		rule string
		msg  string = fmt.Sprintf("%s格式不正确：%v", field.Title, value)
	)

	// 检验必填
	if field.IsRequired {
		if err = gvalid.Check(value, "required", fmt.Sprintf("%s不能为空", field.Title)); err != nil {
			return err
		}
	}

	// 通过字段类型做校验
	switch FieldType(field.Type) {
	case Email:
		rule = "email"
	case Phone:
		rule = "phone"
	case URL:
		rule = "url"
	case Date:
		rule = "date"
	case JSON:
		rule = "json"
	case Int:
		rule = "integer"
	case Float:
		rule = "float"
	case Bool:
		rule = "boolean"
	case Enum:
		var enumOptions []string
		if len(field.EnumOptions) > 0 {
			if j, err := gjson.LoadContent(fmt.Sprintf(`{"options":%s}`, field.EnumOptions)); err == nil {
				for i := 0; i < len(j.GetArray("options")); i++ {
					enumOptions = append(enumOptions, j.GetString(fmt.Sprintf("%d.value", i)))
				}
			}
		}
		enumOptionsStr := strings.Join(enumOptions, ",")
		rule = fmt.Sprintf("in:%s", enumOptionsStr)
		msg = fmt.Sprintf("%s不存在于%s", field.Title, enumOptionsStr)
	}
	if len(rule) > 0 {
		if err = gvalid.Check(value, rule, msg); err != nil {
			return err
		}
	}

	// 后台配置的校验规则
	if len(field.Validator) > 0 {
		return gvalid.Check(value, field.Validator, msg)
	}

	return nil
}
