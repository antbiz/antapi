package service

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

type baseSysSrv struct {
	collectionName string
}

// CollectionName .
func (srv *baseSysSrv) CollectionName() string {
	return srv.collectionName
}

func (srv *baseSysSrv) DefaultFieldNames() []string {
	return []string{"_id", "createdAt", "updatedAt", "createdBy", "updatedBy"}
}

func (srv *baseSysSrv) ExportJSONFile(data *gjson.Json, path string) error {
	g.Log().Infof("Auto Export Json File To %s", path)
	// 将data复制一份
	_data := new(gjson.Json)
	*_data = *data
	for _, fieldName := range srv.DefaultFieldNames() {
		_data.Remove(fieldName)
	}
	if err := gfile.PutContents(path, _data.MustToJsonIndentString()); err != nil {
		g.Log().Error(err)
	}
	return nil
}

func (srv *baseSysSrv) DeleteJSONFile(path string) error {
	g.Log().Infof("Auto Delete Exported Json File From %s", path)
	if err := gfile.Remove(path); err != nil {
		g.Log().Fatal(err)
	}
	return nil
}
