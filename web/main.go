package main

import (
	"web/database"
	"web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDataBase()
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/resources", "./resources")
	router.GET("/", handlers.GetMainPage)
	router.GET("/:object", handlers.GetObject)
	router.Run(":8000")
}
