package handler

import (
	"cowork_system/internal/application/usecase/space"
	"cowork_system/internal/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SpaceHandler struct {
    CreateSpaceUseCase *space.CreateSpaceUseCase
    ListSpacesUseCase  *space.ListSpacesUseCase
}

func NewSpaceHandler(createSpaceUseCase *space.CreateSpaceUseCase, listSpacesUseCase *space.ListSpacesUseCase) *SpaceHandler {
    return &SpaceHandler{
        CreateSpaceUseCase: createSpaceUseCase,
        ListSpacesUseCase:  listSpacesUseCase,
    }
}

func (h *SpaceHandler) GetSpaces(c *gin.Context) {
    spaces, err := h.ListSpacesUseCase.Execute()
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
    createdSpace, err := h.CreateSpaceUseCase.Execute(space)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, createdSpace)
}
