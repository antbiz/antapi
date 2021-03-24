package dao

// TODO: goframe 框架目前不支持嵌入式的where查询，所以目前只能全部放到where里

// GetFuncArg .
type GetFuncArg struct {
	Where               interface{}
	Or                  interface{}
	Having              interface{}
	Fields              []string
	IgnoreFieldsCheck   bool
	IncludeHiddenField  bool
	IncludePrivateField bool
}

// GetListFuncArg .
type GetListFuncArg struct {
	Where               interface{}
	Or                  interface{}
	Having              interface{}
	Fields              []string
	IgnoreFieldsCheck   bool
	Group               string
	Order               string
	PageNum             int
	PageSize            int
	IncludeHiddenField  bool
	IncludePrivateField bool
}

// ExistsAndCountFuncArg .
type ExistsAndCountFuncArg struct {
	Where  interface{}
	Or     interface{}
	Having interface{}
}

// InsertFuncArg .
type InsertFuncArg struct {
	IgnoreFieldsCheck   bool
	IgnoreFieldRequired bool
	IncludeHiddenField  bool
	IncludePrivateField bool
}

// InsertListFuncArg .
type InsertListFuncArg struct {
	IgnoreFieldsCheck   bool
	IgnoreFieldRequired bool
	IncludeHiddenField  bool
	IncludePrivateField bool
}

// UpdateFuncArg .
type UpdateFuncArg struct {
	Where               interface{}
	Or                  interface{}
	IgnoreFieldsCheck   bool
	IgnoreFieldRequired bool
	IncludeHiddenField  bool
	IncludePrivateField bool
}

// UpdateListFuncArg .
type UpdateListFuncArg struct {
	Where               interface{}
	Or                  interface{}
	IgnoreFieldsCheck   bool
	IgnoreFieldRequired bool
	IncludeHiddenField  bool
	IncludePrivateField bool
}

// DeleteFuncArg .
type DeleteFuncArg struct {
	Where               interface{}
	Or                  interface{}
	Having              interface{}
	IncludeHiddenField  bool
	IncludePrivateField bool
}
