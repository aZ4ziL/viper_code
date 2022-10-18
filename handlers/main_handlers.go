package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Home this is main handler
func Home(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get("user")

	ctx.HTML(http.StatusOK, "index", gin.H{
		"user": user,
	})
}
