package dto

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/util/gvalid"
)

// Schema .
type Schema struct {
	Name        string
	Title       string
	ProjectName string
	Description string
	Fields      []*SchemaField
}

// SchemaField .
type SchemaField struct {
	Name              string
	Title             string
	IsRequired        bool
	IsHidden          bool
	IsReadOnly        bool
	IsUnique          bool
	IsPrivate         bool
	IsIndexField      bool
	IsSysField        bool
	Default           string
	ConnectCollection string
	ConnectField      string
	ConnectMany       bool
	Validator         string
	Enum              []string
	EnumNames         []string
	Description       string
}

// GetFieldNames 获取字段名
// includeHidden: 包括 is_hidden 的字段
// includePrivate: 包括 is_private 的字段
func (schema *Schema) GetFieldNames(includeHidden, includePrivate bool) []string {
	fieldNames := make([]string, 0)
	for _, field := range schema.Fields {
		if !field.IsSysField {
			if field.IsHidden && !includeHidden {
				continue
			}
			if field.IsPrivate && !includePrivate {
				continue
			}
		}
		fieldNames = append(fieldNames, field.Name)
	}
	return fieldNames
}

// GetFields : 获取字段
func (schema *Schema) GetFields(includeHidden, includePrivate bool) []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if !field.IsSysField {
			if field.IsHidden && !includeHidden {
				continue
			}
			if field.IsPrivate && !includePrivate {
				continue
			}
		}
		fields = append(fields, field)
	}
	return fields
}

// GetHiddenFieldNames : 获取所有隐藏字段名
func (schema *Schema) GetHiddenFieldNames() []string {
	fieldNames := make([]string, 0)
	for _, field := range schema.Fields {
		if !field.IsSysField && field.IsHidden {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

// GetHiddenFields : 获取所有隐藏字段
func (schema *Schema) GetHiddenFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if !field.IsSysField && field.IsHidden {
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
		msg  string = fmt.Sprintf("%s格式不正确：%v", field.Title, value)
	)

	// 检验必填
	if field.IsRequired {
		if err = gvalid.Check(value, "required", fmt.Sprintf("%s不能为空", field.Title)); err != nil {
			return err
		}
	}

	// 校验选项值
	if len(field.Enum) > 0 {
		rule = fmt.Sprintf("in:%s", strings.Join(field.Enum, ","))
		msg = fmt.Sprintf("%s支持的选项: %s", field.Title, strings.Join(field.EnumNames, ","))
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
