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
	SyncCollections()
	SyncSchemas()
	SyncProjects()
	SyncDefaultsData()
}

var modules = []string{"default", "biz"}

// SyncCollections : 读取collection，更新对应数据库表
func SyncCollections() {
	glog.Info("Migrate database tables")
	for _, module := range modules {
		collectionFilePath := gfile.Join(gfile.Pwd(), "app", "model", "collection", module)
		collectionFileNames, err := gfile.DirNames(collectionFilePath)
		if err != nil {
			glog.Fatalf("Scan Collection Dir %s Error: %v", collectionFilePath, err)
		}

		for _, fileName := range collectionFileNames {
			if !gstr.HasSuffix(fileName, ".json") {
				continue
			}
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
}

// SyncSchemas : 同步collection的schema到 `schema` 数据表
func SyncSchemas() {
	glog.Info("Update table schema data")
	db := g.DB()
	tx, err := db.Begin()
	if err != nil {
		glog.Fatalf("DB Begin Error: %v", err)
	}
	for _, module := range modules {
		collectionFilePath := gfile.Join(gfile.Pwd(), "app", "model", "collection", module)
		collectionFileNames, err := gfile.DirNames(collectionFilePath)
		if err != nil {
			glog.Fatalf("Scan Collection Dir %s Error: %v", collectionFilePath, err)
			return
		}

		for _, fileName := range collectionFileNames {
			if !gstr.HasSuffix(fileName, ".json") {
				continue
			}
			glog.Infof("SyncSchema %s", fileName)
			j, err := gjson.Load(gfile.Join(collectionFilePath, fileName))
			if err != nil {
				glog.Fatalf("SyncSchemas %s Error: %v", fileName, err)
			}

			// 填充默认字段
			if err := logic.Schema.CheckFields(j); err != nil {
				glog.Fatalf("SyncSchemas %s Error When CheckFields: %v", fileName, err)
			}

			schema := new(model.Schema)
			if err := j.Struct(schema); err != nil {
				glog.Fatalf("SyncSchemas %s Error: %v", fileName, err)
			}

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

			for idx, field := range schema.Fields {
				fieldID, err := db.Table("schema_field").Value("id", g.Map{"pid": schema.ID, "name": field.Name})
				if err != nil {
					glog.Fatalf("Find SchemaField %s-%s Error: %v", schema.Name, field.Name, err)
				}
				field.Pid = schema.ID
				field.Pcn = "schema"
				field.Pfd = "fields"
				field.Idx = idx

				if fieldID.IsEmpty() {
					field.ID = guid.S()
					if _, err := db.Table("schema_field").Data(field).Insert(); err != nil {
						tx.Rollback()
						glog.Fatalf("Create SchemaField %s Error: %v", field.Name, err)
					}
				} else {
					field.ID = fieldID.String()
					if _, err := db.Table("schema_field").Where("id", field.ID).Data(field).Update(); err != nil {
						tx.Rollback()
						glog.Fatalf("Update SchemaField %s-%s Error: %v", field.Name, field.ID, err)
					}
				}
			}
		}
	}
	tx.Commit()
}

// SyncProjects : 同步projects数据到 `project` 数据表
func SyncProjects() {
	glog.Info("Update table project data")
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
