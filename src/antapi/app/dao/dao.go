package dao

import "github.com/gogf/gf/encoding/gjson"

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
