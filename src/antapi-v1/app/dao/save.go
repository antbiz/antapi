package dao

import (
	"antapi/app/global"
)

// Save 更新或创建单个数据
func Save(collectionName string, arg *SaveFuncArg, data interface{}) (string, error) {
	dataGJson := dataToJson(data)
	schema := global.GetSchema(collectionName)

	// 自动查找唯一列去获取唯一id，如果不存在则新建
	// 查找条件：is_unique为true并且data中的值不能为空
	var id string
	for _, fieldName := range schema.GetUniqueFieldNames() {
		val := dataGJson.GetString(fieldName)
		if val != "" {
			getFunArg := &GetFuncArg{
				Where:             fieldName,
				WhereArgs:         val,
				Fields:            []string{"id"},
				IgnoreFieldsCheck: true,
			}
			if record, err := Get(collectionName, getFunArg); err != nil {
				return "", err
			} else {
				id = record.GetString("id")
			}
		}
	}

	if id == "" {
		updateArg := &UpdateFuncArg{
			IgnoreFieldValueCheck: arg.IgnoreFieldValueCheck,
			IncludeHiddenField:    arg.IncludeHiddenField,
			IncludePrivateField:   arg.IncludePrivateField,
		}
		return id, Update(collectionName, updateArg, id, data)
	} else {
		insertArg := &InsertFuncArg{
			IgnoreFieldValueCheck: arg.IgnoreFieldValueCheck,
			IncludeHiddenField:    arg.IncludeHiddenField,
			IncludePrivateField:   arg.IncludePrivateField,
		}
		return Insert(collectionName, insertArg, data)
	}
}
