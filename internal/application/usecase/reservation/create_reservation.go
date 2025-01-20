package reservation

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
	"time"
)

type CreateReservationUseCase struct {
	ReservationRepo ports.ReservationRepository
	SpaceRepo       ports.SpaceRepository
	UserRepo        ports.UserRepository
}

func NewCreateReservationUseCase(reservationRepo ports.ReservationRepository, spaceRepo ports.SpaceRepository, userRepo ports.UserRepository) *CreateReservationUseCase {
	return &CreateReservationUseCase{
		ReservationRepo: reservationRepo,
		SpaceRepo:       spaceRepo,
		UserRepo:        userRepo,
	}
}

func (uc *CreateReservationUseCase) Execute(spaceID string, userID string, startTime time.Time, endTime time.Time, numPersons int) (entity.Reservation, error) {
	space, err := uc.SpaceRepo.GetByID(spaceID)
	if err != nil {
		return entity.Reservation{}, errors.New("space not found")
	}

	if startTime.After(endTime) {
		return entity.Reservation{}, errors.New("invalid time range")
	}

	activeReservationsCount, err := uc.ReservationRepo.CountActiveBySpace(spaceID)
	if err != nil {
		return entity.Reservation{}, errors.New("error fetching active reservations")
	}

	if activeReservationsCount+numPersons > space.Capacity {
		return entity.Reservation{}, errors.New("space is fully booked at the selected time")
	}

	reservation := entity.Reservation{
		SpaceID:    spaceID,
		UserID:     userID,
		StartTime:  startTime,
		EndTime:    endTime,
		NumPersons: numPersons,
		Status:     entity.Confirmed,
	}

	createdReservation, err := uc.ReservationRepo.Create(reservation)
	if err != nil {
		return entity.Reservation{}, err
	}

	totalActiveReservations := activeReservationsCount + numPersons
	isAvailable := totalActiveReservations < space.Capacity

	if err := uc.SpaceRepo.SetAvailability(spaceID, isAvailable); err != nil {
		return entity.Reservation{}, errors.New("error updating space availability")
	}

	return createdReservation, nil
}
