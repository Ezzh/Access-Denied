package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"web/database"
	"web/models"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	// host := c.Request.Host
	data := map[string]interface{}{
		"access": false,
	}
	c.JSON(200, data)
}

func access_verification(c *gin.Context, object models.SCP) bool {
	result := map[string]interface{}{"access": false}
	response, err := http.Get("http://" + c.Request.Host + "/validate")
	if err != nil {
		return false
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&result)
	if access, ok := result["access"].(bool); ok {
		return access
	} else {
		return false
	}
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// for _, u := range users {
	// 	if u.Username == user.Username {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким именем уже существует"})
	// 		return
	// 	}
	// }
}
func GetMainPage(c *gin.Context) {
	objects := database.GetAll()
	c.HTML(200, "main_page.html", gin.H{"objects": objects})
}

func GetObject(c *gin.Context) {
	object := database.GetByName(c.Param("object"))
	if object.IsSecret {
		if !access_verification(c, object) {
			c.String(200, "Access denied!!!")
			return
		}
	}
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

	c.String(200, "Данные успешно сохранены")
}

func GetCreateSCP(c *gin.Context) {
	c.HTML(200, "create_scp.html", gin.H{})
}
