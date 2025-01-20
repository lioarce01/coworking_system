package reservation

import (
	"cowork_system/internal/application/ports"
	"errors"
)

type DeleteReservationUseCase struct {
	ReservationRepo ports.ReservationRepository
	SpaceRepo       ports.SpaceRepository
}

func NewDeleteReservationUseCase(repo ports.ReservationRepository, spaceRepo ports.SpaceRepository) *DeleteReservationUseCase {
	return &DeleteReservationUseCase{
		ReservationRepo: repo,
		SpaceRepo:       spaceRepo,
	}
}

func (uc *DeleteReservationUseCase) Execute(reservationID string) error {
    reservation, err := uc.ReservationRepo.GetByID(reservationID)
    if err != nil || reservation == nil {
        return errors.New("reservation not found")
    }

    if reservation.SpaceID == "" {
        return errors.New("reservation does not have a valid space ID")
    }

    space, err := uc.SpaceRepo.GetByID(reservation.SpaceID)
    if err != nil || space == nil {
        return errors.New("space not found")
    }

    err = uc.ReservationRepo.Delete(reservationID)
    if err != nil {
        return errors.New("failed to delete reservation")
    }

    activeReservationsCount, err := uc.ReservationRepo.CountActiveBySpace(space.ID)
    if err != nil {
        return errors.New("error fetching active reservations")
    }

    isAvailable := activeReservationsCount < space.Capacity
    if err := uc.SpaceRepo.SetAvailability(space.ID, isAvailable); err != nil {
        return errors.New("failed to update space availability")
    }

    return nil
}
