package handler

import (
	"cowork_system/internal/application/usecase/reservation"
	"cowork_system/internal/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	CreateReservationUseCase *reservation.CreateReservationUseCase
	GetReservationsUseCase *reservation.GetReservationsUseCase
	GetReservationUseCase *reservation.GetReservationUseCase
	GetSpaceReservationsUseCase *reservation.GetSpaceReservationsUseCase
	GetUserReservationsUseCase *reservation.GetUserReservationsUseCase
	UpdateReservationUseCase *reservation.UpdateReservationUseCase
	DeleteReservationUseCase *reservation.DeleteReservationUseCase
}

func NewReservationHandler(
	createReservationUseCase *reservation.CreateReservationUseCase,
	getReservationsUseCase *reservation.GetReservationsUseCase,
	getReservationUseCase *reservation.GetReservationUseCase,
	getSpaceReservationsUseCase *reservation.GetSpaceReservationsUseCase,
	getUserReservationsUseCase *reservation.GetUserReservationsUseCase,
	updateReservationUseCase *reservation.UpdateReservationUseCase,
	deleteReservationUseCase *reservation.DeleteReservationUseCase,
) *ReservationHandler {
	return &ReservationHandler{
		CreateReservationUseCase: createReservationUseCase,
		GetReservationsUseCase: getReservationsUseCase,
		GetReservationUseCase: getReservationUseCase,
		GetSpaceReservationsUseCase: getSpaceReservationsUseCase,
		GetUserReservationsUseCase: getUserReservationsUseCase,
		UpdateReservationUseCase: updateReservationUseCase,
		DeleteReservationUseCase: deleteReservationUseCase,
	}
}
	
func (h *ReservationHandler) GetReservations(c *gin.Context) {
	reservations, err := h.GetReservationsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) GetReservation(c *gin.Context) {
	id := c.Param("id")
	reservation, err := h.GetReservationUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

func (h *ReservationHandler) GetSpaceReservations(c *gin.Context) {
	spaceId := c.Param("id")
	reservations, err := h.GetSpaceReservationsUseCase.Execute(spaceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) GetUserReservations(c *gin.Context) {
	userID := c.Param("id")
	reservations, err := h.GetUserReservationsUseCase.Execute(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var reservation entity.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdReservation, err := h.CreateReservationUseCase.Execute(
		reservation.SpaceID,
		reservation.UserID,
		reservation.StartTime,
		reservation.EndTime,
		reservation.NumPersons,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdReservation)
}


func (h *ReservationHandler) UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	var reservation entity.Reservation
	
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedReservation, err := h.UpdateReservationUseCase.Execute(id, reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedReservation)
}

func (h *ReservationHandler) DeleteReservation(c *gin.Context) {
    reservationID := c.Param("id")

    err := h.DeleteReservationUseCase.Execute(reservationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "reservation deleted successfully"})
}
