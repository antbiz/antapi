package logic

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
)

type SchemaLogic struct{}

var DefaultSchemaLogic = SchemaLogic{}

// CheckFields : 校验collection的字段，并填充系统必要字段
func (SchemaLogic) CheckFields(data *gjson.Json) error {
	fieldsLen := len(data.GetArray("fields"))
	if fieldsLen == 0 {
		return errors.New("fields is required")
	}

	isChildTable := data.GetBool("is_child")

	var (
		hasIdField  bool
		hasPcnField bool
		hasIdxField bool
		hasPidField bool
		hasPfdField bool
	)

	getDataPathForField := func(i int, name string) string {
		return fmt.Sprintf("fields.%d.%s", i, name)
	}

	for i := 0; i < fieldsLen; i++ {
		fieldName := getDataPathForField(i, "name")
		switch fieldName {
		case "id":
			hasIdField = true
		case "pcn":
			hasPcnField = true
		case "idx":
			hasIdxField = true
		case "pid":
			hasPidField = true
		case "pfd":
			hasPfdField = true
		}
	}

	updateField := func(i int, name string, v interface{}) {
		data.Set(fmt.Sprintf("fields.%d.%s", i, name), v)
	}

	if !hasIdField {
		fieldsLen++
		updateField(fieldsLen, "type", "UUID")
		updateField(fieldsLen, "title", "ID")
		updateField(fieldsLen, "name", "id")
	}
	if isChildTable {
		if !hasPcnField {
			fieldsLen++
			updateField(fieldsLen, "type", "String")
			updateField(fieldsLen, "title", "Parent Collection")
			updateField(fieldsLen, "name", "pcn")
			updateField(fieldsLen, "can_index", true)
		}
		if !hasIdxField {
			fieldsLen++
			updateField(fieldsLen, "type", "Int")
			updateField(fieldsLen, "title", "Index")
			updateField(fieldsLen, "name", "idx")
		}
		if !hasPidField {
			fieldsLen++
			updateField(fieldsLen, "type", "String")
			updateField(fieldsLen, "title", "Parent ID")
			updateField(fieldsLen, "name", "pid")
			updateField(fieldsLen, "can_index", true)
		}
		if !hasPfdField {
			fieldsLen++
			updateField(fieldsLen, "type", "String")
			updateField(fieldsLen, "title", "Parent Field")
			updateField(fieldsLen, "name", "pfd")
			updateField(fieldsLen, "can_index", true)
		}
	}

	return nil
}
