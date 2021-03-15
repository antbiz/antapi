package types

type DBType string

// Supported database types
const (
	POSTGRES DBType = "postgres"
	SQLITE   DBType = "sqlite3"
	MYSQL    DBType = "mysql"
	MSSQL    DBType = "mssql"
	MONGO    DBType = "mongo"
	// ORACLE   DBType = "oracle"
)

// Supported field types
const (
	Data      string = "Data"
	Color     string = "Color"
	Email     string = "Email"
	Tel       string = "Tel"
	URL       string = "Url"
	Password  string = "Password"
	Text      string = "Text"
	LongText  string = "LongText"
	RichText  string = "RichText"
	Markdown  string = "Markdown"
	Code      string = "Code"
	HTML      string = "HTML"
	Signature string = "Signature"
	File      string = "File"
	Enum      string = "Enum"
	JSON      string = "Json"
	UUID      string = "Uuid"
	Int       string = "Int"
	BigInt    string = "BigInt"
	Float     string = "Float"
	Money     string = "Money"
	Date      string = "Date"
	DateTime  string = "DateTime"
	Time      string = "Time"
	TimeStamp string = "TimeStamp"
	Year      string = "Year"
	Bool      string = "Bool"
	Array     string = "Array"
	Connect   string = "Connect"
)
