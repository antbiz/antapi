package router

import (
	"antapi/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	// 通用的biz增删改查
	// TODO: 后台可配置鉴权
	s.Group("/api/business", func(group *ghttp.RouterGroup) {
		// 查询 {collection} 列表
		group.GET("/{collection}", api.Biz.GetList)
		// 获取 {collection} 详情
		group.GET("/{collection}/{id}", api.Biz.Get)
		// 添加 {collection}
		group.POST("/{collection}", api.Biz.Create)
		// 修改 {collection}
		group.PATCH("/{collection}/{id}", api.Biz.Update)
		// 删除 {collection}
		group.DELETE("/{collection}/{id}", api.Biz.Delete)
	})

	// 用户登录
	// TODO: 集成三方登录
	s.Group("/api/signin", func(group *ghttp.RouterGroup) {
		// 账号登录
		group.POST("/byuser", api.SignIn.SignInByUser)
	})

	// 用户注册
	s.Group("/api/signup", func(group *ghttp.RouterGroup) {
		// 邮箱注册
		group.POST("/email", api.SignUp.SignUpWithEmail)
		// 手机号注册
		group.POST("/phone", api.SignUp.SignUpWithPhone)
	})

	// 用户个人相关的接口，需要做auth鉴权
	s.Group("/api/user", func(group *ghttp.RouterGroup) {
		// 退出登录
		group.ALL("/signout", api.User.SignOut)
		// 个人信息
		group.GET("/me", api.User.MyProfile)
	})
}
