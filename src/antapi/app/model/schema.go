package model

import (
	"antapi/app/model/fieldtype"
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
	ID          string         `orm:"id"`
	CreatedAt   *gtime.Time    `orm:"created_at"`
	UpdatedAt   *gtime.Time    `orm:"updated_at"`
	DeletedAt   *gtime.Time    `orm:"deleted_at"`
	CreatedBy   string         `orm:"created_by"`
	UpdatedBy   string         `orm:"updated_by"`
	Title       string         `orm:"title"`
	Name        string         `orm:"name"`
	ProjectName string         `orm:"project_name"`
	Description string         `orm:"description"`
	Fields      []*SchemaField `orm:"-"`
}

// SchemaField .
type SchemaField struct {
	ID                  string      `orm:"id"`
	CreatedAt           *gtime.Time `orm:"created_at"`
	UpdatedAt           *gtime.Time `orm:"updated_at"`
	DeletedAt           *gtime.Time `orm:"deleted_at"`
	CreatedBy           string      `orm:"created_by"`
	UpdatedBy           string      `orm:"updated_by"`
	Pcn                 string      `orm:"pcn"`
	Pid                 string      `orm:"pid"`
	Idx                 int         `orm:"idx"`
	Pfd                 string      `orm:"pfd"`
	Type                string      `orm:"type"`
	Title               string      `orm:"title"`
	Name                string      `orm:"name"`
	Description         string      `orm:"description"`
	IsRequired          bool        `orm:"is_required"`
	IsHidden            bool        `orm:"is_hidden"`
	IsUnique            bool        `orm:"is_unique"`
	IsPrivate           bool        `orm:"is_private"`
	IsEncrypted         bool        `orm:"is_encrypted"`
	CanSort             bool        `orm:"can_sort"`
	CanIndex            bool        `orm:"can_index"`
	IsMultiple          bool        `orm:"is_multiple"`
	Default             string      `orm:"default"`
	RelatedCollection   string      `orm:"related_collection"`
	RelatedDisplayField string      `orm:"related_display_field"`
	Min                 int         `orm:"min"`
	Max                 int         `orm:"max"`
	Validator           string      `orm:"validator"`
	Style               string      `orm:"style"`
	EnumType            string      `orm:"enum_type"`
	EnumOptions         string      `orm:"enum_options"`
}

// DefaultFieldNames : 所有默认的字段
var DefaultFieldNames = g.SliceStr{"id", "pcn", "idx", "pid", "pfd", "created_at", "updated_at", "deleted_at", "created_by", "updated_by"}

// GetFieldNames 获取字段名
// includeHidden: 包括 is_hidden 的字段
// includePrivate: 包括 is_private 的字段
func (schema *Schema) GetFieldNames(includeHidden, includePrivate bool) []string {
	fieldNames := make([]string, 0)
	for _, field := range schema.Fields {
		if fieldtype.FieldType(field.Type) == fieldtype.Table {
			continue
		}
		if field.IsHidden && !includeHidden {
			continue
		}
		if field.IsPrivate && !includePrivate {
			continue
		}
		fieldNames = append(fieldNames, field.Name)
	}
	return fieldNames
}

// GetFields : 获取字段
func (schema *Schema) GetFields(includeHidden, includePrivate bool) []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if fieldtype.FieldType(field.Type) == fieldtype.Table {
			continue
		}
		if field.IsHidden && !includeHidden {
			continue
		}
		if field.IsPrivate && !includePrivate {
			continue
		}
		fields = append(fields, field)
	}
	return fields
}

// GetHiddenFieldNames : 获取所有隐藏字段名
func (schema *Schema) GetHiddenFieldNames() []string {
	fieldNames := make([]string, 0)
	for _, field := range schema.Fields {
		if fieldtype.FieldType(field.Type) == fieldtype.Table {
			continue
		}
		if field.IsHidden {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

// GetHiddenFields : 获取所有隐藏字段
func (schema *Schema) GetHiddenFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if fieldtype.FieldType(field.Type) == fieldtype.Table {
			continue
		}
		if field.IsHidden {
			fields = append(fields, field)
		}
	}
	return fields
}

// GetUniqueFieldNames : 获取所有私密字段名
func (schema *Schema) GetUniqueFieldNames() []string {
	fieldNames := make([]string, 0)
	for _, field := range schema.Fields {
		if fieldtype.FieldType(field.Type) == fieldtype.Table {
			continue
		}
		if field.IsUnique {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

// GetUniqueFields : 获取所有私密字段
func (schema *Schema) GetUniqueFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if fieldtype.FieldType(field.Type) == fieldtype.Table {
			continue
		}
		if field.IsUnique {
			fields = append(fields, field)
		}
	}
	return fields
}

// GetPrivateFieldNames : 获取所有私密字段名
func (schema *Schema) GetPrivateFieldNames() []string {
	fieldNames := make([]string, 0)
	for _, field := range schema.Fields {
		if fieldtype.FieldType(field.Type) == fieldtype.Table {
			continue
		}
		if field.IsPrivate {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

// GetPrivateFields : 获取所有私密字段
func (schema *Schema) GetPrivateFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if fieldtype.FieldType(field.Type) == fieldtype.Table {
			continue
		}
		if field.IsPrivate {
			fields = append(fields, field)
		}
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
		if field.RelatedCollection != "" && fieldtype.FieldType(field.Type) == fieldtype.Link {
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
		if field.RelatedCollection != "" && fieldtype.FieldType(field.Type) == fieldtype.Link {
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

// GetTableFieldNames : 获取所有Table类型字段名
func (schema *Schema) GetTableFieldNames() []string {
	fieldNames := make([]string, 0)
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if field.RelatedCollection != "" && fieldtype.FieldType(field.Type) == fieldtype.Table {
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
		if field.RelatedCollection != "" && fieldtype.FieldType(field.Type) == fieldtype.Table {
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
	switch fieldtype.FieldType(field.Type) {
	case fieldtype.Email:
		rule = "email"
	case fieldtype.Phone:
		rule = "phone"
	case fieldtype.URL:
		rule = "url"
	case fieldtype.Date:
		rule = "date"
	case fieldtype.JSON:
		rule = "json"
	case fieldtype.Int:
		rule = "integer"
	case fieldtype.Float:
		rule = "float"
	case fieldtype.Bool:
		rule = "boolean"
	case fieldtype.Enum:
		var enumOptions []string
		if field.EnumOptions != "" {
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
	if rule != "" {
		if err = gvalid.Check(value, rule, msg); err != nil {
			return err
		}
	}

	// 后台配置的校验规则
	if field.Validator != "" {
		return gvalid.Check(value, field.Validator, msg)
	}

	return nil
}
