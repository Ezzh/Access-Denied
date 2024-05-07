package logic

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"web/database"
	"web/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GenerateRandomKey() (string, error) {
	buffer := make([]byte, 32)

	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	key := hex.EncodeToString(buffer)

	return key, nil
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.Next()
}

func AccessVerification(c *gin.Context, object models.SCP) bool {
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

func GetUserFromSession(c *gin.Context) models.User {
	session := sessions.Default(c)
	username := session.Get("username")
	return database.GetUserByName(username.(string))
}

func ExitSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()
}
