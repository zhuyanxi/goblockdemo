package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/zhuyanxi/goblockdemo/block"
	Route "github.com/zhuyanxi/goblockdemo/routes"
)

func main() {
	bc := block.NewBlockchain()
	//bc.AddBlock("Send 1 btc to Alice")
	//bc.AddBlock("Send 1.1 btc to Bob")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev hash:%x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println(block.Hash)
		fmt.Printf("Timestamp: %x\n", block.Timestamp)
		fmt.Printf("Nouce: %d\n", block.Nouce)
		fmt.Println()
	}

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
		apiRoute.Post("/ComputeBlockHash", Route.ComputeBlockHash)
		apiRoute.Post("/AddUser", AddUser)
	}

	app.Run(iris.Addr(":80"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path:%s | IP:%s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

type User struct {
	ID        int
	Name      string
	Timestamp int64
	Data      []byte
	BaseInfo  BaseInfo
}

type BaseInfo struct {
	Address string
	City    string
	Zipcode int
}

func AddUser(ctx iris.Context) {
	var user User

	//ctx.PostValue("baseinfo")

	ctx.ReadJSON(&user)
	ctx.Writef("%d, %s, %x, %d\n", user.ID, user.Name, user.Data, user.Timestamp)
	ctx.Writef("%s, %s, %d", user.BaseInfo.Address, user.BaseInfo.City, user.BaseInfo.Zipcode)
}
