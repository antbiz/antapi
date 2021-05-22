package global

import (
	"context"
	"sync"

	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/app/utils"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"go.mongodb.org/mongo-driver/bson"
)

// schema 常驻内存数据
var (
	schemaLocker sync.RWMutex
	schemasMap   map[string]*dto.Schema
)

// LoadSchemas 加载全部Collection的Schema
func LoadSchemas() {
	docs := make([]map[string]interface{}, 0)
	if err := db.DB().Collection("schema").Find(context.Background(), bson.M{}).All(&docs); err != nil {
		g.Log().Errorf("LoadSchemas schema read fail: %v", err)
		return
	}

	schemas := make([]*dto.Schema, len(docs))
	for i, doc := range docs {
		jsonDoc := gjson.New(doc)
		schemas[i] = utils.ParseFormRenderSchema(jsonDoc)
	}

	schemaLocker.Lock()
	defer schemaLocker.Unlock()

	schemasMap = make(map[string]*dto.Schema)
	for _, schema := range schemas {
		schemasMap[schema.Name] = schema
	}

	g.Log().Debug("LoadSchemas successfully!")
}

// GetSchema : 从内存中获取某个Collection的Schema
func GetSchema(collectionName string) *dto.Schema {
	return schemasMap[collectionName]
}
