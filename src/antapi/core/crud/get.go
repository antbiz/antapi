package crud

import (
	"antapi/model"
	"fmt"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// GetOne : 获取单个数据
func GetOne(collectionName string, where interface{}, args ...interface{}) (*gjson.Json, error) {
	db := g.DB()
	schema, err := model.GetSchema(collectionName)
	if err != nil {
		return nil, err
	}

	record, err := db.Table(collectionName).Fields(schema.GetPublicFieldNames()).Where(where, args...).One()
	if err != nil {
		return nil, err
	}
	dataGJson := gjson.New(record.Json())

	for _, field := range schema.GetLinkFields() {
		linkSchema, err := model.GetSchema(field.RelatedCollection)
		if err != nil {
			return nil, err
		}
		linkRecord, err := db.Table(field.RelatedCollection).Fields(linkSchema.GetPublicFieldNames()).Where("id", dataGJson.GetString(field.Name)).One()
		if err != nil {
			return nil, err
		}
		dataGJson.Set(field.RelatedCollection, linkRecord.Json())
	}

	for _, field := range schema.GetTableFields() {
		tableSchema, err := model.GetSchema(field.RelatedCollection)
		if err != nil {
			return nil, err
		}
		tableRecords, err := db.Table(field.RelatedCollection).
			Fields(tableSchema.GetTableFieldNames()).
			Order("idx asc").
			Where("pid", dataGJson.Get("id")).
			Where("pcn", collectionName).
			Where("pfd", field.Name).
			All()
		if err != nil {
			return nil, err
		}
		dataGJson.Set(field.RelatedCollection, tableRecords.Json())
	}

	// 执行 AfterFindHooks 勾子
	for _, hook := range model.AfterFindHooks[collectionName] {
		if err := hook(dataGJson); err != nil {
			return dataGJson, err
		}
	}

	return dataGJson, nil
}

// GetList : 获取列表数据
func GetList(collectionName string, pageNum, pageSize int, where interface{}, args ...interface{}) (*gjson.Json, error) {
	db := g.DB()
	schema, err := model.GetSchema(collectionName)
	if err != nil {
		return nil, err
	}

	// 查询指定范围内主体数据list
	orm := db.Table(collectionName).Fields(schema.GetPublicFieldNames()).Where(where, args...)
	if pageNum > 0 && pageSize > 0 {
		orm = orm.Limit((pageNum-1)*pageSize, pageSize)
	}
	records, err := orm.All()
	if err != nil {
		return nil, err
	}
	recordsLen := records.Len()
	listDataGJson := gjson.New(records.Json())

	ids := make([]string, 0, recordsLen)
	for i := 0; i < recordsLen; i++ {
		ids = append(ids, listDataGJson.GetString(fmt.Sprintf("%d.id", i)))
	}

	// 查询指定范围内Link字段的关联的数据，先批量获取然后再按属性分配
	var linkRecordsMap map[string]map[string]interface{}
	for _, linkCollectionName := range schema.GetLinkCollectionNames() {
		linkSchema, err := model.GetSchema(linkCollectionName)
		if err != nil {
			return nil, err
		}
		linkRecords, err := db.Table(linkCollectionName).Fields(linkSchema.GetPublicFieldNames()).Where("id", ids).All()
		if err != nil {
			return nil, err
		}
		for k, v := range linkRecords.MapKeyStr("id") {
			linkRecordsMap[fmt.Sprintf("%s@%s", linkCollectionName, k)] = v
		}
	}
	for _, field := range schema.GetLinkFields() {
		for i := 0; i < recordsLen; i++ {
			linkId := listDataGJson.GetString(fmt.Sprintf("%d.%s", i, field.Name))
			if err := listDataGJson.Set(fmt.Sprintf("%d.%s", i, field.RelatedCollection), linkRecordsMap[fmt.Sprintf("%s@%s", field.RelatedCollection, linkId)]); err != nil {
				return nil, err
			}
		}
	}

	// 查询指定范围内Table字段的关联的数据，先批量获取然后再按属性分配
	for _, tableCollectionName := range schema.GetTableCollectionNames() {
		tableSchema, err := model.GetSchema(tableCollectionName)
		if err != nil {
			return nil, err
		}
		tableRecords, err := db.Table(tableCollectionName).
			Fields(tableSchema.GetPublicFieldNames()).
			Order("idx asc").
			Where("pcn", collectionName).
			Where("pid", ids).
			All()
		if err != nil {
			return nil, err
		}

		var (
			tableGroupRecords map[string][]string
			pfdArr            = garray.NewStrArray()
		)
		for _, tableRecord := range tableRecords {
			tableRecordMap := tableRecord.GMap()
			pid := tableRecordMap.GetVar("pid").String()
			pfd := tableRecordMap.GetVar("pfd").String()
			if !pfdArr.Contains(pfd) {
				pfdArr.Append(pfd)
			}
			tableGroupRecords[fmt.Sprintf("%s@%s", pid, pfd)] = append(tableGroupRecords[pid], tableRecord.Json())
		}

		pfdSlice := pfdArr.Slice()
		for i := 0; i < recordsLen; i++ {
			pid := listDataGJson.GetString(fmt.Sprintf("%d.id", i))
			for _, pfd := range pfdSlice {
				if err := listDataGJson.Set(fmt.Sprintf("%d.%s", i, pfd), tableGroupRecords[fmt.Sprintf("%s@%s", pid, pfd)]); err != nil {
					return nil, err
				}
			}
		}
	}

	// 执行 AfterFindHooks 勾子
	for i := 0; i < recordsLen; i++ {
		dataGJson := listDataGJson.GetJson(fmt.Sprintf("%d", i))
		for _, hook := range model.AfterFindHooks[collectionName] {
			if err := hook(dataGJson); err != nil {
				return dataGJson, err
			}
		}
	}

	return listDataGJson, nil
}
