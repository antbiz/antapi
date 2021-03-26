package global

import (
	"antapi/app/model"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gmode"
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
	schemaFields := ([]*model.SchemaField)(nil)

	if err := db.Table("schema").Structs(&schemas); err != nil {
		glog.Error("LoadSchemas schema read fail:", err)
		return err
	}

	if err := db.Table("schema_field").Order("idx asc").Structs(&schemaFields); err != nil {
		glog.Error("LoadSchemas schema_field read fail:", err)
		return err
	}

	schemaLocker.Lock()
	defer schemaLocker.Unlock()

	_schemasMap := map[string]*model.Schema{}
	for _, schema := range schemas {

		for _, field := range schemaFields {
			if field.Pid == schema.ID {
				schema.Fields = append(schema.Fields, field)
			}
		}

		_schemasMap[schema.Name] = schema
	}
	schemasMap = _schemasMap

	glog.Info("LoadSchemas successfully!")
	if gmode.IsDevelop() {
		g.Dump(_schemasMap)
	}
	return nil
}

// GetSchema : 从内存中获取某个Collection的Schema
func GetSchema(collectionName string) *model.Schema {
	return schemasMap[collectionName]
}
