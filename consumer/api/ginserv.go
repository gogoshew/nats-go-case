package api

import (
	"L0_Case/consumer/inner/repo"
	"L0_Case/consumer/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type Handler struct {
//	Db *repo.Database
//}
//
//func (h *Handler) orderPage() *gin.Engine {
//	router := gin.Default()
//	router.GET("/", func(c *gin.Context) {
//		h.Db.GetRowById(4)
//	})
//	return router
//}

type Handler struct {
	Repo   *repo.Repository
	Orders *models.Order
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/orders/:id", h.getById)
	return router
}

func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")
	OrderId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(c.Writer, "err to parce id param %s", err)
	}

	order, err := h.Repo.Cache.GetById(uint(OrderId))
	if err != nil {
		fmt.Fprintf(c.Writer, "no searched key %s", err)
	}

	c.JSON(http.StatusOK, order)
}
