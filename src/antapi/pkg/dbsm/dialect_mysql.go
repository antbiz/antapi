package dbsm

import (
	"antapi/pkg/dbsm/types"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/database/gdb"
)

// MySQLDialect Implementation of Dialect for MySQL databases.
type MySQLDialect struct {
	DBName      string
	StoreEngine string
	Charset     string
}

func (dialect *MySQLDialect) DBType() types.DBType {
	return types.MYSQL
}

func (dialect *MySQLDialect) AutoIncrStr() string {
	return "AUTO_INCREMENT"
}

func (dialect *MySQLDialect) GetQuoter() Quoter {
	quoter := Quoter{
		Prefix: '`',
		Suffix: '`',
	}
	return quoter
}

func (dialect *MySQLDialect) SQLType(column *Column) string {
	var res string
	switch t := column.Type; t {
	case types.VARCHAR:
		res = t
		column.Size = 255
	case types.SMALLTEXT:
		res = types.VARCHAR
		column.Size = 1024
	case types.JSON:
		res = types.TEXT
	case types.BOOL:
		res = types.TINYINT
		column.Size = 1
	default:
		res = t
	}

	if column.Precision > 0 {
		res += "(" + strconv.Itoa(column.Size) + "," + strconv.Itoa(column.Precision) + ")"
	} else if column.Size > 0 {
		res += "(" + strconv.Itoa(column.Size) + ")"
	}
	return res
}

func (dialect *MySQLDialect) GetIndexes(tx *gdb.TX, tableName string) (map[string]*Index, error) {
	sql := fmt.Sprintf("SELECT `INDEX_NAME`, `NON_UNIQUE`, `COLUMN_NAME` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = '%s' AND `TABLE_NAME` = '%s'", dialect.DBName, tableName)
	rows, err := tx.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	indexes := make(map[string]*Index, 0)
	for rows.Next() {
		var indexIsUnique bool
		var indexName, colName, nonUnique string
		err = rows.Scan(&indexName, &nonUnique, &colName)
		if err != nil {
			return nil, err
		}

		if indexName == "PRIMARY" {
			continue
		}

		indexIsUnique = "YES" != nonUnique && nonUnique != "1"

		colName = strings.Trim(colName, "` ")
		if strings.HasPrefix(indexName, "IDX_"+tableName) || strings.HasPrefix(indexName, "UQE_"+tableName) {
			indexName = indexName[5+len(tableName):]
		}

		var index *Index
		var ok bool
		if index, ok = indexes[indexName]; !ok {
			index = new(Index)
			index.Name = indexName
			index.IsUnique = indexIsUnique
			indexes[indexName] = index
		}
		index.Cols = append(index.Cols, colName)
	}
	return indexes, nil
}

func (dialect *MySQLDialect) IndexCheckSQL(tableName, idxName string) string {
	return fmt.Sprintf("SELECT `INDEX_NAME` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = '%s' AND `TABLE_NAME` = '%s' AND `INDEX_NAME` = '%s'", dialect.DBName, tableName, idxName)
}

func (dialect *MySQLDialect) CreateIndexSQL(tableName string, index *Index) string {
	quoter := dialect.GetQuoter()
	var unique string
	var idxName string
	if index.IsUnique {
		unique = " UNIQUE"
	}
	idxName = index.XName(tableName)
	return fmt.Sprintf("CREATE%s INDEX %v ON %v (%v)", unique, quoter.Quote(idxName), quoter.Quote(tableName), quoter.Join(index.Cols, ","))
}

func (dialect *MySQLDialect) DropIndexSQL(tableName string, index *Index) string {
	quote := dialect.GetQuoter().Quote
	return fmt.Sprintf("DROP INDEX `%s` ON %s", quote(index.XName(tableName)), quote(tableName))
}

func (dialect *MySQLDialect) GetTables(tx *gdb.TX) ([]*Table, error) {
	sql := fmt.Sprintf("SELECT `TABLE_NAME`, `ENGINE`, `AUTO_INCREMENT`, `TABLE_COMMENT` FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA` = '%s' AND (`ENGINE` = 'MyISAM' OR `ENGINE` = 'InnoDB' OR `ENGINE` = 'TokuDB')", dialect.DBName)
	rows, err := tx.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := make([]*Table, 0)
	for rows.Next() {
		table := &Table{
			Columns:     make([]*Column, 0),
			Indexes:     make(map[string]*Index, 0),
			PrimaryKeys: make([]string, 0),
		}
		var name, engine string
		var autoIncr, comment *string
		err = rows.Scan(&name, &engine, &autoIncr, &comment)
		if err != nil {
			return nil, err
		}

		table.Name = name
		if comment != nil {
			table.Comment = *comment
		}
		table.StoreEngine = engine
		tables = append(tables, table)
	}

	for _, table := range tables {
		cols, err := dialect.GetColumns(tx, table.Name)
		if err != nil {
			return nil, err
		}
		table.UpdateWithCols(cols)
	}

	return tables, nil
}

func (dialect *MySQLDialect) IsTableExist(tx *gdb.TX, tableName string) (bool, error) {
	sql := fmt.Sprintf("SELECT `TABLE_NAME` FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA` = '%s' and `TABLE_NAME` = '%s'", dialect.DBName, tableName)
	v, err := tx.GetCount(sql)
	if err != nil {
		return false, err
	}
	return v > 0, nil
}

func (dialect *MySQLDialect) CreateTableSQL(table *Table) string {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (", table.Name)

	if len(table.Columns) > 0 {
		pkList := table.PrimaryKeys

		for _, col := range table.Columns {
			s, _ := ColumnString(dialect, col, col.IsPrimaryKey && len(pkList) == 1)
			sql += s
			sql = strings.TrimSpace(sql)
			if col.Comment != "" {
				sql += " COMMENT '" + col.Comment + "'"
			}
			sql += ", "
		}

		if len(pkList) > 1 {
			sql += fmt.Sprintf("PRIMARY KEY ( %v ), ", dialect.GetQuoter().Join(pkList, ","))
		}

		sql = sql[:len(sql)-2]
	}
	sql += ")"

	var storeEngine = table.StoreEngine
	if storeEngine == "" {
		storeEngine = dialect.StoreEngine
	}
	if table.StoreEngine != "" {
		sql += " ENGINE=" + table.StoreEngine
	}

	var charset = table.Charset
	if charset == "" {
		charset = dialect.Charset
	}
	if charset != "" {
		sql += " DEFAULT CHARSET " + charset
	}

	return sql
}

func (dialect *MySQLDialect) DropTableSQL(tableName string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS `%s`", tableName)
}

func (dialect *MySQLDialect) GetColumns(tx *gdb.TX, tableName string) ([]*Column, error) {
	alreadyQuoted := "(INSTR(VERSION(), 'maria') > 0 && " +
		"(SUBSTRING_INDEX(VERSION(), '.', 1) > 10 || " +
		"(SUBSTRING_INDEX(VERSION(), '.', 1) = 10 && " +
		"(SUBSTRING_INDEX(SUBSTRING(VERSION(), 4), '.', 1) > 2 || " +
		"(SUBSTRING_INDEX(SUBSTRING(VERSION(), 4), '.', 1) = 2 && " +
		"SUBSTRING_INDEX(SUBSTRING(VERSION(), 6), '-', 1) >= 7)))))"
	sql := "SELECT `COLUMN_NAME`, `IS_NULLABLE`, `COLUMN_DEFAULT`, `COLUMN_TYPE`," +
		" `COLUMN_KEY`, `EXTRA`, `COLUMN_COMMENT`, " +
		alreadyQuoted + " AS NEEDS_QUOTE " +
		fmt.Sprintf("FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = '%s' AND `TABLE_NAME` = '%s'", dialect.DBName, tableName) +
		" ORDER BY `COLUMNS`.ORDINAL_POSITION"

	rows, err := tx.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols := make([]*Column, 0)
	for rows.Next() {
		col := new(Column)

		var columnName, isNullable, colType, colKey, extra, comment string
		var alreadyQuoted bool
		var colDefault *string
		err = rows.Scan(&columnName, &isNullable, &colDefault, &colType, &colKey, &extra, &comment, &alreadyQuoted)
		if err != nil {
			return nil, err
		}
		col.Name = strings.Trim(columnName, "` ")
		col.Comment = comment
		if "YES" == isNullable {
			col.Nullable = true
		}

		if colDefault != nil && (!alreadyQuoted || *colDefault != "NULL") {
			col.Default = *colDefault
			col.DefaultIsEmpty = false
		} else {
			col.DefaultIsEmpty = true
		}

		cts := strings.Split(colType, "(")
		colName := cts[0]
		colType = strings.ToUpper(colName)
		var len1, len2 int
		if len(cts) == 2 {
			// Support Enum / Set ?
			idx := strings.Index(cts[1], ")")
			lens := strings.Split(cts[1][0:idx], ",")
			len1, err = strconv.Atoi(strings.TrimSpace(lens[0]))
			if err != nil {
				return nil, err
			}
			if len(lens) == 2 {
				len2, err = strconv.Atoi(lens[1])
				if err != nil {
					return nil, err
				}
			}
		}
		if colType == "FLOAT UNSIGNED" {
			colType = "FLOAT"
		}
		if colType == "DOUBLE UNSIGNED" {
			colType = "DOUBLE"
		}
		col.Size = len1
		col.Precision = len2
		col.Type = colType

		if colKey == "PRI" {
			col.IsPrimaryKey = true
		}
		if colKey == "UNI" {
			col.IsUnique = true
		}

		if extra == "auto_increment" {
			col.IsAutoIncrement = true
		}

		cols = append(cols, col)
	}
	return cols, nil
}

func (dialect *MySQLDialect) IsColumnExist(tx *gdb.TX, tableName string, colName string) (bool, error) {
	sql := fmt.Sprintf("SELECT `COLUMN_NAME` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = '%s' AND `TABLE_NAME` = '%s' AND `COLUMN_NAME` = '%s'", dialect.DBName, tableName, colName)
	v, err := tx.GetCount(sql)
	if err != nil {
		return false, err
	}
	return v > 0, nil
}

func (dialect *MySQLDialect) AddColumnSQL(tableName string, col *Column) string {
	s, _ := ColumnString(dialect, col, true)
	sql := fmt.Sprintf("ALTER TABLE `%s` ADD %v", tableName, s)
	if col.Comment != "" {
		sql += " COMMENT '" + col.Comment + "'"
	}
	return sql
}

func (dialect *MySQLDialect) ModifyColumnSQL(tableName string, col *Column) string {
	s, _ := ColumnString(dialect, col, false)
	return fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s", tableName, s)
}
