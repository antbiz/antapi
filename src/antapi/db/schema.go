package db

// Schema .
type Schema struct {
	Title       string
	Name        string
	ProjectName string
	Description string
	Fields      []*SchemaField
}

// SchemaField .
type SchemaField struct {
	Index             int
	Pid               string
	Type              string
	Title             string
	Name              string
	Description       string
	IsRequired        bool
	IsHidden          bool
	IsUnique          bool
	IsPrivate         bool
	IsEncrypted       bool
	CanSort           bool
	Default           string
	ConnectCollection string
	ConnectField      string
	ConnectMany       bool
	Min               int
	Max               int
	Validator         string
	Style             string
	EnumType          string
	EnumOptions       string
	IsMultiple        bool
}
