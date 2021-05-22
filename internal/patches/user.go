package patches

import (
	"context"
	"time"

	"github.com/antbiz/antapi/internal/app/service"
	"github.com/antbiz/antapi/internal/db"
	"github.com/gogf/gf/frame/g"
	"go.mongodb.org/mongo-driver/bson"
)

// initAdminAccount 初始化超级管理员账号
func initAdminAccount() {
	username := g.Cfg().GetString("admin.username")
	password := g.Cfg().GetString("admin.password")
	email := g.Cfg().GetString("admin.email")
	if username == "" || password == "" {
		panic("admin username or password is empty")
	}
	password = service.User.EncryptPwd(username, password)
	_, err := db.
		DB().
		Collection(service.User.CollectionName()).
		Upsert(context.Background(), bson.M{"username": username}, g.Map{
			"createdAt": time.Now().Unix(),
			"updatedAt": time.Now().Unix(),
			"username":  username,
			"password":  password,
			"isSysUser": true,
			"email":     email,
		})
	if err != nil {
		panic(err)
	}
	g.Log().Debug("Init admin account successfully!")
}
