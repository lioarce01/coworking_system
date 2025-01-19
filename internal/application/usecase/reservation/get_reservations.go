package reservation

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
)

type GetReservationsUseCase struct {
	ReservationRepo ports.ReservationRepository
}

func NewGetReservationsUseCase(repo ports.ReservationRepository) *GetReservationsUseCase {
	return &GetReservationsUseCase{ReservationRepo: repo}
}

func (uc *GetReservationsUseCase) Execute() ([]entity.Reservation, error) {
	return uc.ReservationRepo.GetAll()
}