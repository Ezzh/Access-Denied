package logic

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Необходима аутентификация"})
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
