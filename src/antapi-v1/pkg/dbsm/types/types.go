package types

type DBType string

// Supported database types
const (
	POSTGRES DBType = "pgsql"
	SQLITE   DBType = "sqlite"
	MYSQL    DBType = "mysql"
	MSSQL    DBType = "mssql"
	ORACLE   DBType = "oracle"
	// MONGO    DBType = "mongo"
)

// Supported field types
const (
	VARCHAR   string = "VARCHAR"
	SMALLTEXT string = "SMALLTEXT"
	TEXT      string = "TEXT"
	LONGTEXT  string = "LONGTEXT"
	JSON      string = "JSON"
	BLOB      string = "BLOB"
	BOOL      string = "BOOL"
	TINYINT   string = "TINYINT"
	INT       string = "INT"
	BIGINT    string = "BIGINT"
	FLOAT     string = "FLOAT"
	DATE      string = "DATE"
	DATETIME  string = "DATETIME"
	TIME      string = "TIME"
	TIMESTAMP string = "TIMESTAMP"
	YEAR      string = "YEAR"
)
