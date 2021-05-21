package router

import (
	middlewares "github.com/BeanWei/apikit/middleware"
	"github.com/antbiz/antapi/internal/app/api"
	"github.com/antbiz/antapi/internal/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Use(middlewares.CORS, middleware.Ctx, middleware.ErrorHandler)

	// 通用的biz增删改查
	// TODO: 后台可配置鉴权
	s.Group("/api/biz", func(group *ghttp.RouterGroup) {
		// 查询 {collection} 列表
		group.GET("/{collection}", api.Biz.List)
		// 获取 {collection} 详情
		group.GET("/{collection}/{id}", api.Biz.Get)
		// 添加 {collection}
		group.POST("/{collection}", api.Biz.Create)
		// 修改 {collection}
		group.PUT("/{collection}/{id}", api.Biz.Update)
		// 删除 {collection}
		group.DELETE("/{collection}/{id}", api.Biz.Delete)
	})

	// 文件上传
	s.Group("/api/file", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		// 上传
		group.POST("/upload", api.File.Upload)
	})
	// 文件预览或下载
	s.BindHandler("GET:/upload/{year}/{month}/{day}/{filename}", api.File.Preview)
}