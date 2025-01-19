package reservation

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
)

type GetUserReservationsUseCase struct {
	ReservationRepo ports.ReservationRepository
}

func NewGetUserReservationsUseCase(repo ports.ReservationRepository) *GetUserReservationsUseCase {
	return &GetUserReservationsUseCase{ReservationRepo: repo}
}

func (uc *GetUserReservationsUseCase) Execute(id string) ([]entity.Reservation, error) {
	reservations, err := uc.ReservationRepo.GetByUser(id)
	if err != nil {
		return nil, err
	}

	if len(reservations) == 0 {
		return []entity.Reservation{}, nil
	}
	
	return reservations, nil
}