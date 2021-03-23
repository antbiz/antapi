package logic

import (
	"antapi/app/dao"
	"antapi/app/errcode"

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
func (userLogic) GetUserByLogin(login, pwd string) (data *gjson.Json, err error) {
	if data, err = dao.Get("user", "username=? or phone=? or email=?", g.Slice{login, login, login}); err != nil {
		return
	}
	username := data.GetString("username")
	password := data.GetString("password")
	if userLogic.EncryptPwd(username, pwd) != password {
		return nil, gerror.NewCode(errcode.IncorrectUsernameOrPassword, errcode.IncorrectUsernameOrPasswordMsg)
	}
	return data, nil
}

// // SignIn 用户登录
// func (userLogic) SignIn(ctx context.Context, signInReq *model.UserSignInReq) error {

// }
