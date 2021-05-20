package db

import (
	"context"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/qmgo"
)

var (
	cli     *qmgo.Client
	initCli sync.Once
)

// Init  初始化数据库连接
func Init() (err error) {
	initCli.Do(func() {
		cli, err = qmgo.NewClient(
			context.Background(),
			&qmgo.Config{
				Uri: g.Cfg().GetString("mongo.default.uri"),
			},
		)
		if err != nil {
			return
		}
	})
	return
}

// DB 数据库实例
func DB(database ...string) *qmgo.Database {
	var dbName string
	if len(database) > 0 {
		dbName = database[0]
	} else {
		dbName = g.Cfg().GetString("mongo.default.database")
	}
	return cli.Database(dbName)
}
