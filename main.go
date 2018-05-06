package main

import (
	"block"
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	bc := block.NewBlockchain()
	bc.AddBlock("Send 1 btc to Alice")
	bc.AddBlock("Send 1.1 btc to Bob")
	fmt.Println("hello")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev hash:%x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %x\n", block.Timestamp)
		fmt.Println()
	}

	app := iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))
	app.StaticWeb("/", "./web/content")

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMsg := ctx.Values().GetString("error")
		if errMsg != "" {
			ctx.Writef("Internal server error: %s", errMsg)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	// Post: scheme://mysubdomain.$domain.com/decode
	app.Subdomain("mysubdomain.").Post("/decode", func(ctx iris.Context) {})
	// Method POST: http://localhost:8080/decode
	app.Post("/decode", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s %s is %d years old and comes from %s", user.Firstname, user.Lastname, user.Age, user.City)
	})

	// http://localhost:8080/encode
	app.Get("/encode", func(ctx iris.Context) {
		doe := User{
			Username:  "Johndoe",
			Firstname: "John",
			Lastname:  "Doe",
			City:      "blabla",
			Age:       25,
		}
		ctx.JSON(doe)
	})

	// http://localhost:8080/profile/anytypeofstring
	app.Get("/profile/{username:string}", profileByUsername)

	usersRoutes := app.Party("/users", logThisMiddleware)
	{
		// /users/42
		usersRoutes.Get("/{id:int min(1)}", getUserByID)
		// /users/create
		usersRoutes.Post("/create", createUser)
	}

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", bc.Blocks)
		ctx.View("index.html")
	})

	// app.Get("/user/{id:long}", func(ctx iris.Context) {
	// 	userID, _ := ctx.Params().GetInt64("id")
	// 	ctx.Writef("User ID:%d", userID)
	// })

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}

type User struct {
	Username  string
	Firstname string
	Lastname  string
	City      string
	Age       int
}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path:%s|IP:%s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

func profileByUsername(ctx iris.Context) {
	username := ctx.Params().Get("username")
	ctx.ViewData("Username", username)
	ctx.View("user/profile.html")
}

func getUserByID(ctx iris.Context) {
	userID := ctx.Params().Get("id")
	user := User{Username: "username" + userID}
	ctx.XML(user)
}
func createUser(ctx iris.Context) {
	var user User
	err := ctx.ReadForm(&user)
	if err != nil {
		ctx.Values().Set("error", "create user, read and parse form failed, "+err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.ViewData("", user)
	ctx.View("user/create_verification.html")
}
