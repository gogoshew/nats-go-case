package handler

import (
	"L0_Case/consumer/inner/repository"
	"L0_Case/consumer/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo   *repository.Repository
	Orders chan *models.Order
}

func NewHandler(repository *repository.Repository, orders chan *models.Order) *Handler {
	return &Handler{Repo: repository, Orders: orders}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/", "Consumer/static/index.html")
	router.GET("/orders/:id", h.getById)

	return router
}
