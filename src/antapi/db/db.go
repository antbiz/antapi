package db

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	upperdb "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
	"github.com/upper/db/v4/adapter/mongo"
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"github.com/upper/db/v4/adapter/ql"
	"github.com/upper/db/v4/adapter/sqlite"
)

var (
	DB             upperdb.Session
	UseCockroachdb bool
	UseMongo       bool
	UseMssql       bool
	UseMysql       bool
	UsePostgresql  bool
	UseQl          bool
	UseSqlite      bool
)

func init() {
	driver := gstr.ToLower(g.Cfg().GetString("db.Driver"))
	link := g.Cfg().GetString("db.Link")
	switch driver {
	case "mysql":
		UseMysql = true
		if mysqlURL, err := mysql.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(mysqlURL); err != nil {
				panic(err)
			}
		}
	case "postgresql":
		UsePostgresql = true
		if pgsqlURL, err := postgresql.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(pgsqlURL); err != nil {
				panic(err)
			}
		}
	case "mssql":
		UseMssql = true
		if mssqlURL, err := mssql.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(mssqlURL); err != nil {
				panic(err)
			}
		}
	case "sqlite":
		UseSqlite = true
		if sqliteURL, err := sqlite.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(sqliteURL); err != nil {
				panic(err)
			}
		}
	case "ql":
		UseQl = true
		if qlURL, err := ql.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = mysql.Open(qlURL); err != nil {
				panic(err)
			}
		}
	case "cockroachdb":
		UseCockroachdb = true
		if cockroachdbURL, err := cockroachdb.ParseURL(link); err != nil {
			panic(err)
		} else {
			if DB, err = cockroachdb.Open(cockroachdbURL); err != nil {
				panic(err)
			}
		}
	case "mongo":
		UseMongo = true
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
