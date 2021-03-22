package cmd

import (
	"antapi/model"
	"antapi/pkg/dbsm"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/guid"
)

// Sync : 同步collection和默认数据
func Sync() error {
	if err := SyncCollections(); err != nil {
		return err
	}
	if err := SyncDefaultsData(); err != nil {
		return err
	}
	return nil
}

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

		if field.GetBool("can_index") {
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
	collectionFilePath := fmt.Sprintf("%s/model/collection", gfile.MainPkgPath())
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

// SyncProjects : 同步projects
func SyncProjects() {
	projectFilePath := fmt.Sprintf("%s/model/project", gfile.MainPkgPath())
	projectFileNames, err := gfile.DirNames(projectFilePath)
	if err != nil {
		glog.Debugf("Scan Projects Dir Error: %v", err)
		return
	}

	db := g.DB()
	for _, fileName := range projectFileNames {
		j, err := gjson.Load(fmt.Sprintf("%s/%s.json", projectFilePath, fileName))
		if err != nil {
			glog.Debugf("SyncProjects %s Error: %v", fileName, err)
			continue
		}
		project := new(model.Project)
		if err := j.Struct(project); err != nil {
			glog.Debugf("SyncProjects %s Error: %v", fileName, err)
			continue
		}
		glog.Debugf("SyncProjects %s", fileName)

		val, err := db.Table("project").Value("id", "project", project.Name)
		if err != nil {
			glog.Debugf("Find Project %s Error: %v", project.Name, err)
			continue
		}
		if val.IsEmpty() {
			project.ID = guid.S()
			if _, err := db.Table("project").Data(project).Insert(); err != nil {
				glog.Debugf("Create Project %s Error: %v", project.Name, err)
				continue
			}
		} else {
			if _, err := db.Table("project").Where("id", val.String()).Data(project).Update(); err != nil {
				glog.Debugf("Update Project %s-%s Error: %v", project.Name, val.String(), err)
				continue
			}
		}
	}
}

// SyncDefaultsData : 初始化默认数据
func SyncDefaultsData() error {
	return nil
}
