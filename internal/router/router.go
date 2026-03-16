package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	HealthRouter(api)

	return r
}
