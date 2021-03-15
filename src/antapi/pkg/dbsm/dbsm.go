package dbsm

import (
	"antapi/pkg/dbsm/types"
	"database/sql"
	"strings"
)

// Sync .
func Sync(tx *sql.Tx, dialect Dialect, tables []*Table) error {
	if err := sync(tx, dialect, tables); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// sync .
func sync(tx *sql.Tx, dialect Dialect, tables []*Table) error {
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
				oricolType := oriCol.Type
				if colType != oricolType {
					if colType == types.Text && strings.HasPrefix(oricolType, "VARCHAR") {
						if dialect.DBType() == types.MYSQL || dialect.DBType() == types.POSTGRES {
							needModifyCol = true
						}
					} else if strings.HasPrefix(oricolType, "VARCHAR") && strings.HasPrefix(colType, "VARCHAR") {
						if dialect.DBType() == types.MYSQL {
							if oriCol.Size < col.Size {
								needModifyCol = true
							}
						}
					}
				} else if colType == "VARCHAR" {
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
					if index.Equal(oriindex) {
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
