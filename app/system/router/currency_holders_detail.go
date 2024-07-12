// ==========================================================================
// GFast自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/router/currency_holders_detail.go
// 生成人：gfast
// ==========================================================================

package router

import (
	"gfast/app/system/api"
	"gfast/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//加载路由
func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Group("/system", func(group *ghttp.RouterGroup) {
			group.Group("/currencyHoldersDetail", func(group *ghttp.RouterGroup) {
				//gToken拦截器
				api.GfToken.AuthMiddleware(group)
				//context拦截器
				group.Middleware(middleware.Ctx, middleware.Auth)
				group.GET("list", api.CurrencyHoldersDetail.List)
				group.GET("get", api.CurrencyHoldersDetail.Get)
				group.POST("add", api.CurrencyHoldersDetail.Add)
				group.PUT("edit", api.CurrencyHoldersDetail.Edit)
				group.DELETE("delete", api.CurrencyHoldersDetail.Delete)
			})
		})
	})
}
