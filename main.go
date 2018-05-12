package main

import (
	"github.com/kataras/iris"
	Route "github.com/zhuyanxi/goblockdemo/routes"
)

func main() {
	//bc := block.NewBlockchain()
	// bc.AddBlock("Send 1 btc to Alice")
	// bc.AddBlock("Send 1.1 btc to Bob")
	// fmt.Println("hello")
	// for _, block := range bc.Blocks {
	// 	fmt.Printf("Prev hash:%x\n", block.PrevHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("Timestamp: %x\n", block.Timestamp)
	// 	fmt.Println()
	// }

	app := iris.New()

	viewTmpl := iris.HTML("./web/views", ".html")
	viewTmpl.Reload(true)
	viewTmpl.Layout("layouts/layout.html")
	app.RegisterView(viewTmpl)
	app.StaticWeb("/", "./web/content")

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMsg := ctx.Values().GetString("error")
		if errMsg != "" {
			ctx.Writef("Internal server error: %s", errMsg)
			return
		}
		ctx.Writef("(Unexpected) internal server error: %s", errMsg)
	})

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	app.Get("/", Route.Index)
	app.Get("/hash", Route.Hash)
	app.Get("/block", Route.Block)

	apiRoute := app.Party("/api", logThisMiddleware)
	{
		apiRoute.Post("/SHA256", Route.SHA256)
		apiRoute.Post("/MineBlock", Route.MineBlock)
		apiRoute.Post("/AddUser", AddUser)
	}

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path:%s|IP:%s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

type User struct {
	Id   int
	Name string
}

func AddUser(ctx iris.Context) {
	var user User
	ctx.ReadJSON(&user)
	ctx.Writef("%d %s", user.Id, user.Name)
}
