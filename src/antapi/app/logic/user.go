package logic

import (
	"antapi/app/dao"
	"antapi/app/model"
	"antapi/common/errcode"
	"context"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

var User = new(userLogic)

type userLogic struct{}

// EncryptPwd 加密账号密码
func (userLogic) EncryptPwd(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}

// GetUserByLogin 根据 用户名/手机号/邮箱 + 密码 查询用户信息
func (self *userLogic) GetUserByLogin(login, pwd string) (data *gjson.Json, err error) {
	arg := &dao.GetFuncArg{
		Where:               "username=? or phone=? or email=?",
		WhereArgs:           g.Slice{login, login, login},
		IncludeHiddenField:  true,
		IncludePrivateField: true,
	}
	if data, err = dao.Get("user", arg); err != nil {
		return
	}
	username := data.GetString("username")
	password := data.GetString("password")
	if self.EncryptPwd(username, pwd) != password || data == nil {
		return nil, gerror.NewCode(errcode.IncorrectUsernameOrPassword, errcode.IncorrectUsernameOrPasswordMsg)
	}
	return data, nil
}

// SignIn 用户登录
func (self *userLogic) SignIn(ctx context.Context, signInReq *model.UserSignInReq) error {
	// userData, err := self.GetUserByLogin(signInReq.Login, signInReq.Pwd)
	// if err != nil {
	// 	return err
	// }
	// user := &model.UserSession{
	// 	Username:  userData.GetString("username"),
	// 	Phone:     userData.GetString("phone"),
	// 	Email:     userData.GetString("email"),
	// 	Blocked:   userData.GetString("blocked"),
	// 	IsSysuser: userData.GetString("is_sysuser"),
	// }
	return nil
}
