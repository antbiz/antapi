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

var codeMsg = map[int]string{
	ServerError:           "Internal Server Error",
	ServiceUnavailable:    "Service Unavailable",
	TooManyRequests:       "Too Many Requests",
	CallHTTPError:         "Call HTTP Error",
	IllegalRequest:        "Illegal Request",
	AuthorizationError:    "Authorization Error",
	IPDenied:              "IP Denied",
	PermissionDenied:      "Permission Denied",
	IPRequestsLimit:       "IP Requests Out Of Rate Limit",
	UserRequestsLimit:     "User Requests Out Of Rate Limit",
	APINotFound:           "API Not Found",
	RequestMethodError:    "Request Method Error",
	ParameterBindError:    "Parameter Bind Error",
	MissRequiredParameter: "Miss Required Parameter",
	ResubmitError:         "Resubmit Error",
	ResubmitWarning:       "Do Not Resubmit",

	IllegalUsername:         "Illegal Username",
	UsernameOrPasswordError: "Username Or Password Error",
}

// ErrMsg .
func ErrMsg(code int) string {
	return codeMsg[code]
}
