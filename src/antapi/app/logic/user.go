package logic

import (
	"antapi/app/dao"
	"antapi/app/model"
	"antapi/common/errcode"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var User = new(userLogic)

type userLogic struct{}

// EncryptPwd 加密账号密码
func (userLogic) EncryptPwd(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}

// GetUserByLogin 根据 用户名/手机号/邮箱 + 密码 查询用户信息
func (u *userLogic) GetUserByLogin(login, pwd string) (data *gjson.Json, err error) {
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
	if u.EncryptPwd(username, pwd) != password || data == nil {
		return nil, gerror.NewCode(errcode.IncorrectUsernameOrPassword, errcode.IncorrectUsernameOrPasswordMsg)
	}
	return data, nil
}

// SignIn 用户登录
func (u *userLogic) SignIn(signInReq *model.UserSignInReq, r *ghttp.Request) (map[string]interface{}, error) {
	data, err := u.GetUserByLogin(signInReq.Login, signInReq.Pwd)
	if err != nil {
		return nil, err
	}

	sessionData := g.Map{
		"id":         data.GetString("id"),
		"username":   data.GetString("username"),
		"phone":      data.GetString("phone"),
		"email":      data.GetString("email"),
		"blocked":    data.GetBool("blocked"),
		"is_sysuser": data.GetBool("is_sysuser"),
	}
	r.Session.Set(r.Session.Id(), sessionData)
	return sessionData, nil
}

// CheckUsernameUnique 用户名唯一性校验
func (u *userLogic) CheckUsernameUnique(username string) error {
	arg := &dao.ExistsAndCountFuncArg{
		Where:     "username",
		WhereArgs: username,
	}
	exists, err := dao.Exists("user", arg)
	if err != nil {
		return err
	}
	if exists {
		return gerror.WrapCodef(errcode.ExistsUserName, err, errcode.ExistsUserNameMsg, username)
	}
	return nil
}

// CheckEmailUnique 邮箱唯一性校验
func (u *userLogic) CheckEmailUnique(email string) error {
	arg := &dao.ExistsAndCountFuncArg{
		Where:     "email",
		WhereArgs: email,
	}
	exists, err := dao.Exists("user", arg)
	if err != nil {
		return err
	}
	if exists {
		return gerror.WrapCodef(errcode.ExistsUserEmail, err, errcode.ExistsUserEmailMsg, email)
	}
	return nil
}

// CheckPhoneUnique 手机号唯一性校验
func (u *userLogic) CheckPhoneUnique(phone string) error {
	arg := &dao.ExistsAndCountFuncArg{
		Where:     "phone",
		WhereArgs: phone,
	}
	exists, err := dao.Exists("user", arg)
	if err != nil {
		return err
	}
	if exists {
		return gerror.WrapCodef(errcode.ExistsUserPhone, err, errcode.ExistsUserPhoneMsg, phone)
	}
	return nil
}
