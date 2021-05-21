package dao

import "github.com/antbiz/antapi/internal/common/types"

// ListOptions .
type ListOptions struct {
	Filter              interface{}
	Fields              []string
	Sort                []string
	Limit               int64
	Offset              int64
	IgnoreFieldsCheck   bool
	IncludeHiddenField  bool
	IncludePrivateField bool
	CtxUser             types.ContextUser
}

// GetOptions .
type GetOptions struct {
	Filter              interface{}
	Fields              []string
	IgnoreFieldsCheck   bool
	IncludeHiddenField  bool
	IncludePrivateField bool
	RaiseNotFound       bool
	CtxUser             types.ContextUser
}

// InsertOptions .
type InsertOptions struct {
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
	CtxUser               types.ContextUser
}

// UpsertOptions .
type UpsertOptions struct {
	Filter                interface{}
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
	CtxUser               types.ContextUser
}

// UpdateOptions .
type UpdateOptions struct {
	Filter                interface{}
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
	CtxUser               types.ContextUser
}

// DeleteOptions .
type DeleteOptions struct {
	Filter        interface{}
	RaiseNotFound bool
	CtxUser       types.ContextUser
}
