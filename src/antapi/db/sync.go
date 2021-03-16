package db

import (
	"antapi/db/dbsm"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

// JSONCollection2TableSchema : collection数据转数据表机构
func JSONCollection2TableSchema(collection *gjson.Json) *dbsm.Table {
	tableName := collection.GetString("name")
	columns := make([]*dbsm.Column, 0)
	for _, field := range collection.GetJsons("fields") {
		col := &dbsm.Column{
			Name:     field.GetString("name"),
			Type:     field.GetString("type"),
			Default:  field.GetString("default"),
			Nullable: true,
			IsUnique: field.GetBool("is_unique"),
			Comment:  field.GetString("description"),
		}

		if field.GetBool("create_index") {
			col.IndexName = col.Name
		}
		if col.Name == "id" {
			col.IsAutoIncrement = true
			col.IsPrimaryKey = true
		}

		columns = append(columns, col)
	}
	return dbsm.NewTable(tableName, columns)
}

// SyncCollections : 同步collections
func SyncCollections() error {
	collectionFilePath := fmt.Sprintf("%s/db/collection", gfile.MainPkgPath())
	collectionFileNames, err := gfile.DirNames(collectionFilePath)
	if err != nil {
		return err
	}

	tables := make([]*dbsm.Table, 0)
	for _, fileName := range collectionFileNames {
		glog.Debugf("SyncCollections %s", fileName)
		collection, err := gjson.Load(fmt.Sprintf("%s/%s.json", collectionFilePath, fileName))
		if err != nil {
			glog.Debugf("SyncCollections %s Error: %v", fileName, err)
			continue
		}
		tables = append(tables, JSONCollection2TableSchema(collection))
	}

	return nil
}

// SyncDefaultsData : 初始化默认数据
func SyncDefaultsData() error {
	return nil
}
