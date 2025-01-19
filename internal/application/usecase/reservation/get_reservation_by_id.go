package reservation

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
)

type GetReservationUseCase struct {
	ReservationRepo ports.ReservationRepository
}

func NewGetReservationUseCase(repo ports.ReservationRepository) *GetReservationUseCase {
	return &GetReservationUseCase{ReservationRepo: repo}
}

func (uc *GetReservationUseCase) Execute(id string) (entity.Reservation, error) {
	reservation, err := uc.ReservationRepo.GetByID(id)
	if err != nil {
		return entity.Reservation{}, err
	}

	if reservation.ID == "" {
		return entity.Reservation{}, errors.New("reservation not found")
	}

	return reservation, nil
}