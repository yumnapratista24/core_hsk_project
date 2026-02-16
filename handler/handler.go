package handler

import (
	"core_hsk_project/errors"
	"core_hsk_project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.ServiceInterface
}

func NewHandler(service services.ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

// Ping returns a simple pong response
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"success": true,
	})
}

// GetWordsByHskSourceID retrieves all words for a given hsk_source_id
func (h *Handler) GetWordsByHskSourceID(c *gin.Context) {
	hskSourceID, err := strconv.Atoi(c.Param("hsk_source_id"))
	if err != nil {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "Invalid hsk_source_id",
		})
		return
	}

	response, err := h.service.GetWordsByHskSourceID(hskSourceID)
	if len(response.List) == 0 {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "hsk_source_id not exists",
		})
		return
	}
	if err != nil {
		HandleError(c, errors.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve words",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
