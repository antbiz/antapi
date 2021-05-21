package service

import (
	"context"

	"github.com/antbiz/antapi/internal/app/dao"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"go.mongodb.org/mongo-driver/bson"
)

var User = &userSrv{
	collectionName: "user",
}

type userSrv struct {
	collectionName string
}

// CollectionName .
func (srv *userSrv) CollectionName() string {
	return srv.collectionName
}

// EncryptPwd 加密账号密码
func (srv *userSrv) EncryptPwd(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}

// GetUserByLogin 根据 用户名/手机号/邮箱 + 密码 查询用户信息
func (srv *userSrv) GetUserByLogin(ctx context.Context, login, pwd string) (*gjson.Json, error) {
	doc, err := dao.Get(ctx, srv.collectionName, &dao.GetOptions{
		Filter: bson.D{{"$or", bson.D{{"username", login}, {"phone", login}, {"email", login}}}},
	})
	if err != nil {
		return nil, err
	}

	username := doc.GetString("username")
	password := doc.GetString("password")
	if srv.EncryptPwd(username, password) != password {
		return nil, nil
	}
	return doc, nil
}
