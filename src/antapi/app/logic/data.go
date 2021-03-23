package logic

import (
	"antapi/app/model"
	"sync"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

// 常驻内存数据
var (
	schemaLocker sync.RWMutex
	schemasMap   map[string]*model.Schema
)

// LoadSchemas : 加载全部Collection的Schema
func LoadSchemas() error {
	db := g.DB()
	schemas := ([]*model.Schema)(nil)

	if err := db.Table("schema").Structs(&schemas); err != nil {
		glog.Error("LoadSchemas schema read fail:", err)
		return err
	}

	if err := db.Table("schema_field").Where("pid", gdb.ListItemValuesUnique(schemas, "ID")).ScanList(&schemas, "Fields", "pid:ID"); err != nil {
		glog.Error("LoadSchemas schema_field read fail:", err)
		return err
	}

	schemaLocker.Lock()
	defer schemaLocker.Unlock()

	for _, schema := range schemas {
		schemasMap[schema.Name] = schema
	}

	glog.Info("LoadSchemas successfully!")
	return nil
}

// GetSchema : 从内存中获取某个Collection的Schema
func GetSchema(collectionName string) *model.Schema {
	return schemasMap[collectionName]
}
