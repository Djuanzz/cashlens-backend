package router

import (
	"github.com/Djuanzz/cashlens-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()

	health := r.Group("/health")

	health.GET("/", handler.CheckHealth)
}
