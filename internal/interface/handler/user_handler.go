package handler

import (
	"cowork_system/internal/application/usecase/user"
	"cowork_system/internal/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	CreateUserUseCase *user.CreateUserUseCase
	GetUsersUseCase *user.GetUsersUseCase
	GetUserUseCase *user.GetUserUseCase
	UpdateUserUseCase *user.UpdateUserUseCase
	DeleteUserUseCase *user.DeleteUserUseCase
	ChangeRoleUseCase *user.ChangeRoleUseCase
}

func NewUserHandler(
	createUserUseCase *user.CreateUserUseCase,
	getUsersUseCase *user.GetUsersUseCase,
	getUserUseCase *user.GetUserUseCase,
	updateUserUseCase *user.UpdateUserUseCase,
	deleteUserUseCase *user.DeleteUserUseCase,
	changeRoleUseCase *user.ChangeRoleUseCase,
) *UserHandler {
	return &UserHandler{
		CreateUserUseCase: createUserUseCase,
		GetUsersUseCase:    getUsersUseCase,
		GetUserUseCase:     getUserUseCase,
		UpdateUserUseCase:  updateUserUseCase,
		DeleteUserUseCase:  deleteUserUseCase,
		ChangeRoleUseCase: changeRoleUseCase,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.GetUsersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.GetUserUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.CreateUserUseCase.Execute(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user entity.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.UpdateUserUseCase.Execute(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	err := h.DeleteUserUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (h *UserHandler) ChangeRole(c *gin.Context) {
	var body struct {
		AdminID string `json:"admin_id"`
		TargetID string `json:"target_id"`
		NewRole string `json:"new_role"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.AdminID == "" || body.TargetID == "" || body.NewRole == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "all fields are required"})
		return
	}

	newRole := entity.Role(body.NewRole)

	if err := h.ChangeRoleUseCase.Execute(body.AdminID,body.TargetID,newRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "role updated successfully"})
}