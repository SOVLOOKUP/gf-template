package middleware

import (
	"gf-app/utils/resp"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status == http.StatusNotFound {
		r.Response.ClearBuffer()
		r.Response.WriteJsonExit(resp.NotFound("NoDataFound",""))
	}
}