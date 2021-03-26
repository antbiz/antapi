package dbsm

import (
	"antapi/pkg/dbsm/types"
	"strings"

	"github.com/gogf/gf/database/gdb"
)

// Dialect represents a kind of database
type Dialect interface {
	DBType() types.DBType
	AutoIncrStr() string
	GetQuoter() Quoter
	SQLType(column *Column) string

	GetIndexes(tx *gdb.TX, tableName string) (map[string]*Index, error)
	CreateIndexSQL(tableName string, index *Index) string
	DropIndexSQL(tableName string, index *Index) string

	GetTables(tx *gdb.TX) ([]*Table, error)
	IsTableExist(tx *gdb.TX, tableName string) (bool, error)
	CreateTableSQL(table *Table) string
	DropTableSQL(tableName string) string

	GetColumns(tx *gdb.TX, tableName string) ([]*Column, error)
	IsColumnExist(tx *gdb.TX, tableName string, colName string) (bool, error)
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
