package router

import (
	"antapi/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

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

	// TODO: 集成三方登录
	s.Group("/api/signin", func(group *ghttp.RouterGroup) {
		group.POST("/byuser", api.SignIn.SignInByUser)
	})
}
