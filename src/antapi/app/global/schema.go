package global

import (
	"antapi/app/model"
	"antapi/app/utils"
	"sync"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

// schema 常驻内存数据
var (
	schemaLocker sync.RWMutex
	schemasMap   map[string]*model.Schema
)

// LoadSchemas 加载全部Collection的Schema
func LoadSchemas() error {
	db := g.DB()
	schemas := ([]*model.Schema)(nil)

	res, err := db.Table("schema").All()
	if err != nil {
		g.Log().Error("LoadSchemas schema read fail:", err)
		return err
	}
	for _, record := range res.Array() {
		jsonSchema, err := gjson.LoadContent(record.String())
		if err != nil {
			g.Log().Error("LoadSchemas schema json load fail:", err)
			return err
		}
		schemas = append(schemas, utils.ParseFormRenderSchema(jsonSchema)...)
	}

	schemaLocker.Lock()
	defer schemaLocker.Unlock()

	schemasMap = make(map[string]*model.Schema)
	for _, schema := range schemas {
		schemasMap[schema.CollectionName] = schema
	}

	glog.Info("LoadSchemas successfully!")
	return nil
}

// GetSchema : 从内存中获取某个Collection的Schema
func GetSchema(collectionName string) *model.Schema {
	return schemasMap[collectionName]
}
