package gdbsm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gogf/gf/frame/g"
)

func TestSyncFromJsonSchema(t *testing.T) {
	g.Config().SetPath("C:\\BeanWei\\Workspace\\antbiz\\antapi\\config")

	var (
		fieldSchemaGroup = []map[string]string{
			{
				`"id"`: `{"db:type": "BigInt", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "", "db:nullable": 0, "db:isUnique": 0, "db:isAutoIncrement": 1, "db:isPrimaryKey": 1, "db:comment": "id"}`,
			},
			{
				`"id"`:   `{"db:type": "BigInt", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "", "db:nullable": 0, "db:isUnique": 0, "db:isAutoIncrement": 1, "db:isPrimaryKey": 1, "db:comment": "id"}`,
				`"name"`: `{"db:type": "VARCHAR", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "name", "db:nullable": 0, "db:isUnique": 1, "db:isAutoIncrement": 0, "db:isPrimaryKey": 0, "db:comment": "姓名"}`,
			},
			{
				`"id"`:   `{"db:type": "BigInt", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "", "db:nullable": 0, "db:isUnique": 0, "db:isAutoIncrement": 1, "db:isPrimaryKey": 1, "db:comment": "id"}`,
				`"name"`: `{"db:type": "VARCHAR", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "name", "db:nullable": 0, "db:isUnique": 1, "db:isAutoIncrement": 0, "db:isPrimaryKey": 0, "db:comment": "姓名"}`,
				`"age"`:  `{"db:type": "TinyInt", "db:size": 3, "db:precision": 0, "db:default": "", "db:indexName": "", "db:nullable": 0, "db:isUnique": 0, "db:isAutoIncrement": 0, "db:isPrimaryKey": 0, "db:comment": "年龄"}`,
			},
		}
		tableName = "gdbsm_test"
		dialect   = &MySQLDialect{}
		db        = g.DB()
	)

	if _, err := db.Exec(dialect.DropTableSQL(tableName)); err != nil {
		panic(err)
	}

	for idx, schema := range fieldSchemaGroup {
		t.Logf("test group %d", idx)
		var fields []string
		for field, opts := range schema {
			fields = append(fields, field+":"+opts)
		}
		if err := SyncFromJsonSchema(dialect, fmt.Sprintf(`{"schema":{"db:tableName": "%s","properties":{%s}}}`, tableName, strings.Join(fields, ","))); err != nil {
			panic(err)
		}
	}

	if _, err := db.Table(tableName).Insert(map[string]interface{}{
		"name": "beanwei",
		"age":  26,
	}); err != nil {
		panic(err)
	}

	if res, err := db.Table(tableName).All(); err != nil {
		panic(err)
	} else {
		t.Logf("res-all: %v", res)
	}
}
