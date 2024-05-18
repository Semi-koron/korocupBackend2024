package main

import (
	"github.com/semikoron/korocupbackend/database"
	"github.com/semikoron/korocupbackend/originalmiddleware"
	"github.com/semikoron/korocupbackend/services"
	"github.com/semikoron/korocupbackend/utils/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Open a PostgreSQL database.
	config.LoadEnv()
	database.ConnectDB()
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	// User
	e.GET("/fetch/users", services.FetchUsers)           // すべてのユーザーを取得
	e.GET("/get/user/:user_id", services.GetUserProfile) // ユーザーのプロフィールを取得
	// Post
	e.POST("/create/post", services.NewPost)
	e.GET("/get/posts", services.GetPosts)
	e.GET("/get/post/:postid", services.GetPostDetail)
	r := e.Group("/auth")
	r.Use(originalmiddleware.FirebaseAuthMiddleware)
	r.GET("/test", services.Hello)

	// User
	r.PUT("/update/user", services.UpdateUser) // ユーザーのプロフィールを更新
	r.POST("/login", services.Login)           // ユーザーをログイン
	r.POST("/create/user", services.NewUser)   // ユーザーを作成

	// Post
	r.POST("/create/post", services.NewPost) // 投稿を作成
	e.Logger.Fatal(e.Start(":8080"))
}
