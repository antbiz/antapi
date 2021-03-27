package logic

import (
	"antapi/app/model"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

var Project = new(projectLogic)

type projectLogic struct{}

// AutoExportProjectData 保存数据到 app/model/project 以便项目初始化
func (projectLogic) AutoExportProjectData(data *gjson.Json) error {
	glog.Info("Auto Export Project Data To app/model/project")
	exportPath := gfile.Join(gfile.Pwd(), "app", "model", "project", data.GetString("name"))

	exportData := data.Map()
	if exportData != nil {
		for _, fieldName := range model.DefaultFieldNames {
			delete(exportData, fieldName)
		}
		if err := gfile.PutContents(exportPath, gjson.New(exportData).MustToJsonString()); err != nil {
			glog.Fatal(err)
		}
	}

	return nil
}

// AutoDeleteExportedJsonFile 删除导出的json文件
func (projectLogic) AutoDeleteExportedJsonFile(data *gjson.Json) error {
	glog.Info("Auto Delete Exported Json File From app/model/project")
	jsonFilePath := gfile.Join(gfile.Pwd(), "app", "model", "project", data.GetString("name"))

	if err := gfile.Remove(jsonFilePath); err != nil {
		glog.Fatal(err)
	}

	return nil
}
