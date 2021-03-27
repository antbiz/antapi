package logic

import (
	"antapi/app/model"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

var Project = new(projectLogic)

type projectLogic struct{}

// AutoExportProjectData 保存数据到 app/model/project 以便项目初始化
func (projectLogic) AutoExportProjectData(data *gjson.Json) error {
	glog.Info("Auto Export Project Data To app/model/project")
	exportPath := gfile.Join(gfile.Pwd(), "app", "model", "project", fmt.Sprintf("%s.json", data.GetString("name")))

	// 将data复制一份
	_data := new(gjson.Json)
	*_data = *data

	for _, fieldName := range model.DefaultFieldNames {
		_data.Remove(fieldName)
	}
	if err := gfile.PutContents(exportPath, _data.MustToJsonIndentString()); err != nil {
		glog.Fatal(err)
	}

	return nil
}

// AutoDeleteExportedJsonFile 删除导出的json文件
func (projectLogic) AutoDeleteExportedJsonFile(data *gjson.Json) error {
	glog.Info("Auto Delete Exported Json File From app/model/project")
	jsonFilePath := gfile.Join(gfile.Pwd(), "app", "model", "project", fmt.Sprintf("%s.json", data.GetString("name")))

	if err := gfile.Remove(jsonFilePath); err != nil {
		glog.Fatal(err)
	}

	return nil
}
