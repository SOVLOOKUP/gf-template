package handler

import (
	"gf-app/module/base/model"
	"gf-app/utils/resp"

	"github.com/gogf/gf/net/ghttp"
)

// Hello is a demonstration route handler for output "Hello World!".
func Tasks(r *ghttp.Request) {
	world := &model.Hello{
		Name: "world",
	}

	r.Response.Writeln(
		resp.Succ(world.SayHello()),
	)
}
