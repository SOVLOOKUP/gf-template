package boot

import (
	"gf-app/middleware/gtoken"
	"gf-app/middleware/interceptor"

	"github.com/gogf/gf/frame/g"
)

func InitConfig() {
	s := g.Server()

	//response拦截器
	s.Use(interceptor.MiddlewareErrorHandler)

	//开启Gtoken鉴权
	gtoken.Tokenizer.Start()

	// error日志
	s.SetErrorLogEnabled(true)

	//admin平滑重启
	g.SetServerGraceful(true)
	s.EnableAdmin()
}
