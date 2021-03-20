package model

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
)

// Schema .
type Schema struct {
	ID          string
	CreatedAt   *gtime.Time
	UpdatedAt   *gtime.Time
	DeletedAt   *gtime.Time
	CreatedBy   string
	UpdatedBy   string
	Title       string
	Name        string
	ProjectName string
	Description string
	Fields      []*SchemaField
}

// SchemaField .
type SchemaField struct {
	ID                  string
	CreatedAt           *gtime.Time
	UpdatedAt           *gtime.Time
	DeletedAt           *gtime.Time
	CreatedBy           string
	UpdatedBy           string
	Pcn                 string
	Pid                 string
	Idx                 int
	Pfd                 string
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

// DefaultFieldNames : 所有默认的字段
// var DefaultFieldNames = g.SliceStr{"id", "pcn", "idx", "pid", "pfd", "created_at", "updated_at", "deleted_at", "created_by", "updated_by"}

// GetSchema : 获取collection的schema
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
	fieldNames := make([]string, 0)
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
		if field.IsPrivate || field.IsHidden || FieldType(field.Type) == Table {
			continue
		}
		fields = append(fields, field)
	}
	return fields
}

// GetRequiredFieldNames : 获取所有必填的字段名
func (schema *Schema) GetRequiredFieldNames() []string {
	fieldNames := make([]string, 0)
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
	fieldNames := make([]string, 0)
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

// GetLinkCollectionNames : 获取所有Link类型字段关联的collection
func (schema *Schema) GetLinkCollectionNames() []string {
	collectionNames := garray.NewStrArray()
	for _, field := range schema.GetLinkFields() {
		if !collectionNames.Contains(field.RelatedCollection) {
			collectionNames.Append(field.RelatedCollection)
		}
	}
	return collectionNames.Slice()
}

// GetLinkPathIncludeTableInner : 获取所有link字段的路径，包括子表
func (schema *Schema) GetLinkPathIncludeTableInner() (paths map[string][]string) {
	for _, linkField := range schema.GetLinkFields() {
		paths[linkField.RelatedCollection] = append(paths[linkField.RelatedCollection], linkField.Name)
	}
	for _, tableField := range schema.GetTableFields() {
		tableSchema, _ := GetSchema(tableField.RelatedCollection)
		for _, tableLinkField := range tableSchema.GetLinkFields() {
			paths[tableLinkField.RelatedCollection] = append(paths[tableLinkField.RelatedCollection], fmt.Sprintf("%s.%s", tableField.Name, tableLinkField.Name))
		}
	}
	return
}

// GetTableFieldNames : 获取所有Table类型字段名
func (schema *Schema) GetTableFieldNames() []string {
	fieldNames := make([]string, 0)
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

// GetTableFields : 获取所有Table类型字段
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

// GetTableCollectionNames : 获取所有Table类型字段关联的collection
func (schema *Schema) GetTableCollectionNames() []string {
	collectionNames := garray.NewStrArray()
	for _, field := range schema.GetTableFields() {
		if !collectionNames.Contains(field.RelatedCollection) {
			collectionNames.Append(field.RelatedCollection)
		}
	}
	return collectionNames.Slice()
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
