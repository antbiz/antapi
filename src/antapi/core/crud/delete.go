package crud

import (
	"antapi/model"

	"github.com/gogf/gf/frame/g"
)

// Delete : 删除指定数据
func Delete(collectionName string, where interface{}, args ...interface{}) error {
	db := g.DB()
	schema, err := model.GetSchema(collectionName)

	// 查询需要删除的id
	var delIds []string
	values, err := db.Table(collectionName).Where(where, args...).Array("id")
	if err != nil {
		return err
	}
	for _, val := range values {
		delIds = append(delIds, val.String())
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

	return nil
}
