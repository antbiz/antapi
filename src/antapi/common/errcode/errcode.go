package errcode

const (
	// 系统级错误代码
	UnknownError          = 10000
	ServerError           = 10001
	ServiceUnavailable    = 10002
	TooManyRequests       = 10003
	CallHTTPError         = 10004
	IllegalRequest        = 10005
	AuthorizationError    = 10006
	IPDenied              = 10007
	PermissionDenied      = 10008
	IPRequestsLimit       = 10009
	UserRequestsLimit     = 10010
	APINotFound           = 10011
	RequestMethodError    = 10012
	ParameterBindError    = 10013
	MissRequiredParameter = 10014
	ResubmitError         = 10015
	JSONError             = 10016
	DuplicateError        = 10017

	// 服务级错误代码 - 用户
	IncorrectUsernameOrPassword = 20001
	IncorrectOldPassword        = 20002
	IllegalUsername             = 20003
	ExistsUserName              = 20004
	ExistsUserEmail             = 20005
	ExistsUserPhone             = 20006
)

const (
	UnknownErrorMsg                = "未知错误"
	ServerErrorMsg                 = "服务器错误"
	ServiceUnavailableMsg          = "服务不可用"
	TooManyRequestsMsg             = "当前请求过多，系统繁忙"
	CallHTTPErrorMsg               = "调用第三方HTTP接口失败"
	IllegalRequestMsg              = "非法请求"
	AuthorizationErrorMsg          = "签名信息错误"
	IPDeniedMsg                    = "IP(%s) 限制不能请求该资源"
	PermissionDeniedMsg            = "权限限制不能请求该资源"
	IPRequestsLimitMsg             = "IP(%s) 请求频次超过上限"
	UserRequestsLimitMsg           = "用户 (%s) 请求频次超过上限"
	APINotFoundMsg                 = "接口不存在"
	RequestMethodErrorMsg          = "请求方法错误"
	ParameterBindErrorMsg          = "参数值非法，需为 (%s)，实际为 (%s)"
	MissRequiredParameterMsg       = "缺失必选参数 (%s)"
	ResubmitErrorMsg               = "请勿重复提交"
	JSONErrorMsg                   = "无效的JSON"
	DuplicateErrorMsg              = "(%s)已存在，请勿重新创建"
	IncorrectUsernameOrPasswordMsg = "账号或密码错误"
	IncorrectOldPasswordMsg        = "旧密码错误"
	IllegalUsernameMsg             = "非法用户名 (%s)"
	ExistsUserNameMsg              = "用户名 (%s) 已被占用"
	ExistsUserEmailMsg             = "邮箱 (%s) 已被占用"
	ExistsUserPhoneMsg             = "手机号 (%s) 已被占用"
)
