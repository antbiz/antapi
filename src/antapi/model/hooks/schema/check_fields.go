package schema

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
)

// CheckFields : 校验collection的字段，并填充系统必要字段
func CheckFields(data *gjson.Json) error {
	fieldsLen := len(data.GetArray("fields"))
	if fieldsLen == 0 {
		return errors.New("fields is required")
	}

	isChildTable := data.GetBool("is_child")

	var (
		hasIdField  bool
		hasIdxField bool
		hasPidField bool
		hasPcnField bool
	)

	getDataPathForField := func(i int, name string) string {
		return fmt.Sprintf("fields.%d.%s", i, name)
	}

	for i := 0; i < fieldsLen; i++ {
		fieldName := getDataPathForField(i, "name")
		switch fieldName {
		case "id":
			hasIdField = true
		case "idx":
			hasIdxField = true
		case "pid":
			hasPidField = true
		case "pcn":
			hasPcnField = true
		}
	}

	updateField := func(i int, name string, v interface{}) {
		data.Set(fmt.Sprintf("fields.%d.%s", i, name), v)
	}

	if !hasIdField {
		fieldsLen += 1
		updateField(fieldsLen, "type", "UUID")
		updateField(fieldsLen, "title", "ID")
		updateField(fieldsLen, "name", "id")
	}
	if isChildTable {
		if !hasIdxField {
			fieldsLen += 1
			updateField(fieldsLen, "type", "Int")
			updateField(fieldsLen, "title", "Index")
			updateField(fieldsLen, "name", "idx")
		}
		if !hasPidField {
			fieldsLen += 1
			updateField(fieldsLen, "type", "String")
			updateField(fieldsLen, "title", "Parent ID")
			updateField(fieldsLen, "name", "pid")
			updateField(fieldsLen, "can_index", true)
		}
		if !hasPcnField {
			fieldsLen += 1
			updateField(fieldsLen, "type", "String")
			updateField(fieldsLen, "title", "Parent Collection")
			updateField(fieldsLen, "name", "pcn")
			updateField(fieldsLen, "can_index", true)
		}
	}

	return nil
}
