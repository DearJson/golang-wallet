package middleware

import (
	"gfast/app/common/service"
	"gfast/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

// CheckIp 检查请求IP是否符合要求
func CheckIp(r *ghttp.Request) {
	requestIp := library.GetClientIp(r)
	allowIps, err := service.SysConfig.GetConfigByKey("sys.whitelist")
	if err != nil || allowIps == nil {
		library.FailJson(true, r, "受限制的ip不能访问")
	}
	allowIpssdss := strings.Split(allowIps.ConfigValue, ",")
	if library.ElementIsInSlice(requestIp, allowIpssdss) == false {
		g.Log().File("web-request-{Y-m-d}.log").Printf("异常请求,ip【%v】,域名【%#v】,参数【%v】 \n", library.GetClientIp(r), r.Request.URL, r.GetFormMap())
		library.FailJson(true, r, "没有访问权限")
	}
	g.Log().File("web-request-{Y-m-d}.log").Printf("接收请求,ip【%v】,域名【%v】,参数【%v】 \n", library.GetClientIp(r), r.Request.URL, r.GetFormMap())
	r.Middleware.Next()
}
