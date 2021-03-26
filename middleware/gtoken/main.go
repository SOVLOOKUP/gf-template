package gtoken

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
)

var cfg = g.Cfg()

func getPaths(path string) g.SliceStr {
	dataSlice := cfg.GetArray(path)
	authPaths := make(g.SliceStr, len(dataSlice))
	for i, path := range dataSlice {
		authPaths[i] = path.(string)
	}
	return authPaths
}

//启动gtoken
var Tokenizer = &gtoken.GfToken{
	LoginPath:        cfg.GetString("tokenizer.loginPath"),
	LoginBeforeFunc:  Login,
	LogoutPath:       cfg.GetString("tokenizer.logoutPath"),
	AuthPaths:        getPaths("tokenizer.authPaths"),
	AuthExcludePaths: getPaths("tokenizer.authExcludePaths"),
	GlobalMiddleware: cfg.GetBool("tokenizer.globalMiddleware"),
	CacheMode:        cfg.GetInt8("tokenizer.cacheMode"),
}
