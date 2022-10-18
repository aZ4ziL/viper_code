package main

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/aZ4ziL/viper_code/models"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(time.Time{})
	gob.Register(map[string]interface{}{})
}

func main() {
	r := gin.Default()

	// Use Session

	store := gormsessions.NewStore(models.GetDB(), true, []byte("secret"))
	r.Use(sessions.Sessions("ginSessionID", store))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title": "Hello World",
		})
	})

	r.Run(":8000")
}
