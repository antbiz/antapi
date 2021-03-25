package middleware

import "github.com/gogf/gf/net/ghttp"

// CORS 允许跨域请求中间件
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
