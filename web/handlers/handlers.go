package handlers

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"web/database"
	"web/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	// host := c.Request.Host
	data := map[string]interface{}{
		"access": false,
	}
	c.JSON(200, data)
}

func PostRegister(c *gin.Context) {
	var user models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	session := sessions.Default(c)

	for _, u := range database.GetAllUser() {
		if u.Username == user.Username {
			c.HTML(http.StatusOK, "register.html", gin.H{"error": "Такой пользователь уже существует!"})
			return
		}
	}
	database.CreateUser(user)
	session.Set("username", user.Username)
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func GetRegister(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}

func PostLogin(c *gin.Context) {
	var user models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	session := sessions.Default(c)

	for _, u := range database.GetAllUser() {
		if u.Username == user.Username && u.Password == user.Password {
			session.Set("username", user.Username)
			session.Save()
			c.Redirect(http.StatusFound, "/")
			return
		}
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"error": "Неправильное имя пользователя или пароль"})

}

func GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func GetMainPage(c *gin.Context) {
	objects := database.GetAll()
	c.HTML(200, "main_page.html", gin.H{"objects": objects})
}

func GetObject(c *gin.Context) {
	object := database.GetByName(c.Param("object"))
	// if object.IsSecret {
	// 	if !access_verification(c, object) {
	// 		c.String(200, "Access denied!!!")
	// 		return
	// 	}
	// }
	if object == (models.SCP{}) {
		c.String(http.StatusNotFound, "Object not found")
		return
	}
	imageData, err := os.ReadFile("./secret-data/images/" + object.ImagePath)
	if err != nil {
		log.Print(err.Error())
		c.String(http.StatusInternalServerError, "Error reading image")
		return
	}

	description, err := os.ReadFile("./secret-data/description/" + object.DescryptionPath)
	if err != nil {
		log.Print(err.Error())
		c.String(http.StatusInternalServerError, "Error reading descrptoin")
		return
	}
	log.Print(string(description))
	encodedImage := base64.StdEncoding.EncodeToString(imageData)
	c.HTML(200, "object.html", gin.H{"name": object.Name, "imagedata": encodedImage, "description": strings.Split(string(description), "\n")})
}

func PostCreateSCP(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	image, err := c.FormFile("image")

	if err != nil {
		c.String(400, "Ошибка при загрузке файла")
		return
	}
	ImageFilePath := filepath.Join("./secret-data/images", filepath.Base(image.Filename))
	file, err := os.Create(ImageFilePath)
	if err != nil {
		c.String(500, "Ошибка при создании файла для изображения")
		return
	}
	defer file.Close()

	src, err := image.Open()
	if err != nil {
		c.String(500, "Ошибка при открытии файла: %v", err)
		return
	}
	defer src.Close()

	if _, err := io.Copy(file, src); err != nil {
		c.String(500, "Ошибка при сохранении изображения")
		return
	}

	txtFilePath := filepath.Join("./secret-data/description", filepath.Base(name)+".txt")
	txtFile, err := os.OpenFile(txtFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.String(500, "Ошибка при открытии файла")
		return
	}
	defer txtFile.Close()

	if _, err := txtFile.WriteString(description); err != nil {
		c.String(500, "Ошибка при записи в файл")
		return
	}

	database.CreateSCP(models.SCP{DescryptionPath: name + ".txt", ImagePath: image.Filename, Name: name, IsSecret: false})

	c.String(200, "Данные успешно сохранены")
}

func GetCreateSCP(c *gin.Context) {
	c.HTML(200, "create_scp.html", gin.H{})
}
