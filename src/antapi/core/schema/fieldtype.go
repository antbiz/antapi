package schema

// FieldType .
type FieldType string

// Supported field types
const (
	String    FieldType = "String"
	Color     FieldType = "Color"
	Email     FieldType = "Email"
	Tel       FieldType = "Tel"
	URL       FieldType = "Url"
	Password  FieldType = "Password"
	Text      FieldType = "Text"
	LongText  FieldType = "LongText"
	RichText  FieldType = "RichText"
	Markdown  FieldType = "Markdown"
	Code      FieldType = "Code"
	HTML      FieldType = "HTML"
	Signature FieldType = "Signature"
	Media     FieldType = "Media"
	Enum      FieldType = "Enum"
	JSON      FieldType = "JSON"
	UUID      FieldType = "UUID"
	Int       FieldType = "Int"
	BigInt    FieldType = "BigInt"
	Float     FieldType = "Float"
	Money     FieldType = "Money"
	Date      FieldType = "Date"
	DateTime  FieldType = "DateTime"
	Time      FieldType = "Time"
	TimeStamp FieldType = "TimeStamp"
	Year      FieldType = "Year"
	Bool      FieldType = "Bool"
	Array     FieldType = "Array"
	BelongsTo FieldType = "BelongsTo"
	HasMany   FieldType = "HasMany"
)
