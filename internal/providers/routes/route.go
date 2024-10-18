package routes

import (
	homeRoutes "go-blog/internal/modules/home"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
}