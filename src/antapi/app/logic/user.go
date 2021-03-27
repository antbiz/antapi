package logic

import (
	"antapi/app/dao"
	"antapi/app/global"
	"antapi/app/model"
	"antapi/common/errcode"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

var User = &userLogic{
	collectionName: "user",
}

type userLogic struct {
	collectionName string
}

// EncryptPwd 加密账号密码
func (userLogic) EncryptPwd(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}

// GetUserByLogin 根据 用户名/手机号/邮箱 + 密码 查询用户信息
func (u *userLogic) GetUserByLogin(login, pwd string) (data *gjson.Json, err error) {
	arg := &dao.GetFuncArg{
		Where:                 "username=? or phone=? or email=?",
		WhereArgs:             g.Slice{login, login, login},
		IncludeHiddenField:    true,
		IncludePrivateField:   true,
		IgnorePermissionCheck: true,
	}
	if data, err = dao.Get(u.collectionName, arg); err != nil {
		return
	} else if data == nil {
		return nil, gerror.NewCode(errcode.IncorrectUsernameOrPassword, errcode.IncorrectUsernameOrPasswordMsg)
	}
	username := data.GetString("username")
	password := data.GetString("password")
	if u.EncryptPwd(username, pwd) != password || data == nil {
		return nil, gerror.NewCode(errcode.IncorrectUsernameOrPassword, errcode.IncorrectUsernameOrPasswordMsg)
	}

	return data, nil
}

// CheckUserFieldUnique 用户唯一性校验
func (u *userLogic) CheckUserUnique(fieldname, value string) error {
	userSchema := global.GetSchema(u.collectionName)
	if !garray.NewStrArrayFrom(userSchema.GetFieldNames(true, true)).Contains(fieldname) {
		return gerror.NewCode(errcode.IllegalRequest, errcode.IllegalRequestMsg)
	}
	arg := &dao.ExistsAndCountFuncArg{
		Where:     fieldname,
		WhereArgs: value,
	}
	exists, err := dao.Exists(u.collectionName, arg)
	if err != nil {
		return err
	}
	if exists {
		switch fieldname {
		case "username":
			return gerror.WrapCodef(errcode.ExistsUserName, err, errcode.ExistsUserNameMsg, value)
		case "email":
			return gerror.WrapCodef(errcode.ExistsUserEmail, err, errcode.ExistsUserEmailMsg, value)
		case "phone":
			return gerror.WrapCodef(errcode.ExistsUserPhone, err, errcode.ExistsUserPhoneMsg, value)
		default:
			return gerror.WrapCodef(errcode.DuplicateError, err, errcode.DuplicateErrorMsg, "该用户")
		}
	}

	return nil
}

// SignIn 用户登录
func (u *userLogic) SignIn(req *model.UserSignInReq) (map[string]interface{}, error) {
	data, err := u.GetUserByLogin(req.Login, req.Pwd)
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
	return sessionData, nil
}

// SignUpWithEmail 用户邮箱注册
func (u *userLogic) SignUpWithEmail(req *model.UserSignUpWithEmailReq) error {
	if err := u.CheckUserUnique("username", req.Username); err != nil {
		return err
	}
	if err := u.CheckUserUnique("email", req.Email); err != nil {
		return err
	}

	// 创建用户时password字段需要暴露，前面做了数据检查，这里可以忽略
	arg := &dao.InsertFuncArg{
		IgnoreFieldValueCheck: true,
		IncludeHiddenField:    true,
		IncludePrivateField:   true,
		IgnorePermissionCheck: true,
	}
	data := g.Map{
		"username": req.Username,
		"email":    req.Email,
		"password": u.EncryptPwd(req.Username, req.Password),
	}
	if _, err := dao.Insert(u.collectionName, arg, data); err != nil {
		return err
	}

	return nil
}

// SignUpWithPhone 用户手机号注册
func (u *userLogic) SignUpWithPhone(req *model.UserSignUpWithPhoneReq) error {
	if err := u.CheckUserUnique("username", req.Username); err != nil {
		return err
	}
	if err := u.CheckUserUnique("phone", req.Phone); err != nil {
		return err
	}

	// 创建用户时password字段需要暴露，前面做了数据检查，这里可以忽略
	arg := &dao.InsertFuncArg{
		IgnoreFieldValueCheck: true,
		IncludeHiddenField:    true,
		IncludePrivateField:   true,
		IgnorePermissionCheck: true,
	}
	data := g.Map{
		"username": req.Username,
		"phone":    req.Phone,
		"password": u.EncryptPwd(req.Username, req.Password),
	}
	if _, err := dao.Insert(u.collectionName, arg, data); err != nil {
		return err
	}

	return nil
}

// UpdatePassword 修改密码
func (u *userLogic) UpdatePassword(userID string, req *model.UserUpdatePasswordReq) error {
	getArg := &dao.GetFuncArg{
		Where:                 "id=?",
		WhereArgs:             userID,
		IncludeHiddenField:    true,
		IncludePrivateField:   true,
		IgnorePermissionCheck: true,
	}
	data, err := dao.Get(u.collectionName, getArg)
	if err != nil {
		return err
	}

	if u.EncryptPwd(data.GetString("username"), req.OldPassword) != data.GetString("password") {
		return gerror.NewCode(errcode.IncorrectOldPassword, errcode.IncorrectOldPasswordMsg)
	}

	updateArg := &dao.UpdateFuncArg{
		IncludeHiddenField:  true,
		IncludePrivateField: true,
	}
	updateData := g.Map{
		"password": req.Password,
	}
	if err := dao.Update(u.collectionName, updateArg, userID, updateData); err != nil {
		return err
	}

	return nil
}

// GetProfileByID 获取个人信息
func (u *userLogic) GetProfileByID(userID string) (*gjson.Json, error) {
	arg := &dao.GetFuncArg{
		Where:                 "id=?",
		WhereArgs:             userID,
		IgnorePermissionCheck: true,
	}
	return dao.Get(u.collectionName, arg)
}
