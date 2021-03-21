package resp

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gvalid"
)

const (
	SUCCESS = 0
	ERROR   = 1
)

// Resp : normal resp struct
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ListsData : list data resp struct
type ListsData struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

type apiResp struct {
	Resp
	r *ghttp.Request
}

// NewResp : resp build
func NewResp(r *ghttp.Request, code int, msg string, data ...interface{}) *apiResp {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}

	return &apiResp{
		Resp: Resp{
			Code: code,
			Msg:  msg,
			Data: d,
		},
		r: r,
	}
}

// Success : success resp
func Success(r *ghttp.Request, data ...interface{}) *apiResp {
	return NewResp(r, SUCCESS, "", data...)
}

// Error : fail resp
func Error(r *ghttp.Request) *apiResp {
	return NewResp(r, ERROR, "fail")
}

func (apiResp *apiResp) SetCode(code int) *apiResp {
	apiResp.Code = code
	return apiResp
}

func (apiResp *apiResp) SetMsg(msg string) *apiResp {
	apiResp.Msg = msg
	return apiResp
}

func (apiResp *apiResp) SetError(err error) *apiResp {
	switch v := err.(type) {
	case *gvalid.Error:
		apiResp.Msg = v.FirstString()
	default:
		apiResp.Msg = err.Error()
	}
	glog.Skip(1).Line(true).Println(err.Error())
	return apiResp
}

func (apiResp *apiResp) SetData(data interface{}) *apiResp {
	apiResp.Data = data
	return apiResp
}

func (apiResp *apiResp) Json() {
	_ = apiResp.r.Response.WriteJsonExit(apiResp.Resp)
}
