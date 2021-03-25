package cmd

import (
	"antapi/app/logic"
	"antapi/app/model"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/guid"
)

// Sync : 同步collection和默认数据
func Sync() {
	SyncCollections()
	SyncSchemas()
	SyncProjects()
	SyncDefaultsData()
}

// SyncCollections : 读取collection，更新对应数据库表
func SyncCollections() {
	glog.Info("Migrate database tables")
	collectionFilePath := gfile.Join(gfile.Pwd(), "app", "model", "collection")
	collectionFileNames, err := gfile.DirNames(collectionFilePath)
	if err != nil {
		glog.Fatalf("Scan Collection Dir %s Error: %v", collectionFilePath, err)
	}

	for _, fileName := range collectionFileNames {
		glog.Infof("SyncCollections %s", fileName)
		j, err := gjson.Load(gfile.Join(collectionFilePath, fileName))
		if err != nil {
			glog.Fatalf("SyncCollections %s Error: %v", fileName, err)
		}
		if err := logic.Schema.CheckFields(j); err != nil {
			glog.Fatalf("Check Collections %s Fields Error: %v", fileName, gerror.Cause(err))
		}
		if err := logic.Schema.MigrateCollectionSchema(j); err != nil {
			glog.Fatalf("Migrate Collections %s Error: %v", fileName, gerror.Cause(err))
		}
	}
}

// SyncSchemas : 同步collection的schema到 `schema` 数据表
func SyncSchemas() {
	glog.Info("Update table schema data")
	collectionFilePath := gfile.Join(gfile.Pwd(), "app", "model", "collection")
	collectionFileNames, err := gfile.DirNames(collectionFilePath)
	if err != nil {
		glog.Fatalf("Scan Collection Dir %s Error: %v", collectionFilePath, err)
		return
	}

	db := g.DB()
	for _, fileName := range collectionFileNames {
		glog.Fatalf("SyncSchemas %s", fileName)
		j, err := gjson.Load(gfile.Join(collectionFilePath, fileName))
		if err != nil {
			glog.Fatalf("SyncSchemas %s Error: %v", fileName, err)
		}
		schema := new(model.Schema)
		if err := j.Struct(schema); err != nil {
			glog.Fatalf("SyncSchemas %s Error: %v", fileName, err)
		}
		glog.Infof("SyncSchema %s", fileName)

		schemaID, err := db.Table("schema").Value("id", "name", schema.Name)
		if err != nil {
			glog.Fatalf("Find Schema %s Error: %v", schema.Name, err)
		}
		if schemaID.IsEmpty() {
			schema.ID = guid.S()
			if _, err := db.Table("schema").Data(schema).Insert(); err != nil {
				glog.Fatalf("Create Schema %s Error: %v", schema.Name, err)
			}
		} else {
			schema.ID = schemaID.String()
			if _, err := db.Table("schema").Where("id", schema.ID).Data(schema).Update(); err != nil {
				glog.Fatalf("Update Schema %s-%s Error: %v", schema.Name, schema.ID, err)
			}
		}

		for _, field := range schema.Fields {
			fieldID, err := db.Table("schema_field").Value("id", "pid", schema.ID)
			if err != nil {
				glog.Fatalf("Find SchemaField %s-%s Error: %v", schema.Name, field.Name, err)
			}
			if fieldID.IsEmpty() {
				field.ID = guid.S()
				if _, err := db.Table("schema").Data(field).Insert(); err != nil {
					glog.Fatalf("Create SchemaField %s Error: %v", field.Name, err)
				}
			} else {
				field.ID = fieldID.String()
				if _, err := db.Table("schema_field").Where("id", field.ID).Data(field).Update(); err != nil {
					glog.Fatalf("Update SchemaField %s-%s Error: %v", field.Name, field.ID, err)
				}
			}
		}
	}
}

// SyncProjects : 同步projects数据到 `project` 数据表
func SyncProjects() {
	glog.Info("Update table project data")
	projectFilePath := gfile.Join(gfile.Pwd(), "app", "model", "project")
	projectFileNames, err := gfile.DirNames(projectFilePath)
	if err != nil {
		glog.Fatalf("Scan Project Dir %s Error: %v", projectFilePath, err)
		return
	}

	db := g.DB()
	for _, fileName := range projectFileNames {
		j, err := gjson.Load(gfile.Join(projectFilePath, fileName))
		if err != nil {
			glog.Fatalf("SyncProjects %s Error: %v", fileName, err)
		}
		project := new(model.Project)
		if err := j.Struct(project); err != nil {
			glog.Fatalf("SyncProjects %s Error: %v", fileName, err)
		}
		glog.Infof("SyncProjects %s", fileName)

		projectID, err := db.Table("project").Value("id", "project", project.Name)
		if err != nil {
			glog.Fatalf("Find Project %s Error: %v", project.Name, err)
		}
		if projectID.IsEmpty() {
			project.ID = guid.S()
			if _, err := db.Table("project").Data(project).Insert(); err != nil {
				glog.Fatalf("Create Project %s Error: %v", project.Name, err)
			}
		} else {
			project.ID = projectID.String()
			if _, err := db.Table("project").Where("id", project.ID).Data(project).Update(); err != nil {
				glog.Fatalf("Update Project %s-%s Error: %v", project.Name, project.ID, err)
			}
		}
	}
}

// SyncDefaultsData : 初始化默认数据
func SyncDefaultsData() {
}
