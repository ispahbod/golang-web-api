package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(context *gin.Context) {
	context.JSON(http.StatusOK, "Working")
	return
}
func (h *HealthHandler) HealthPost(context *gin.Context) {
	context.JSON(http.StatusOK, "Working Post Method")
	return
}

func (h *HealthHandler) HealthPostID(context *gin.Context) {
	id := context.Params.ByName("id")
	context.JSON(http.StatusOK, fmt.Sprintf("Working Post Method id: %s", id))
	return
}
