package main

import (
	"encoding/gob"
	"time"

	"github.com/aZ4ziL/viper_code/models"
	"github.com/aZ4ziL/viper_code/renderer"
	"github.com/aZ4ziL/viper_code/routers"
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

	r.HTMLRender = renderer.CreateMyRender()

	// Set Static
	r.Static("/static", "./static")

	// Use Session
	store := gormsessions.NewStore(models.GetDB(), true, []byte("secret"))
	r.Use(sessions.Sessions("ginSessionID", store))

	routers.HomeRouter(r)

	r.Run(":8000")
}
