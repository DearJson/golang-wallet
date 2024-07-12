/*
* @desc:xxxx功能描述
* @company:NULL
* @Author: Json   wbjson@gmail.com
* @Date:   2021/3/18 16:42
 */

package middleware

import "github.com/gogf/gf/net/ghttp"

// CORS 跨域处理中间件
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
