package reservation

import (
	"cowork_system/internal/application/ports"
	"errors"
)

type DeleteReservationUseCase struct {
	ReservationRepo ports.ReservationRepository
}

func NewDeleteReservationUseCase(repo ports.ReservationRepository) *DeleteReservationUseCase {
	return &DeleteReservationUseCase{ReservationRepo: repo}
}

func (uc *DeleteReservationUseCase) Execute(id string) error {
	reservation, err := uc.ReservationRepo.GetByID(id)
	if err != nil {
		return err
	}

	if reservation.ID == "" {
		return errors.New("reservation not found")
	}

	err = uc.ReservationRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}