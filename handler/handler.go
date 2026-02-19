package handler

import (
	"core_hsk_project/ai"
	"core_hsk_project/errors"
	"core_hsk_project/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   services.ServiceInterface
	aiService ai.ServiceInterface
}

func NewHandler(service services.ServiceInterface, aiSvc ai.ServiceInterface) *Handler {
	return &Handler{
		service:   service,
		aiService: aiSvc,
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

func (h *Handler) GenerateDialogueFromAI(c *gin.Context) {
	hskSourceID, err := strconv.Atoi(c.Param("hsk_source_id"))
	if err != nil {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "Invalid hsk_source_id",
		})
		return
	}

	complexity, err := strconv.Atoi(c.Query("complexity"))
	if err != nil {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "complexity is not valid",
		})
		return
	}

	// Validate text complexity range
	if complexity < 1 || complexity > 3 {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "complexity must be between 1 and 3",
		})
		return
	}

	// Validate complexity for HSK level
	if hskSourceID == 1 && complexity == 3 {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "HSK level 1 only supports complexity 1 and 2",
		})
		return
	}

	words, previousLevelWords, err := h.service.GetWordsWithPreviousLevel(hskSourceID)
	var stringifiedWords strings.Builder
	for _, word := range words {
		stringifiedWords.WriteString(word.Hanzi + "-")
	}

	var stringifiedPrevWords strings.Builder
	for _, word := range previousLevelWords {
		stringifiedPrevWords.WriteString(word.Hanzi + "-")
	}

	result, err := h.aiService.GenerateDialogueFromAI(ai.GenerateDialogueFromAIRequest{
		HSKLevel:           hskSourceID,
		StringifiedWords:   stringifiedWords.String(),
		PreviousLevelWords: stringifiedPrevWords.String(),
		TextComplexity:     complexity,
	})
	if err != nil {
		HandleError(c, errors.CustomError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Errorf("failed to generate dialogue from AI: %v", err).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

func (h *Handler) GenerateGradedTextFromAI(c *gin.Context) {
	// Parse hsk_source_id
	hskSourceID, err := strconv.Atoi(c.Param("hsk_source_id"))
	if err != nil {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "Invalid hsk_source_id",
		})
		return
	}

	// Parse complexity
	complexity, err := strconv.Atoi(c.Query("complexity"))
	if err != nil {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "complexity is not valid",
		})
		return
	}

	// Validate complexity range
	if complexity < 1 || complexity > 3 {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "complexity must be between 1 and 3",
		})
		return
	}

	topic := ai.GetRandomTopic(hskSourceID)

	// HSK level 1 restriction (same as dialogue)
	if hskSourceID == 1 && complexity == 3 {
		HandleError(c, errors.CustomError{
			Code:    http.StatusBadRequest,
			Message: "HSK level 1 only supports complexity 1 and 2",
		})
		return
	}

	words, previousLevelWords, err := h.service.GetWordsWithPreviousLevel(hskSourceID)
	var wordItems []ai.WordItem
	for _, word := range words {
		wordItems = append(wordItems, ai.WordItem{
			Hanzi:    word.Hanzi,
			Pinyin:   word.Pinyin,
			English:  word.GetEnglish(),
			HSKLevel: hskSourceID,
		})
	}

	var prevWordItems []ai.WordItem
	for _, word := range previousLevelWords {
		wordItems = append(wordItems, ai.WordItem{
			Hanzi:    word.Hanzi,
			Pinyin:   word.Pinyin,
			English:  word.GetEnglish(),
			HSKLevel: word.HSKSource.Level,
		})
	}

	result, err := h.aiService.GenerateGradedTextFromAI(ai.GenerateGradedTextFromAIRequest{
		HSKLevel:       hskSourceID,
		Words:          wordItems,
		PrevLevelWords: prevWordItems,
		TextComplexity: complexity,
		Topic:          topic,
	})
	if err != nil {
		HandleError(c, errors.CustomError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Errorf("failed to generate graded text from AI: %v", err).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}
