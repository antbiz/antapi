package service

import (
	"fmt"

	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

var Project = &projectSrv{
	collectionName: "project",
}

type projectSrv struct {
	collectionName string
}

// CollectionName .
func (srv *projectSrv) CollectionName() string {
	return srv.collectionName
}

// GetJSONFilePath 获取文件备份导出路径
func (srv *projectSrv) GetJSONFilePath(projectName string) string {
	return gfile.Join(gfile.Pwd(), "internal", "data", "schemas", "biz", fmt.Sprintf("%s.json", projectName))
}

// AutoExportJSONFile 保存数据到 app/model/project/biz 以便项目初始化
func (srv *projectSrv) AutoExportJSONFile(data *gjson.Json) error {
	g.Log().Info("Auto Export Project Data To app/model/project/biz")
	jsonFilePath := srv.GetJSONFilePath(data.GetString("name"))

	// 将data复制一份
	_data := new(gjson.Json)
	*_data = *data

	for _, fieldName := range dto.DefaultFieldNames {
		_data.Remove(fieldName)
	}
	if err := gfile.PutContents(jsonFilePath, _data.MustToJsonIndentString()); err != nil {
		g.Log().Fatal(err)
	}

	return nil
}

// AutoDeleteExportedJsonFile 删除导出的json文件
func (srv *projectSrv) AutoDeleteJSONFile(data *gjson.Json) error {
	g.Log().Info("Auto Delete Exported Json File From app/model/project/biz")
	jsonFilePath := srv.GetJSONFilePath(data.GetString("name"))

	if err := gfile.Remove(jsonFilePath); err != nil {
		g.Log().Fatal(err)
	}

	return nil
}
