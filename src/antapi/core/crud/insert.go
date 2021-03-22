package crud

import (
	"antapi/common/errcode"
	"antapi/hooks"
	"antapi/logic"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/guid"
)

// Insert : 插入单个数据，返回插入的主体id
func Insert(collectionName string, data interface{}) (string, error) {
	res, err := InsertList(collectionName, data)
	if err != nil {
		return "", err
	}
	return res[0], nil
}

// InsertList : 插入多个数据，返回一组插入的主体id
// TODO: 需要考虑子表数据校验的提示信息
func InsertList(collectionName string, data ...interface{}) ([]string, error) {
	dataLen := len(data)
	if dataLen == 0 {
		return nil, nil
	}
	db := g.DB()
	schema := logic.GetSchema(collectionName)

	ids := make([]string, 0, dataLen)

	dataGJsonSlice := make([]*gjson.Json, 0, dataLen)
	for i := 0; i < dataLen; i++ {
		dataGJson, err := gjson.LoadJson(data[i])
		if err != nil {
			return nil, gerror.WrapCode(errcode.JSONError, err, errcode.JSONErrorMsg)
		}
		dataGJsonSlice = append(dataGJsonSlice, dataGJson)
	}

	// 执行 BeforeInsertHooks, BeforeSaveHooks 勾子
	for _, dataGJson := range dataGJsonSlice {
		for _, hook := range hooks.GetBeforeInsertHooksByCollectionName(collectionName) {
			if err := hook(dataGJson); err != nil {
				return nil, err
			}
		}
		for _, hook := range hooks.GetBeforeSaveHooksByCollectionName(collectionName) {
			if err := hook(dataGJson); err != nil {
				return nil, err
			}
		}
	}

	// 批量插入主体数据
	contents := make([]map[string]interface{}, 0, dataLen)
	for i := 0; i < dataLen; i++ {
		id := guid.S()
		ids = append(ids, id)
		dataGJson := dataGJsonSlice[i]
		dataGJson.Set("id", id)

		var content map[string]interface{}
		for _, field := range schema.GetPublicFields() {
			val := dataGJson.Get(field.Name)
			if validErr := field.CheckFieldValue(val); validErr != nil {
				return nil, gerror.NewCode(errcode.ParameterBindError, validErr.String())
			}
			content[field.Name] = val
		}

		contents = append(contents, content)
	}
	if _, err := db.Table(collectionName).Insert(contents); err != nil {
		return nil, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}

	// 批量插入子表数据
	for _, field := range schema.GetTableFields() {
		tableContent := make([]map[string]interface{}, 0)
		for i := 0; i < dataLen; i++ {
			dataGJson := dataGJsonSlice[i]
			tableRowsLen := len(dataGJson.GetArray(field.Name))
			if tableRowsLen == 0 {
				continue
			}
			tableSchema := logic.GetSchema(field.RelatedCollection)

			for j := 0; j < tableRowsLen; j++ {
				var tableRowContent map[string]interface{}
				for _, tableField := range tableSchema.GetPublicFields() {
					val := dataGJson.Get(fmt.Sprintf("%s.%d.%s", field.Name, j, tableField.Name))
					if validErr := field.CheckFieldValue(val); validErr != nil {
						return nil, gerror.NewCode(errcode.ParameterBindError, validErr.String())
					}
					tableRowContent[tableField.Name] = val
				}
				tableRowContent["pcn"] = collectionName
				tableRowContent["id"] = guid.S()
				tableRowContent["idx"] = j
				tableRowContent["pid"] = ids[i]
				tableRowContent["pfd"] = field.Name

				// 更新 dataGJson 方便 执行 AfterInsertHooks 勾子的一些业务逻辑
				for _, defaultField := range []string{"pcn", "id", "idx", "pid", "pfd"} {
					dataGJson.Set(fmt.Sprintf("%s.%d.%s", field.Name, j, defaultField), tableRowContent[defaultField])
				}

				tableContent = append(tableContent, tableRowContent)
			}

			dataGJsonSlice[i] = dataGJson
		}

		if _, err := db.Table(field.RelatedCollection).Insert(tableContent); err != nil {
			return nil, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
		}
	}

	// 执行 AfterInsertHooks, AfterSaveHooks 勾子
	for _, dataGJson := range dataGJsonSlice {
		for _, hook := range hooks.GetAfterInsertHooksByCollectionName(collectionName) {
			if err := hook(dataGJson); err != nil {
				return nil, err
			}
		}
		for _, hook := range hooks.GetAfterSaveHooksByCollectionName(collectionName) {
			if err := hook(dataGJson); err != nil {
				return nil, err
			}
		}
	}

	return ids, nil
}
