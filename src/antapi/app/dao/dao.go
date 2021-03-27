package dao

import (
	"antapi/app/global"
	"antapi/common/errcode"
	"fmt"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// TODO: goframe 框架目前不支持嵌入式的where查询，所以目前只能全部放到where里

// GetFuncArg .
type GetFuncArg struct {
	Where                 interface{}
	WhereArgs             interface{}
	Or                    interface{}
	OrArgs                interface{}
	Having                interface{}
	HavingArgs            interface{}
	Fields                []string
	SessionUsername       string
	IgnorePermissionCheck bool
	IgnoreFieldsCheck     bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
	RaiseNotFound         bool
}

// GetListFuncArg .
type GetListFuncArg struct {
	Where                 interface{}
	WhereArgs             interface{}
	Or                    interface{}
	OrArgs                interface{}
	Having                interface{}
	HavingArgs            interface{}
	Fields                []string
	Group                 string
	Order                 string
	PageNum               int
	PageSize              int
	SessionUsername       string
	IgnorePermissionCheck bool
	IgnoreFieldsCheck     bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
}

// ExistsAndCountFuncArg .
type ExistsAndCountFuncArg struct {
	Where      interface{}
	WhereArgs  interface{}
	Or         interface{}
	OrArgs     interface{}
	Having     interface{}
	HavingArgs interface{}
}

// InsertFuncArg .
type InsertFuncArg struct {
	SessionUsername       string
	IgnorePermissionCheck bool
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
}

// UpdateFuncArg .
type UpdateFuncArg struct {
	SessionUsername       string
	IgnorePermissionCheck bool
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
	RaiseNotFound         bool
}

// SaveFuncArg .
type SaveFuncArg struct {
	SessionUsername       string
	IgnorePermissionCheck bool
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
}

// DeleteFuncArg .
type DeleteFuncArg struct {
	SessionUsername       string
	IgnorePermissionCheck bool
	Where                 interface{}
	WhereArgs             interface{}
	Or                    interface{}
	OrArgs                interface{}
	Having                interface{}
	HavingArgs            interface{}
}

// dataToJson 任意类型数据转gjson对象
func dataToJson(data interface{}) *gjson.Json {
	if val, ok := data.(*gjson.Json); ok {
		return val
	}
	return gjson.New(data)
}

// CheckDuplicate 唯一性校验
// 此处的校验不是非常完美，唯一列必须保证在设计时属性强制设定为必填，数据入库时保证不能为null。否则此处校验会有会有
func CheckDuplicate(collectionName string, data *gjson.Json, excludeID ...string) error {
	schema := global.GetSchema(collectionName)
	fields := schema.GetUniqueFieldNames()
	if len(fields) == 0 {
		return nil
	}
	fields = append(fields, "id")
	m := g.DB().Table(collectionName).Fields(fields)
	fieldNameMap := map[string]string{}

	for _, field := range schema.GetUniqueFields() {
		fieldNameMap[field.Name] = field.Title
		m.Or(field.Name, data.Get(field.Name))
	}

	res, err := m.All()
	if err != nil {
		return gerror.WrapCode(errcode.ServerError, err, errcode.ServerErrorMsg)
	}
	if res.IsEmpty() {
		return nil
	}

	excludeIDArray := garray.NewStrArrayFrom(excludeID, true)
	resJson := gjson.New(res.Json())
	for i := 0; i < res.Len(); i++ {
		id := resJson.GetString(fmt.Sprintf("%d.id", i))
		if excludeIDArray.Contains(id) {
			continue
		}

		for fieldName, fieldTitle := range fieldNameMap {
			fieldVal := resJson.GetString(fmt.Sprintf("%d.%s", i, fieldName))
			if fieldVal == data.GetString(fieldName) {
				return gerror.NewCodef(errcode.DuplicateError, "已存在相同的%s: %s", fieldTitle, fieldVal)
			}
		}
	}

	return nil
}
