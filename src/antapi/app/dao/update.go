package dao

import (
	"antapi/app/global"
	"antapi/common/errcode"
	"fmt"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/guid"
)

// Update : 更新单个数据
func Update(collectionName string, arg *UpdateFuncArg, id string, data interface{}) error {
	db := g.DB()
	dataGJson := dataToJson(data)
	schema := global.GetSchema(collectionName)

	// 检查数据的修改时间，如果提交数据中的修改时间要比数据库中的值小则不允许提交
	oriRes, err := db.Table(collectionName).Fields("id", "updated_at").Where("id", id).One()
	if err != nil {
		return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}
	if oriRes.IsEmpty() {
		if arg.RaiseNotFound {
			return gerror.NewCode(errcode.SourceNotFound, errcode.SourceNotFoundMsg)
		}
	} else {
		oriUpdatedAtGTimeObj := oriRes.GMap().GetVar("updated_at").GTime()
		curUpdatedAtGTimeObj := dataGJson.GetGTime("updated_at")
		if oriUpdatedAtGTimeObj != nil && curUpdatedAtGTimeObj != nil && oriUpdatedAtGTimeObj.Second() > curUpdatedAtGTimeObj.Second() {
			return gerror.NewCode(errcode.TimestampMismatchError, errcode.TimestampMismatchErrorMsg)
		}
	}

	if err := CheckDuplicate(collectionName, dataGJson, id); err != nil {
		return err
	}

	// 执行 BeforeInsertHooks, BeforeSaveHooks 勾子
	for _, hook := range global.GetBeforeUpdateHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return err
		}
	}
	for _, hook := range global.GetBeforeSaveHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return err
		}
	}

	// 更新主体数据
	content := map[string]interface{}{
		"updated_by": arg.SessionUsername,
	}
	for _, field := range schema.GetFields(arg.IncludeHiddenField, arg.IncludePrivateField) {
		val := dataGJson.Get(field.Name)
		if !arg.IgnoreFieldValueCheck {
			if validErr := field.CheckFieldValue(val); validErr != nil {
				return gerror.NewCode(errcode.ParameterBindError, validErr.String())
			}
		}
		content[field.Name] = val
	}
	if _, err := db.Table(collectionName).FieldsEx("id,created_by,created_at,deleted_at").Where("id", id).Update(content); err != nil {
		return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}

	// 更新子表数据
	for _, field := range schema.GetTableFields() {
		tableRowsLen := len(dataGJson.GetArray(field.Name))
		if tableRowsLen == 0 {
			continue
		}
		tableContent := make([]map[string]interface{}, 0)
		tableSchema := global.GetSchema(field.RelatedCollection)

		tableIds := make([]string, tableRowsLen)
		for i := 0; i < tableRowsLen; i++ {
			dataGJson.Set(fmt.Sprintf("%s.%d.pcn", field.Name, i), collectionName)
			dataGJson.Set(fmt.Sprintf("%s.%d.idx", field.Name, i), i)
			dataGJson.Set(fmt.Sprintf("%s.%d.pid", field.Name, i), id)
			dataGJson.Set(fmt.Sprintf("%s.%d.pfd", field.Name, i), field.Name)
			tableRowId := dataGJson.GetString(fmt.Sprintf("%s.%d.%s", field.Name, i, "id"))
			if tableRowId == "" {
				tableRowId = guid.S()
				dataGJson.Set(fmt.Sprintf("%s.%d.id", field.Name, i), tableRowId)
			}
			tableIds = append(tableIds, tableRowId)

			tableRowContent := map[string]interface{}{
				"updated_by": arg.SessionUsername,
			}
			for _, tableField := range tableSchema.GetFields(arg.IncludeHiddenField, arg.IncludePrivateField) {
				val := dataGJson.Get(fmt.Sprintf("%s.%d.%s", field.Name, i, tableField.Name))
				if !arg.IgnoreFieldValueCheck {
					if validErr := field.CheckFieldValue(val); validErr != nil {
						return gerror.NewCode(errcode.ParameterBindError, validErr.String())
					}
				}
				tableRowContent[tableField.Name] = val
			}
			tableContent = append(tableContent, tableRowContent)
		}

		// 执行save操作，如果存在则更新，否则插入
		// TODO: 优化这里，为了保证对所有数据库做兼容，不应该使用save方法
		if _, err := db.Table(field.RelatedCollection).FieldsEx("deleted_at").Save(tableContent); err != nil {
			return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
		}

		// 自动处理需要删除的子表数据
		if _, err := db.Table(field.RelatedCollection).
			Where("id not in (?)", tableIds).
			Where("pcn", collectionName).
			Where("pid", id).
			Where("pfd", field.Name).
			Delete(); err != nil {
			return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
		}
	}

	// 执行 AfterUpdateHooks, AfterSaveHooks 勾子
	for _, hook := range global.GetAfterUpdateHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return err
		}
	}
	for _, hook := range global.GetAfterSaveHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return err
		}
	}

	return nil
}
