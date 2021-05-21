package dto

import (
	"fmt"

	"github.com/gogf/gf/util/gvalid"
)

// Schema .
type Schema struct {
	Title          string
	CollectionName string
	ProjectName    string
	Description    string
	Schema         map[string]interface{}
	Column         int
	Fields         []*SchemaField
}

// SchemaField .
type SchemaField struct {
	DisplayName       string
	Name              string
	Description       string
	IsRequired        bool
	IsHidden          bool
	IsReadOnly        bool
	IsUnique          bool
	IsPrivate         bool
	IsIndexField      bool
	Default           string
	ConnectCollection string
	ConnectField      string
	ConnectMany       bool
	Validator         string
	EnumOptions       *EnumOption
	EnumType          string
}

// EnumOption .
type EnumOption struct {
	Labels []string
	Values []string
}

// DefaultFieldNames : 所有默认的字段
var DefaultFieldNames = []string{"_id", "createdAt", "updatedAt", "createdBy", "updatedBy"}

// GetFieldNames 获取字段名
// includeHidden: 包括 is_hidden 的字段
// includePrivate: 包括 is_private 的字段
func (schema *Schema) GetFieldNames(includeHidden, includePrivate bool) []string {
	fieldNames := make([]string, 0)
	for _, field := range schema.Fields {
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

// CheckFieldValue : 校验字段值的合法性
// TODO: 错误信息支持多语言
func (field *SchemaField) CheckFieldValue(value interface{}) *gvalid.Error {
	var (
		err  *gvalid.Error
		rule string
		msg  string = fmt.Sprintf("%s格式不正确：%v", field.DisplayName, value)
	)

	// 检验必填
	if field.IsRequired {
		if err = gvalid.Check(value, "required", fmt.Sprintf("%s不能为空", field.DisplayName)); err != nil {
			return err
		}
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
