package routers

import (
	"github.com/aZ4ziL/viper_code/handlers"
	"github.com/gin-gonic/gin"
)

func HomeRouter(r *gin.Engine) {
	r.GET("/", handlers.Home)
}
