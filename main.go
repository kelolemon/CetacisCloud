package main

import (
	"./src"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	// router
	api := app.Party("/api", crs).AllowMethods(iris.MethodOptions)

	api.Post("/upload", src.PostUpload)

	api.Post("/download/{filename}", src.PostDownload)

	api.Post("/login", src.PostLogin)

	api.Post("/logout", src.PostLogout)

	api.Get("/check", src.IsLog)
	// run the server
	_ = app.Run(iris.Addr(":8080"))
}