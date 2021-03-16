package db

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
	ID                string
	Index             int
	Pid               string
	Type              string
	Title             string
	Name              string
	Description       string
	IsRequired        bool
	IsHidden          bool
	IsUnique          bool
	IsPrivate         bool
	IsEncrypted       bool
	CanSort           bool
	Default           string
	ConnectCollection string
	ConnectField      string
	ConnectMany       bool
	Min               int
	Max               int
	Validator         string
	Style             string
	EnumType          string
	EnumOptions       string
	IsMultiple        bool
}

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

func (schema *Schema) GetOneToOneFieldNames() []string {
	var fieldNames []string
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if len(field.ConnectCollection) > 0 && len(field.ConnectField) > 0 && !field.ConnectMany {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

func (schema *Schema) GetOneToOneFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if len(field.ConnectCollection) > 0 && len(field.ConnectField) > 0 && !field.ConnectMany {
			fields = append(fields, field)
		}
	}
	return fields
}

func (schema *Schema) GetOneToManyFieldNames() []string {
	var fieldNames []string
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if len(field.ConnectCollection) > 0 && len(field.ConnectField) > 0 && field.ConnectMany {
			fieldNames = append(fieldNames, field.Name)
		}
	}
	return fieldNames
}

func (schema *Schema) GetOneToManyFields() []*SchemaField {
	fields := make([]*SchemaField, 0)
	for _, field := range schema.Fields {
		if field.IsPrivate || field.IsHidden {
			continue
		}
		if len(field.ConnectCollection) > 0 && len(field.ConnectField) > 0 && field.ConnectMany {
			fields = append(fields, field)
		}
	}
	return fields
}
