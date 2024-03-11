package handlers

import (
	"web/database"

	"github.com/gin-gonic/gin"
)

func DoRequest(c *gin.Context) {
	host := c.Request.Host
	c.String(200, host)
}

func GetMainPage(c *gin.Context) {
	c.HTML(200, "main_page.html", gin.H{})
}

func GetObject(c *gin.Context) {
	be := database.GetByName("SCP-4065")

	c.HTML(200, "object.html", gin.H{"name": be.Name})
}
