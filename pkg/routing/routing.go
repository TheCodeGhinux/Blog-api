package routing

import (
	"go-blog/internal/providers/routes"

	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	router := GetRouter()

	routes.RegisterRoutes(router)
}
