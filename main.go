package main

import (
	"context"
	"core_hsk_project/ai"
	"core_hsk_project/handler"
	"core_hsk_project/middleware"
	"core_hsk_project/model"
	"core_hsk_project/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-deepseek/deepseek"
)

func main() {
	router := gin.Default()

	dsn := os.Getenv("DBUrl")
	if dsn == "" {
		panic("DBUrl environment variable is empty")
	}

	db := model.NewDatabase(dsn)
	deepseekClient, err := deepseek.NewClient(os.Getenv("DeepseekAPIKey"))
	if err != nil {
		panic("Failed to create deepseek client")
	}
	aiService := ai.NewService(&deepseekClient, context.Background())
	hskModel := model.NewHskModel(db)
	service := services.NewService(hskModel)
	handler := handler.NewHandler(service, aiService)

	setupRoutes(router, handler)

	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, handler *handler.Handler) {
	api := router.Group("/api")
	api.Use(middleware.Authenticate())
	api.GET("/hsk-sources/:hsk_source_id/words", handler.GetWordsByHskSourceID)
	api.GET("/hsk-sources/:hsk_source_id/generate-dialogue", handler.GenerateDialogueFromAI)

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
