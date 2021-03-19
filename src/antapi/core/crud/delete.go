package crud

import (
	"antapi/model"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// Delete : 删除指定数据
func Delete(collectionName string, where interface{}, args ...interface{}) error {
	db := g.DB()
	schema, err := model.GetSchema(collectionName)

	// 查询需要删除的id
	var delIds []string
	records, err := db.Table(collectionName).Where(where, args...).All()
	if err != nil {
		return err
	}
	if records.Len() == 0 {
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
		for _, hook := range model.BeforeDeleteHooks[collectionName] {
			if err := hook(recordGJson); err != nil {
				return err
			}
		}
	}

	// 删除主体数据
	if _, err := db.Table(collectionName).Where("id", delIds).Delete(); err != nil {
		return err
	}

	// 删除子表数据
	for _, field := range schema.GetTableFields() {
		if _, err := db.Table(field.RelatedCollection).Where("pid", delIds).Where("pfd", field.Name).Where("pct", collectionName).Delete(); err != nil {
			return err
		}
	}

	// 执行 AfterDelete 勾子
	for _, recordGJson := range recordsGJsonSlice {
		for _, hook := range model.AfterDeleteHooks[collectionName] {
			if err := hook(recordGJson); err != nil {
				return err
			}
		}
	}

	return nil
}
