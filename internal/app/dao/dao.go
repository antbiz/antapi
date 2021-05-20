package dao

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
}

// GetOptions .
type GetOptions struct {
	Filter              interface{}
	Fields              []string
	IgnoreFieldsCheck   bool
	IncludeHiddenField  bool
	IncludePrivateField bool
	RaiseNotFound       bool
}

// InsertOptions .
type InsertOptions struct {
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
}

// UpsertOptions .
type UpsertOptions struct {
	Filter                interface{}
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
}

// UpdateOptions .
type UpdateOptions struct {
	Filter                interface{}
	IgnoreFieldValueCheck bool
	IncludeHiddenField    bool
	IncludePrivateField   bool
}

// DeleteOptions .
type DeleteOptions struct {
	Filter        interface{}
	RaiseNotFound bool
}
