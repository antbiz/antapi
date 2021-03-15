package db

import (
	"antapi/db/types"
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	upperdb "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mongo"
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"github.com/upper/db/v4/adapter/sqlite"
)

var (
	DB     upperdb.Session
	DBType types.DBType
)

func init() {
	driver := gstr.ToLower(g.Cfg().GetString("db.Driver"))
	link := g.Cfg().GetString("db.Link")
	DBType = types.DBType(driver)
	switch DBType {
	case types.MYSQL:
		if mysqlURL, err := mysql.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(mysqlURL); err != nil {
				panic(err)
			}
		}
	case types.POSTGRES:
		if pgsqlURL, err := postgresql.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(pgsqlURL); err != nil {
				panic(err)
			}
		}
	case types.MSSQL:
		if mssqlURL, err := mssql.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(mssqlURL); err != nil {
				panic(err)
			}
		}
	case types.SQLITE:
		if sqliteURL, err := sqlite.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(sqliteURL); err != nil {
				panic(err)
			}
		}
	case types.MONGO:
		if mongoURL, err := mongo.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mongo.Open(mongoURL); err != nil {
				panic(err)
			}
		}
	default:
		panic(fmt.Errorf("unknown database driver: %s", driver))
	}
}
