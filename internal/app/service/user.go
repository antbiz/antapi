package service

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
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
func (srv *userSrv) GetUserByLogin(login, pwd string) (record *gjson.Json, err error) {
	return
}
