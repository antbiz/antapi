package db

// Schema .
type Schema struct {
	Title       string         `json bson db:"title"`
	Name        string         `json bson db:"name"`
	ProjectName string         `json bson db:"project_name"`
	Description string         `json bson db:"description"`
	Fields      []*SchemaField `json:"fields"`
}

// SchemaField .
type SchemaField struct {
	Index             int    `json bson db:"idx,omitempty"`
	Pid               string `json bson db:"pid,omitempty"`
	Type              string `json bson db:"type"`
	Title             string `json bson db:"title"`
	Name              string `json bson db:"name"`
	Description       string `json bson db:"description"`
	IsRequired        bool   `json bson db:"is_required"`
	IsHidden          bool   `json bson db:"is_hidden"`
	IsUnique          bool   `json bson db:"is_unique"`
	IsPrivate         bool   `json bson db:"is_private"`
	IsEncrypted       bool   `json bson db:"is_encrypted"`
	CanSort           bool   `json bson db:"can_sort"`
	Default           string `json bson db:"default"`
	ConnectCollection string `json bson db:"connect_collection"`
	ConnectField      string `json bson db:"connect_field"`
	ConnectMany       bool   `json bson db:"connect_many"`
	Min               int    `json bson db:"min"`
	Max               int    `json bson db:"max"`
	Validator         string `json bson db:"validator"`
	Style             string `json bson db:"style"`
	EnumType          string `json bson db:"enum_type"`
	EnumOptions       string `json bson db:"enum_options"`
	IsMultiple        bool   `json bson db:"is_multiple"`
}
