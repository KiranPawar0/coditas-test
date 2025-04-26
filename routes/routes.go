package routes

import (
	"github.com/KiranPawar0/coditas-test/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/user", user.CreateUser)
}
