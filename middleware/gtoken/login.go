package gtoken

import "github.com/gogf/gf/net/ghttp"

//鉴权登录
func Login(r *ghttp.Request) (string, interface{}) {
	username := r.GetString("username")
	passwd := r.GetString("passwd")

	//仅供测试
	username = "sss"
	userdata := passwd
	// TODO 进行登录校验

	return username, userdata
}
