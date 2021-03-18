package model

import "github.com/gogf/gf/frame/g"

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
	Tel       FieldType = "Tel"
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
	Array     FieldType = "Array"
	Link      FieldType = "Link"
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

// GetFieldNames : 获取所有对外开放的字段名
func (schema *Schema) GetFieldNames() []string {
	var fieldNames []string
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		fieldNames = append(fieldNames, field.Name)
	}
	return fieldNames
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

// GetLinkFieldNames : 获取所有link字段名
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

// GetLinkFields : 获取所有link字段
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
