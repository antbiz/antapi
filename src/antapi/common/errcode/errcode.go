package errcode

const (
	// 系统级错误代码
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

	// 服务级错误代码 - 用户
	IllegalUsername             = 20001
	IncorrectUsernameOrPassword = 20002
)

const (
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
	IllegalUsernameMsg             = "非法用户名 (%s)"
	IncorrectUsernameOrPasswordMsg = "账号或密码错误"
)
