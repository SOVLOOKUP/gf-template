package interceptor

import (
	"gf-app/utils/resp"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

//请求拦截器
func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 404
	if r.Response.Status == http.StatusNotFound {
		r.Response.ClearBuffer()
		r.Response.WriteJsonExit(resp.NotFound("NoDataFound", ""))
	}
}
