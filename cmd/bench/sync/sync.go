package sync

import (
	"context"
	"fmt"

	"github.com/antbiz/antapi/internal/app/dao"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"go.mongodb.org/mongo-driver/bson"
)

// Help .
func Help() {
	fmt.Println(`
USAGE
    bench sync
ARGUMENT
    OPTION  
OPTION
EXAMPLES
	go run cmd/bench/bench.go sync
DESCRIPTION
    The "sync" command is used for sync schema/projects/defaults`)
}

// Run .
func Run() {
	ctx := context.Background()
	syncSchemas(ctx)
	syncProjects(ctx)
}

var folders = []string{"sys", "biz"}

func syncSchemas(ctx context.Context) {
	g.Log().Info("Sync Schemas")
	for _, folder := range folders {
		schemaFilePath := gfile.Join(gfile.Pwd(), "internal", "data", "schemas", folder)
		schemaFileNames, err := gfile.DirNames(schemaFilePath)
		if err != nil {
			g.Log().Fatal("Scan Schemas Dir %s Error: %v", schemaFilePath, err)
		}

		for _, fileName := range schemaFileNames {
			if !gstr.HasSuffix(fileName, ".json") {
				continue
			}
			g.Log().Infof("Sync Schema Data %s", fileName)
			jsonDoc, err := gjson.Load(gfile.Join(schemaFilePath, fileName))
			if err != nil {
				g.Log().Fatalf("Sync Schema Data %s Error: %v", fileName, err)
			}

			_, err = dao.Upsert(ctx, "schema", jsonDoc, &dao.UpsertOptions{
				Filter: bson.M{"collectionName": jsonDoc.GetString("collectionName")},
			})
			if err != nil {
				g.Log().Fatal(err)
			}
		}
	}
}

func syncProjects(ctx context.Context) {
	g.Log().Info("Sync Projects")
	for _, folder := range folders {
		projectFilePath := gfile.Join(gfile.Pwd(), "internal", "data", "projects", folder)
		projectFileNames, err := gfile.DirNames(projectFilePath)
		if err != nil {
			g.Log().Fatal("Scan Projects Dir %s Error: %v", projectFilePath, err)
		}

		for _, fileName := range projectFileNames {
			if !gstr.HasSuffix(fileName, ".json") {
				continue
			}
			g.Log().Infof("Sync Project Data %s", fileName)
			jsonDoc, err := gjson.Load(gfile.Join(projectFilePath, fileName))
			if err != nil {
				g.Log().Fatalf("Sync Project Data %s Error: %v", fileName, err)
			}

			_, err = dao.Upsert(ctx, "project", jsonDoc, &dao.UpsertOptions{
				Filter: bson.M{"name": jsonDoc.GetString("name")},
			})
			if err != nil {
				g.Log().Fatal(err)
			}
		}
	}
}
