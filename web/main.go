package main

import (
	"web/database"
	"web/handlers"
	"web/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDataBase()
	database.CreateSCP(models.SCP{Name: "bebebe", DescryptionPath: "SCP-1337.txt", ImagePath: "SCP-1337.jpg", IsSecret: false})
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/resources", "./resources")
	router.GET("/", handlers.GetMainPage)
	router.GET("/:object", handlers.GetObject)
	router.GET("/validate", handlers.Validate)
	// router.POST("/register")
	router.GET("/create_scp", handlers.GetCreateSCP)
	router.POST("/create_scp", handlers.PostCreateSCP)
	router.Run(":8000")
}
