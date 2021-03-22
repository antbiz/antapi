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
	ResubmitWarning       = 10016

	// 服务级错误代码 - 用户
	IllegalUsername         = 20001
	UsernameOrPasswordError = 20002
)

const (
	ServerErrorMsg           = "Internal Server Error"
	ServiceUnavailableMsg    = "Service Unavailable"
	TooManyRequestsMsg       = "Too Many Requests"
	CallHTTPErrorMsg         = "Call HTTP Error"
	IllegalRequestMsg        = "Illegal Request"
	AuthorizationErrorMsg    = "Authorization Error"
	IPDeniedMsg              = "IP Denied"
	PermissionDeniedMsg      = "Permission Denied"
	IPRequestsLimitMsg       = "IP Requests Out Of Rate Limit"
	UserRequestsLimitMsg     = "User Requests Out Of Rate Limit"
	APINotFoundMsg           = "API Not Found"
	RequestMethodErrorMsg    = "Request Method Error"
	ParameterBindErrorMsg    = "Parameter Bind Error"
	MissRequiredParameterMsg = "Miss Required Parameter"
	ResubmitErrorMsg         = "Resubmit Error"
	ResubmitWarningMsg       = "Do Not Resubmit"

	IllegalUsernameMsg         = "Illegal Username"
	UsernameOrPasswordErrorMsg = "Username Or Password Error"
)
