package gdbsm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gogf/gf/frame/g"
)

func TestSyncFromJsonSchema(t *testing.T) {
	var (
		fieldSchemaGroup = []map[string]string{
			{
				`"a"`: `{"db:type": "VARCHAR", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "", "db:nullable": 0, "db:isUnique": 0, "db:isAutoIncrement": 0, "db:isPrimaryKey": 0, "db:comment": "a"}`,
			},
			{
				`"a"`: `{"db:type": "VARCHAR", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "", "db:nullable": 0, "db:isUnique": 1, "db:isAutoIncrement": 0, "db:isPrimaryKey": 0, "db:comment": "a"}`,
				`"b"`: `{"db:type": "Int", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "", "db:nullable": 0, "db:isUnique": 0, "db:isAutoIncrement": 0, "db:isPrimaryKey": 0, "db:comment": "b"}`,
			},
			{
				`"a"`: `{"db:type": "VARCHAR", "db:size": 0, "db:precision": 0, "db:default": "foo", "db:indexName": "ab", "db:nullable": 0, "db:isUnique": 0, "db:isAutoIncrement": 0, "db:isPrimaryKey": 0, "db:comment": "a"}`,
				`"b"`: `{"db:type": "Int", "db:size": 0, "db:precision": 0, "db:default": "", "db:indexName": "ab", "db:nullable": 0, "db:isUnique": 0, "db:isAutoIncrement": 0, "db:isPrimaryKey": 0, "db:comment": "b"}`,
			},
		}
		tableName = "gdbsm_test"
		dialect   = &MySQLDialect{}
	)

	g.Config().SetPath("C:\\BeanWei\\Workspace\\antbiz\\antapi\\config")

	if _, err := g.DB().Exec(dialect.DropTableSQL(tableName)); err != nil {
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
}
