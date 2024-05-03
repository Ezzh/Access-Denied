package main

import (
	"web/database"
	"web/handlers"
	"web/logic"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	store := cookie.NewStore([]byte("secret"))
	database.InitDataBase()
	router := gin.Default()
	router.Use(sessions.Sessions("mysession", store))
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/resources", "./resources")
	router.GET("/", handlers.GetMainPage)
	router.GET("/:object", logic.AuthRequired, handlers.GetObject)
	router.GET("/validate", handlers.Validate)
	router.GET("/register", handlers.GetRegister)
	router.POST("/register", handlers.PostRegister)
	router.GET("/login", handlers.GetLogin)
	router.POST("/login", handlers.PostLogin)
	router.GET("/create_scp", handlers.GetCreateSCP)
	router.POST("/create_scp", handlers.PostCreateSCP)
	router.Run(":8000")
}
