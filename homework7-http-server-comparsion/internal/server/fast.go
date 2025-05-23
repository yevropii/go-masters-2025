package server

import "github.com/valyala/fasthttp"

var fastSrv *fasthttp.Server

func StartFast(addr string) error {
	fastSrv = &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			ctx.WriteString("hello, world\n")
		},
	}
	return fastSrv.ListenAndServe(addr)
}
