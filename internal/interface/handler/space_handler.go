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
    GetByIDUseCase     *space.GetSpaceUseCase
    UpdateSpaceUseCase *space.UpdateSpaceUseCase
    DeleteSpaceUseCase *space.DeleteSpaceUseCase
}

func NewSpaceHandler(
    createSpaceUseCase *space.CreateSpaceUseCase,
    listSpacesUseCase *space.ListSpacesUseCase,
    getByIDUseCase *space.GetSpaceUseCase,
    updateSpaceUseCase *space.UpdateSpaceUseCase,
    deleteSpaceUseCase *space.DeleteSpaceUseCase,
) *SpaceHandler {
    return &SpaceHandler{
        CreateSpaceUseCase: createSpaceUseCase,
        ListSpacesUseCase:  listSpacesUseCase,
        GetByIDUseCase:     getByIDUseCase,
        UpdateSpaceUseCase: updateSpaceUseCase,
        DeleteSpaceUseCase: deleteSpaceUseCase,
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

func (h *SpaceHandler) GetSpaceByID(c *gin.Context) {
    id := c.Param("id")
    
    space, err := h.GetByIDUseCase.Execute(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, space)
}

func (h *SpaceHandler) UpdateSpace(c *gin.Context) {
    var space entity.Space
    id := c.Param("id")
    
    if err := c.ShouldBindJSON(&space); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updatedSpace, err := h.UpdateSpaceUseCase.Execute(id, &space)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedSpace)
}

func (h *SpaceHandler) DeleteSpace(c *gin.Context) {
    id := c.Param("id") 
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
        return
    }

    err := h.DeleteSpaceUseCase.Execute(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Space deleted successfully"})
}
