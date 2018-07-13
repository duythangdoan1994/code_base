package viewlist

import "github.com/kataras/iris"

func handler(ctx iris.Context){
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}