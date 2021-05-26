package db

import (
	"context"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/qmgo"
)

var (
	cli  *qmgo.Client
	once sync.Once
)

// Cli is mongo client
func Cli() *qmgo.Client {
	once.Do(func() {
		var err error
		mongoURI := g.Cfg().GetString("mongo.uri")
		cli, err = qmgo.NewClient(
			context.Background(),
			&qmgo.Config{
				Uri: mongoURI,
			},
		)
		if err != nil {
			g.Log().Errorf("failed to connect mongo: %s", mongoURI)
		} else {
			g.Log().Debugf("connected mongo: %s", mongoURI)
		}
	})
	return cli
}

// DB 数据库实例
func DB(database ...string) *qmgo.Database {
	var dbName string
	if len(database) > 0 {
		dbName = database[0]
	} else {
		dbName = g.Cfg().GetString("mongo.default")
	}
	return Cli().Database(dbName)
}
