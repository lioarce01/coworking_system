package handler

import (
	"cowork_system/internal/application/usecase"
	"cowork_system/internal/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SpaceHandler struct {
	Usecase *usecase.SpaceUsecase
}

func NewSpaceHandler(usecase *usecase.SpaceUsecase) *SpaceHandler {
	return &SpaceHandler{Usecase: usecase}
}

func (h *SpaceHandler) GetSpaces(c *gin.Context) {
	spaces, err := h.Usecase.GetSpaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, spaces)
}

func (h *SpaceHandler) CreateSpace(c *gin.Context) {
	var space entity.Space
	if err := c.ShouldBindJSON(&space); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdSpace, err := h.Usecase.CreateSpace(space)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdSpace)
}
