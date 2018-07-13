package user

import (
	"github.com/kataras/iris"
)

func HandleUserUrl(api *iris.Application)  *iris.Application{
	user := api.Party("/user")
	user.Get("/all", handler)
	return api
}

func handler(ctx iris.Context){
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}