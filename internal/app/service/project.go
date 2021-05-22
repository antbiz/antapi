package service

import (
	"context"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
)

var Project = &projectSrv{}

type projectSrv struct {
	baseSysSrv
}

// CollectionName .
func (srv *projectSrv) CollectionName() string {
	return "project"
}

// GetJSONFilePath 获取文件备份导出路径
func (srv *projectSrv) GetJSONFilePath(projectName string) string {
	return gfile.Join(gfile.Pwd(), "internal", "data", "schemas", "biz", fmt.Sprintf("%s.json", projectName))
}

// AutoExportJSONFile .
func (srv *projectSrv) AutoExportJSONFile(ctx context.Context, data *gjson.Json) error {
	return srv.ExportJSONFile(data, srv.GetJSONFilePath(data.GetString("name")))
}

// AutoDeleteJSONFile .
func (srv *projectSrv) AutoDeleteJSONFile(ctx context.Context, data *gjson.Json) error {
	return srv.DeleteJSONFile(srv.GetJSONFilePath(data.GetString("name")))
}
