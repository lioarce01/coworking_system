package reservation

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
)

type GetSpaceReservationsUseCase struct {
	ReservationRepo ports.ReservationRepository
}

func NewGetSpaceReservationsUseCase(repo ports.ReservationRepository) *GetSpaceReservationsUseCase {
	return &GetSpaceReservationsUseCase{ReservationRepo: repo}
}

func (uc *GetSpaceReservationsUseCase) Execute(id string) ([]entity.Reservation, error)  {
	reservations, err := uc.ReservationRepo.GetBySpace(id)
	if err != nil {
		return nil, err
	}

	if len(reservations) == 0 {
		return []entity.Reservation{}, nil
	}
	
	return reservations, nil
}