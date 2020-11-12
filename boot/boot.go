package boot

import (
	"gf-app/middleware"
	_ "gf-app/packed"
	"gf-app/utils/gtoken"
	"github.com/gogf/gf/frame/g"
)

func init() {
	InitConfig()
	InitModules()
}

func InitConfig()  {
	s := g.Server()

	//response拦截器
	s.Use(middleware.MiddlewareErrorHandler)

	//开启Gtoken鉴权
	gtoken.Tokenizer.Start()

	//admin平滑重启
	s.SetErrorLogEnabled(true)
	s.EnableAdmin()
}


