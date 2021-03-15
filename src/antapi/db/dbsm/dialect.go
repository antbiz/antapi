package dbsm

import (
	"database/sql"
	"strings"

	"antapi/db/types"
)

// Dialect represents a kind of database
type Dialect interface {
	DBType() types.DBType
	AutoIncrStr() string
	GetQuoter() Quoter
	SQLType(column *Column) string

	GetIndexes(tx *sql.Tx, tableName string) (map[string]*Index, error)
	CreateIndexSQL(tableName string, index *Index) string
	DropIndexSQL(tableName string, index *Index) string

	GetTables(tx *sql.Tx) ([]*Table, error)
	IsTableExist(tx *sql.Tx, tableName string) bool
	CreateTableSQL(table *Table) string
	DropTableSQL(tableName string) string

	GetColumns(tx *sql.Tx, tableName string) ([]*Column, error)
	IsColumnExist(tx *sql.Tx, tableName string, colName string) bool
	AddColumnSQL(tableName string, col *Column) string
	ModifyColumnSQL(tableName string, col *Column) string
}

// ColumnString generate column description string according dialect
func ColumnString(dialect Dialect, col *Column, includePrimaryKey bool) (string, error) {
	bd := strings.Builder{}

	if err := dialect.GetQuoter().QuoteTo(&bd, col.Name); err != nil {
		return "", err
	}

	if err := bd.WriteByte(' '); err != nil {
		return "", err
	}

	if _, err := bd.WriteString(dialect.SQLType(col)); err != nil {
		return "", err
	}

	if err := bd.WriteByte(' '); err != nil {
		return "", err
	}

	if includePrimaryKey && col.IsPrimaryKey {
		if _, err := bd.WriteString("PRIMARY KEY "); err != nil {
			return "", err
		}

		if col.IsAutoIncrement {
			if _, err := bd.WriteString(dialect.AutoIncrStr()); err != nil {
				return "", err
			}
			if err := bd.WriteByte(' '); err != nil {
				return "", err
			}
		}
	}

	if col.Default != "" {
		if _, err := bd.WriteString("DEFAULT "); err != nil {
			return "", err
		}
		if _, err := bd.WriteString(col.Default); err != nil {
			return "", err
		}
		if err := bd.WriteByte(' '); err != nil {
			return "", err
		}
	}

	if col.Nullable {
		if _, err := bd.WriteString("NULL "); err != nil {
			return "", err
		}
	} else {
		if _, err := bd.WriteString("NOT NULL "); err != nil {
			return "", err
		}
	}

	return bd.String(), nil
}
