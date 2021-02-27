package gdbsm

import (
	"antapi/pkg/gdbsm/types"
	"fmt"
	"strings"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

// Sync .
func Sync(dialect Dialect, tables []*Table) error {
	if tx, err := g.DB().Begin(); err != nil {
		return err
	} else {
		if err = sync(tx, dialect, tables); err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
	}
	return nil
}

// SyncFromJsonSchema .
func SyncFromJsonSchema(dialect Dialect, jsonSchema string) error {
	if j, err := gjson.DecodeToJson(jsonSchema); err != nil {
		return err
	} else {
		tables := make([]*Table, 0)
		columns := make([]*Column, 0)
		for _, colName := range gutil.Keys(j.GetMap("schema.properties")) {
			columns = append(columns, &Column{
				Name:            colName,
				Type:            j.GetString(fmt.Sprintf("schema.properties.%s.db:type", colName)),
				Size:            j.GetInt(fmt.Sprintf("schema.properties.%s.db:size", colName)),
				Precision:       j.GetInt(fmt.Sprintf("schema.properties.%s.db:precision", colName)),
				Default:         j.GetString(fmt.Sprintf("schema.properties.%s.db:default", colName)),
				IndexName:       j.GetString(fmt.Sprintf("schema.properties.%s.db:indexName", colName)),
				Nullable:        j.GetBool(fmt.Sprintf("schema.properties.%s.db:nullable", colName)),
				IsUnique:        j.GetBool(fmt.Sprintf("schema.properties.%s.db:isUnique", colName)),
				IsAutoIncrement: j.GetBool(fmt.Sprintf("schema.properties.%s.db:isAutoIncrement", colName)),
				IsPrimaryKey:    j.GetBool(fmt.Sprintf("schema.properties.%s.db:isPrimaryKey", colName)),
				Comment:         j.GetString(fmt.Sprintf("schema.properties.%s.db:comment", colName)),
			})
		}
		table := NewTable(j.GetString("schema.db:tableName"), columns)
		tables = append(tables, table)
		return Sync(dialect, tables)
	}
}

// sync .
func sync(tx *gdb.TX, dialect Dialect, tables []*Table) error {
	oriTables, err := dialect.GetTables(tx)
	if err != nil {
		return err
	}

	for _, table := range tables {
		var oriTable *Table
		for _, oritable := range oriTables {
			if oritable.Name == table.Name {
				oriTable = oritable
				break
			}
		}

		// this is a new table
		if oriTable == nil {
			if _, err = tx.Exec(dialect.CreateTableSQL(table)); err != nil {
				return err
			}

			for _, index := range table.Indexes {
				if _, err = tx.Exec(dialect.CreateIndexSQL(table.Name, index)); err != nil {
					return err
				}
			}
		} else {
			// this will modify an old table
			for _, col := range table.Columns {
				var oriCol *Column
				for _, oricol := range oriTable.Columns {
					if strings.EqualFold(col.Name, oricol.Name) {
						oriCol = col
						break
					}
				}

				// column is not exists on table
				if oriCol == nil {
					if _, err = tx.Exec(dialect.AddColumnSQL(table.Name, col)); err != nil {
						return err
					}
					continue
				}

				needModifyCol := false
				colType := dialect.SQLType(col)
				oricolType := dialect.SQLType(oriCol)
				if colType != oricolType {
					if colType == types.Text && strings.HasPrefix(oricolType, types.Varchar) {
						if dialect.DBType() == types.MYSQL || dialect.DBType() == types.POSTGRES {
							needModifyCol = true
						}
					} else if strings.HasPrefix(oricolType, types.Varchar) && strings.HasPrefix(colType, types.Varchar) {
						if dialect.DBType() == types.MYSQL {
							if oriCol.Size < col.Size {
								needModifyCol = true
							}
						}
					}
				} else if colType == types.Varchar {
					if dialect.DBType() == types.MYSQL {
						if oriCol.Size < col.Size {
							needModifyCol = true
						}
					}
				}
				if needModifyCol {
					if _, err = tx.Exec(dialect.ModifyColumnSQL(table.Name, col)); err != nil {
						return err
					}
				}
			} // end table columns for-loop

			var (
				foundIndexNames = make(map[string]bool)
				addedNames      = make(map[string]*Index)
			)

			for name, index := range table.Indexes {
				var oriIndex *Index
				for oriname, oriindex := range oriTable.Indexes {
					if index.Equal(oriIndex) {
						oriIndex = oriindex
						foundIndexNames[oriname] = true
						break
					}
				}

				if oriIndex != nil {
					if oriIndex.IsUnique != index.IsUnique {
						if _, err = tx.Exec(dialect.DropIndexSQL(table.Name, oriIndex)); err != nil {
							return err
						}
						oriIndex = nil
					}
				}

				if oriIndex == nil {
					addedNames[name] = index
				}
			}

			for _, index := range addedNames {
				if _, err = tx.Exec(dialect.CreateIndexSQL(table.Name, index)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
