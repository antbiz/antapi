package dao

import (
	"antapi/app/global"
	"antapi/common/errcode"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// Delete : 删除指定数据
func Delete(collectionName string, arg *DeleteFuncArg) error {
	db := g.DB()
	schema := global.GetSchema(collectionName)

	// 查询需要删除的id
	var delIds []string
	records, err := db.Table(collectionName).Where(arg.Where, arg.WhereArgs).Or(arg.Or, arg.OrArgs).Having(arg.Having, arg.WhereArgs).All()
	if err != nil {
		return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}
	if records.IsEmpty() {
		if arg.RaiseNotFound {
			return gerror.NewCode(errcode.SourceNotFound, errcode.SourceNotFoundMsg)
		}
		return nil
	}

	recordsGJson := gjson.New(records.Json())
	recordsGJsonSlice := make([]*gjson.Json, 0, records.Len())
	for i := 0; i < records.Len(); i++ {
		recordGJson := recordsGJson.GetJson(fmt.Sprintf("%d", i))
		delIds = append(delIds, recordGJson.GetString("id"))
		recordsGJsonSlice = append(recordsGJsonSlice, recordGJson)
	}

	// 执行 BeforeDelete 勾子
	for _, recordGJson := range recordsGJsonSlice {
		for _, hook := range global.GetBeforeDeleteHooksByCollectionName(collectionName) {
			if err := hook(recordGJson); err != nil {
				return err
			}
		}
	}

	// 删除主体数据
	if arg.SessionUsername != "" {
		_, _ = db.Table(collectionName).Where("id", delIds).Update("updated_by", arg.SessionUsername)
	}
	if _, err := db.Table(collectionName).Where("id", delIds).Delete(); err != nil {
		return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}

	// 删除子表数据
	for _, field := range schema.GetTableFields() {
		if _, err := db.Table(field.RelatedCollection).Where("pid", delIds).Where("pfd", field.Name).Where("pct", collectionName).Delete(); err != nil {
			return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
		}
	}

	// 执行 AfterDelete 勾子
	for _, recordGJson := range recordsGJsonSlice {
		for _, hook := range global.GetAfterDeleteHooksByCollectionName(collectionName) {
			if err := hook(recordGJson); err != nil {
				return err
			}
		}
	}

	return nil
}
