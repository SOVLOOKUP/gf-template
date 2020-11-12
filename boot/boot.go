package boot

import (
	"gf-app/middleware"
	_ "gf-app/packed"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
)

func init() {
	InitConfig()
	InitModules()
}

func InitConfig()  {
	s := g.Server()
	s.Use(middleware.MiddlewareErrorHandler)

	//启动gtoken
	tokenizer := &gtoken.GfToken{
		LoginPath:       "/login",
		//LoginBeforeFunc: loginFunc,
		LogoutPath:      "/user/logout",
		AuthPaths:        g.SliceStr{"/user", "/system"}, // 这里是按照前缀拦截，拦截/user /user/list /user/add ...
		GlobalMiddleware: false,                           // 开启全局拦截，默认关闭
	}
	tokenizer.Start()

	s.SetErrorLogEnabled(true)
	s.EnableAdmin()
}


