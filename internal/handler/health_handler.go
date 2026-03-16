package handler

import (
	"net/http"

	"github.com/Djuanzz/cashlens-backend/internal/repository"
	"github.com/Djuanzz/cashlens-backend/internal/service"
	"github.com/Djuanzz/cashlens-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	service *service.HealthService
}

func NewHealthHandler() *HealthHandler {
	repo := repository.NewHealthRepository()
	service := service.NewHealthService(repo)

	return &HealthHandler{
		service: service,
	}
}

func (h *HealthHandler) CheckHealth(C *gin.Context) {
	result := h.service.CheckHealth()

	utils.SuccessResponse(C, http.StatusOK, result)
}
