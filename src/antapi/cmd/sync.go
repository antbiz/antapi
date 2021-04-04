package cmd

import (
	"antapi/app/logic"
	"antapi/app/model"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/guid"
)

// Sync : 同步collection和默认数据
func Sync() {
	SyncSchemas()
	SyncSchemasData()
	SyncProjectsData()
	SyncDefaultsData()
}

var modules = []string{"default", "biz"}

// SyncSchemas : 读取schema，更新对应数据库表
func SyncSchemas() {
	glog.Info("Sync Schemas")
	for _, module := range modules {
		schemaFilePath := gfile.Join(gfile.Pwd(), "boot", "schemas", module)
		schemaFileNames, err := gfile.DirNames(schemaFilePath)
		if err != nil {
			glog.Fatalf("Scan Schema Dir %s Error: %v", schemaFilePath, err)
		}

		for _, fileName := range schemaFileNames {
			if !gstr.HasSuffix(fileName, ".json") {
				continue
			}
			glog.Infof("Sync Schema %s", fileName)
			j, err := gjson.Load(gfile.Join(schemaFilePath, fileName))
			if err != nil {
				glog.Fatalf("SyncCollections %s Error: %v", fileName, err)
			}
			if err := logic.Schema.CheckJSONSchema(j); err != nil {
				glog.Fatalf("Check Collections %s Fields Error: %v", fileName, gerror.Cause(err))
			}
			if err := logic.Schema.MigrateSchema(j); err != nil {
				glog.Fatalf("Migrate Collections %s Error: %v", fileName, gerror.Cause(err))
			}
		}
	}
}

// SyncSchemasData : 同步collection的schema到 `schema` 数据表
func SyncSchemasData() {
	glog.Info("Sync Schemas Data")
	db := g.DB()

	for _, module := range modules {
		schemaFilePath := gfile.Join(gfile.Pwd(), "boot", "schemas", module)
		schemaFileNames, err := gfile.DirNames(schemaFilePath)
		if err != nil {
			glog.Fatalf("Scan Schema Dir %s Error: %v", schemaFilePath, err)
			return
		}

		for _, fileName := range schemaFileNames {
			if !gstr.HasSuffix(fileName, ".json") {
				continue
			}
			glog.Infof("Sync Schema Data %s", fileName)
			j, err := gjson.Load(gfile.Join(schemaFilePath, fileName))
			if err != nil {
				glog.Fatalf("Sync Schema Data %s Error: %v", fileName, err)
			}

			// 填充默认字段
			if err := logic.Schema.CheckJSONSchema(j); err != nil {
				glog.Fatalf("Sync Schema Data %s Error When CheckJSONSchema: %v", fileName, err)
			}

			schema := new(model.JSONSchema)
			if err := j.Struct(schema); err != nil {
				glog.Fatalf("Sync Schema Data %s Error: %v", fileName, err)
			}

			schemaID, err := db.Table("schema").Value("id", "name", schema.CollectionName)
			if err != nil {
				glog.Fatalf("Find Schema %s Error: %v", schema.CollectionName, err)
			}
			if schemaID.IsEmpty() {
				schema.ID = guid.S()
				if _, err := db.Table("schema").Data(schema).Insert(); err != nil {
					glog.Fatalf("Create Schema %s Error: %v", schema.CollectionName, err)
				}
			} else {
				schema.ID = schemaID.String()
				if _, err := db.Table("schema").Where("id", schema.ID).Data(schema).Update(); err != nil {
					glog.Fatalf("Update Schema %s-%s Error: %v", schema.CollectionName, schema.ID, err)
				}
			}
		}
	}
}

// SyncProjectsData : 同步projects数据到 `project` 数据表
func SyncProjectsData() {
	glog.Info("Sync Projects Data")
	db := g.DB()
	for _, module := range modules {
		projectFilePath := gfile.Join(gfile.Pwd(), "app", "model", "project", module)
		projectFileNames, err := gfile.DirNames(projectFilePath)
		if err != nil {
			glog.Fatalf("Scan Project Dir %s Error: %v", projectFilePath, err)
			return
		}
		for _, fileName := range projectFileNames {
			if !gstr.HasSuffix(fileName, ".json") {
				continue
			}
			j, err := gjson.Load(gfile.Join(projectFilePath, fileName))
			if err != nil {
				glog.Fatalf("SyncProjects %s Error: %v", fileName, err)
			}

			project := new(model.Project)
			if err := j.Struct(project); err != nil {
				glog.Fatalf("SyncProjects %s Error: %v", fileName, err)
			}
			glog.Infof("SyncProjects %s", fileName)

			projectID, err := db.Table("project").Value("id", "name", project.Name)
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
}

// SyncDefaultsData : 初始化默认数据
func SyncDefaultsData() {
}
