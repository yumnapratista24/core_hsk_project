package main

import (
	"core_hsk_project/handler"
	"core_hsk_project/middleware"
	"core_hsk_project/model"
	"core_hsk_project/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dsn := os.Getenv("DBUrl")
	db := model.NewDatabase(dsn)
	hskModel := model.NewHskModel(db)
	service := services.NewService(hskModel)
	handler := handler.NewHandler(service)

	setupRoutes(router, handler)

	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, handler *handler.Handler) {
	api := router.Group("/api")
	api.Use(middleware.Authenticate())
	api.GET("/hsk-sources/:hsk_source_id/words", handler.GetWordsByHskSourceID)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error": gin.H{
				"code":    http.StatusNotFound,
				"message": "Endpoint not found",
			},
		})
	})
}
