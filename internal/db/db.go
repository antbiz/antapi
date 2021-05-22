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

// Init 初始化数据库连接
// 这里不直接使用 init 的方法的原因：方便单测和patches执行
func Init() (err error) {
	mongoURI := g.Cfg().GetString("mongo.uri")
	initCli.Do(func() {
		cli, err = qmgo.NewClient(
			context.Background(),
			&qmgo.Config{
				Uri: mongoURI,
			},
		)
		if err != nil {
			return
		}
	})
	g.Log().Debugf("Connected mongodb: %s", mongoURI)
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
	if cli == nil {
		Init()
	}
	return cli.Database(dbName)
}
