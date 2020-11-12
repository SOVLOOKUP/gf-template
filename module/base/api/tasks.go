package api

import (
"github.com/gogf/gf/net/ghttp"
)

// Hello is a demonstration route handler for output "Hello World!".
func Tasks(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
