package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ispahbod/golang-clean-web-api/src/api/routers"
	"github.com/ispahbod/golang-clean-web-api/src/config"
)

func InitServer() {
	cfg := config.GetConfig()
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	v1 := router.Group("/api/v1/")
	{
		health := v1.Group("/health")

		routers.Health(health)
	}
	router.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}
