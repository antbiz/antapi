package dao

import (
	"antapi/app/global"
	"antapi/app/model"
	"antapi/common/errcode"
	"fmt"
	"strings"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// getLinkPathIncludeTableInner : 获取所有link字段的路径，包括子表
func getLinkPathIncludeTableInner(schema *model.Schema) (paths map[string][]string) {
	for _, linkField := range schema.GetLinkFields() {
		paths[linkField.RelatedCollection] = append(paths[linkField.RelatedCollection], linkField.Name)
	}
	for _, tableField := range schema.GetTableFields() {
		tableSchema := global.GetSchema(tableField.RelatedCollection)
		for _, tableLinkField := range tableSchema.GetLinkFields() {
			paths[tableLinkField.RelatedCollection] = append(paths[tableLinkField.RelatedCollection], fmt.Sprintf("%s.%s", tableField.Name, tableLinkField.Name))
		}
	}
	return
}

// Get : 获取单个数据
func Get(collectionName string, arg *GetFuncArg) (*gjson.Json, error) {
	db := g.DB()
	schema := global.GetSchema(collectionName)

	m := db.Table(collectionName).Where(arg.Where, arg.WhereArgs).Or(arg.Or, arg.OrArgs).Having(arg.Having, arg.HavingArgs)
	if len(arg.Fields) > 0 {
		if arg.IgnoreFieldsCheck {
			m.Fields(arg.Fields)
		} else {
			passFieldNames := make([]string, 0)
			allFieldNames := garray.NewStrArrayFrom(schema.GetFieldNames(true, true))
			hiddenFieldNames := garray.NewStrArrayFrom(schema.GetHiddenFieldNames())
			privateFieldNames := garray.NewStrArrayFrom(schema.GetPrivateFieldNames())
			for _, fieldName := range arg.Fields {
				if fieldName == "" {
					continue
				}
				if !allFieldNames.Contains(fieldName) {
					continue
				}
				if !arg.IncludeHiddenField && hiddenFieldNames.Contains(fieldName) {
					continue
				}
				if !arg.IncludePrivateField && privateFieldNames.Contains(fieldName) {
					continue
				}
				passFieldNames = append(passFieldNames, fieldName)
			}
			m.Fields(passFieldNames)
		}
	} else {
		m.Fields(schema.GetFieldNames(arg.IncludeHiddenField, arg.IncludePrivateField))
	}
	record, err := m.One()
	if err != nil {
		return nil, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}
	if record.IsEmpty() {
		if arg.RaiseNotFound {
			return nil, gerror.NewCode(errcode.SourceNotFound, errcode.SourceNotFoundMsg)
		}
		return nil, nil
	}
	dataGJson := gjson.New(record.Json())

	// 查询子表数据
	for _, field := range schema.GetTableFields() {
		tableSchema := global.GetSchema(field.RelatedCollection)

		tableRecords, err := db.Table(field.RelatedCollection).
			Fields(tableSchema.GetFieldNames(arg.IncludeHiddenField, arg.IncludePrivateField)).
			Order("idx asc").
			Where("pid", dataGJson.Get("id")).
			Where("pcn", collectionName).
			Where("pfd", field.Name).
			All()
		if err != nil {
			return nil, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
		}
		dataGJson.Set(field.RelatedCollection, tableRecords.Json())
	}

	// 填充LinkInfo, 查询指定范围内父子Link字段的关联的数据，先批量获取然后再按属性分配
	for linkCollectionName, linkPaths := range getLinkPathIncludeTableInner(schema) {
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

		linkCollectionFieldNames := global.GetSchema(linkCollectionName).GetFieldNames(false, false)
		linkRecords, err := db.Table(linkCollectionName).Fields(linkCollectionFieldNames).Where("id", linkIds).All()
		if err != nil {
			return nil, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
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
	for _, hook := range global.GetAfterFindHooksByCollectionName(collectionName) {
		if err := hook(dataGJson); err != nil {
			return dataGJson, err
		}
	}

	return dataGJson, nil
}

// GetList : 获取列表数据
func GetList(collectionName string, arg *GetListFuncArg) (list *gjson.Json, total int, err error) {
	db := g.DB()
	schema := global.GetSchema(collectionName)

	// 查询指定范围内主体数据list
	m := db.Table(collectionName).Where(arg.Where, arg.WhereArgs).Or(arg.Or, arg.OrArgs).Having(arg.Having, arg.HavingArgs)

	if total, err = m.Count(); err != nil {
		return nil, 0, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}

	if len(arg.Fields) > 0 {
		if arg.IgnoreFieldsCheck {
			m.Fields(arg.Fields)
		} else {
			passFieldNames := make([]string, 0)
			allFieldNames := garray.NewStrArrayFrom(schema.GetFieldNames(true, true))
			hiddenFieldNames := garray.NewStrArrayFrom(schema.GetHiddenFieldNames())
			privateFieldNames := garray.NewStrArrayFrom(schema.GetPrivateFieldNames())
			for _, fieldName := range arg.Fields {
				if fieldName == "" {
					continue
				}
				if !allFieldNames.Contains(fieldName) {
					continue
				}
				if !arg.IncludeHiddenField && hiddenFieldNames.Contains(fieldName) {
					continue
				}
				if !arg.IncludePrivateField && privateFieldNames.Contains(fieldName) {
					continue
				}
				passFieldNames = append(passFieldNames, fieldName)
			}
			m.Fields(passFieldNames)
		}
	} else {
		m.Fields(schema.GetFieldNames(arg.IncludeHiddenField, arg.IncludePrivateField))
	}
	if arg.PageNum > 0 && arg.PageSize > 0 {
		m.Limit((arg.PageNum-1)*arg.PageSize, arg.PageSize)
	}
	records, err := m.All()
	if err != nil {
		return nil, 0, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}
	if records.IsEmpty() {
		return nil, 0, nil
	}

	recordsLen := records.Len()
	list = gjson.New(records.Json())

	ids := make([]string, 0, recordsLen)
	for i := 0; i < recordsLen; i++ {
		ids = append(ids, list.GetString(fmt.Sprintf("%d.id", i)))
	}

	// 查询指定范围内子表数据，先批量获取然后再按属性分配
	for _, tableCollectionName := range schema.GetTableCollectionNames() {
		tableSchema := global.GetSchema(tableCollectionName)

		tableRecords, err := db.Table(tableCollectionName).
			Fields(tableSchema.GetFieldNames(arg.IncludeHiddenField, arg.IncludePrivateField)).
			Order("idx asc").
			Where("pcn", collectionName).
			Where("pid", ids).
			All()
		if err != nil {
			return nil, 0, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
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
			pid := list.GetString(fmt.Sprintf("%d.id", i))
			for _, pfd := range pfdSlice {
				if err := list.Set(fmt.Sprintf("%d.%s", i, pfd), tableGroupRecords[fmt.Sprintf("%s@%s", pid, pfd)]); err != nil {
					return nil, 0, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
				}
			}
		}
	}

	// 填充LinkInfo, 查询指定范围内父子Link字段的关联的数据，先批量获取然后再按属性分配
	for linkCollectionName, linkPaths := range getLinkPathIncludeTableInner(schema) {
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
					tableRecordsLen := len(list.GetArray(fmt.Sprintf("%d.%s", i, _path[0])))
					if tableRecordsLen == 0 {
						continue
					}
					tableRecordsLenByLinkPath[fmt.Sprintf("%d.%s", i, path)] = tableRecordsLen
					for j := 0; j < tableRecordsLen; j++ {
						linkIds = append(linkIds, list.GetString(fmt.Sprintf("%d.%s.%d.%s", i, _path[0], j, _path[1])))
					}
				} else {
					linkIds = append(linkIds, list.GetString(fmt.Sprintf("%d.%s", i, path)))
				}
			}

		}

		linkCollectionFieldNames := global.GetSchema(linkCollectionName).GetFieldNames(false, false)
		linkRecords, err := db.Table(linkCollectionName).Fields(linkCollectionFieldNames).Where("id", linkIds).All()
		if err != nil {
			return nil, 0, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
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
						innderLinkRecordId := list.GetString(fmt.Sprintf("%d.%s.%d.%s", i, _path[0], j, _path[1]))
						list.Set(fmt.Sprintf("%d.%s.%d.%s_linkinfo", i, _path[0], j, _path[1]), linkRecordsMap[innderLinkRecordId])
					}
				} else {
					linkRecordId := list.GetString(fmt.Sprintf("%d.%s", i, path))
					list.Set(fmt.Sprintf("%d.%s_linkinfo", i, path), linkRecordsMap[linkRecordId])
				}
			}
		}
	}

	// 执行 AfterFindHooks 勾子
	for i := 0; i < recordsLen; i++ {
		dataGJson := list.GetJson(fmt.Sprintf("%d", i))
		for _, hook := range global.GetAfterFindHooksByCollectionName(collectionName) {
			if err = hook(dataGJson); err != nil {
				return
			}
		}
	}

	return
}

// Count : 查询统计
func Count(collectionName string, arg *ExistsAndCountFuncArg) (int, error) {
	total, err := g.Table(collectionName).Where(arg.Where, arg.WhereArgs).Or(arg.Or, arg.OrArgs).Having(arg.Having, arg.HavingArgs).Count()
	if err != nil {
		return 0, gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}
	return total, nil
}

// Exists : 是否存在
func Exists(collectionName string, arg *ExistsAndCountFuncArg) (bool, error) {
	total, err := Count(collectionName, arg)
	if err != nil {
		return false, err
	}
	return total > 0, nil
}
