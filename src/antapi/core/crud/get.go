package crud

import (
	"antapi/hooks"
	"antapi/logic"
	"fmt"
	"strings"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// GetOne : 获取单个数据
func GetOne(collectionName string, where interface{}, args ...interface{}) (*gjson.Json, error) {
	db := g.DB()
	schema := logic.GetSchema(collectionName)

	record, err := db.Table(collectionName).Fields(schema.GetPublicFieldNames()).Where(where, args...).One()
	if err != nil {
		return nil, err
	}
	dataGJson := gjson.New(record.Json())

	// 查询子表数据
	for _, field := range schema.GetTableFields() {
		tableSchema := logic.GetSchema(field.RelatedCollection)

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

	// 填充LinkInfo, 查询指定范围内父子Link字段的关联的数据，先批量获取然后再按属性分配
	for linkCollectionName, linkPaths := range logic.DefaultSchemaLogic.GetLinkPathIncludeTableInner(schema) {
		var (
			linkIds                   []string
			tableRecordsLenByLinkPath map[string]int
		)
		for _, path := range linkPaths {
			_path := strings.Split(path, ".")
			// 是否为子表内的link字段
			isTableInner := len(_path) > 1
			if isTableInner {
				tableRecordsLen := len(dataGJson.GetArray(fmt.Sprintf("%s", _path[0])))
				if tableRecordsLen == 0 {
					continue
				}
				tableRecordsLenByLinkPath[path] = tableRecordsLen
				for j := 0; j < tableRecordsLen; j++ {
					linkIds = append(linkIds, dataGJson.GetString(fmt.Sprintf("%s.%d.%s", _path[0], j, _path[1])))
				}
			} else {
				linkIds = append(linkIds, dataGJson.GetString(path))
			}

		}
		linkRecords, err := db.Table(linkCollectionName).Fields(linkCollectionName).Where("id", linkIds).All()
		if err != nil {
			return nil, err
		}
		linkRecordsMap := linkRecords.MapKeyStr("id")
		for _, path := range linkPaths {
			_path := strings.Split(path, ".")
			isTableInner := len(_path) > 1
			if isTableInner {
				tableRecordsLen := tableRecordsLenByLinkPath[path]
				if tableRecordsLen == 0 {
					continue
				}
				for j := 0; j < tableRecordsLen; j++ {
					innderLinkRecordId := dataGJson.GetString(fmt.Sprintf("%s.%d.%s", _path[0], j, _path[1]))
					dataGJson.Set(fmt.Sprintf("%s.%d.%s_linkinfo", _path[0], j, _path[1]), linkRecordsMap[innderLinkRecordId])
				}
			} else {
				linkRecordId := dataGJson.GetString(path)
				dataGJson.Set(fmt.Sprintf("%s_linkinfo", path), linkRecordsMap[linkRecordId])
			}
		}
	}

	// 执行 AfterFindHooks 勾子
	for _, hook := range hooks.GetAfterFindHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return dataGJson, err
		}
	}

	return dataGJson, nil
}

// GetList : 获取列表数据
func GetList(collectionName string, pageNum, pageSize int, where interface{}, args ...interface{}) (*gjson.Json, error) {
	db := g.DB()
	schema := logic.GetSchema(collectionName)

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

	// 查询指定范围内子表数据，先批量获取然后再按属性分配
	for _, tableCollectionName := range schema.GetTableCollectionNames() {
		tableSchema := logic.GetSchema(tableCollectionName)

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

	// 填充LinkInfo, 查询指定范围内父子Link字段的关联的数据，先批量获取然后再按属性分配
	for linkCollectionName, linkPaths := range logic.DefaultSchemaLogic.GetLinkPathIncludeTableInner(schema) {
		var (
			linkIds                   []string
			tableRecordsLenByLinkPath map[string]int
		)
		for _, path := range linkPaths {
			_path := strings.Split(path, ".")
			// 是否为子表内的link字段
			isTableInner := len(_path) > 1
			for i := 0; i < recordsLen; i++ {
				if isTableInner {
					tableRecordsLen := len(listDataGJson.GetArray(fmt.Sprintf("%d.%s", i, _path[0])))
					if tableRecordsLen == 0 {
						continue
					}
					tableRecordsLenByLinkPath[fmt.Sprintf("%d.%s", i, path)] = tableRecordsLen
					for j := 0; j < tableRecordsLen; j++ {
						linkIds = append(linkIds, listDataGJson.GetString(fmt.Sprintf("%d.%s.%d.%s", i, _path[0], j, _path[1])))
					}
				} else {
					linkIds = append(linkIds, listDataGJson.GetString(fmt.Sprintf("%d.%s", i, path)))
				}
			}

		}
		linkRecords, err := db.Table(linkCollectionName).Fields(linkCollectionName).Where("id", linkIds).All()
		if err != nil {
			return nil, err
		}
		linkRecordsMap := linkRecords.MapKeyStr("id")
		for _, path := range linkPaths {
			_path := strings.Split(path, ".")
			isTableInner := len(_path) > 1
			for i := 0; i < recordsLen; i++ {
				if isTableInner {
					tableRecordsLen := tableRecordsLenByLinkPath[fmt.Sprintf("%d.%s", i, path)]
					if tableRecordsLen == 0 {
						continue
					}
					for j := 0; j < tableRecordsLen; j++ {
						innderLinkRecordId := listDataGJson.GetString(fmt.Sprintf("%d.%s.%d.%s", i, _path[0], j, _path[1]))
						listDataGJson.Set(fmt.Sprintf("%d.%s.%d.%s_linkinfo", i, _path[0], j, _path[1]), linkRecordsMap[innderLinkRecordId])
					}
				} else {
					linkRecordId := listDataGJson.GetString(fmt.Sprintf("%d.%s", i, path))
					listDataGJson.Set(fmt.Sprintf("%d.%s_linkinfo", i, path), linkRecordsMap[linkRecordId])
				}
			}
		}
	}

	// 执行 AfterFindHooks 勾子
	for i := 0; i < recordsLen; i++ {
		dataGJson := listDataGJson.GetJson(fmt.Sprintf("%d", i))
		for _, hook := range hooks.GetAfterFindHooksByCollectionName(collectionName) {
			if err := hook(dataGJson); err != nil {
				return dataGJson, err
			}
		}
	}

	return listDataGJson, nil
}
