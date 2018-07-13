package main

import (

	"github.com/kataras/iris"
	"github.com/codebase/core"
	"github.com/codebase/user"
	"github.com/kataras/golog"
)

func main() {
	core.InitConfig()
	app := iris.New()
	user.HandleUserUrl(app)
	app.Run(iris.Addr(core.AppConfig.Port))
	golog.Info("Start ting server")
}