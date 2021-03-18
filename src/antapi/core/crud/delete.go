package crud

import "github.com/gogf/gf/frame/g"

// Delete : 删除指定数据
func Delete(collectionName string, where interface{}, args ...interface{}) error {
	db := g.DB()
	if _, err := db.Table(collectionName).Where(where, args...).Delete(); err != nil {
		return err
	}
	return nil
}
