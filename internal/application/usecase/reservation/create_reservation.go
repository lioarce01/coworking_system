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

	existingReservations, err := uc.ReservationRepo.GetBySpaceAndTime(spaceID, startTime, endTime)
	if err != nil {
		return entity.Reservation{}, errors.New("error fetching existing reservations")
	}

	totalReserved := 0
	for _, res := range existingReservations {
		totalReserved += res.NumPersons
	}

	if totalReserved+numPersons > space.Capacity {
		return entity.Reservation{}, errors.New("not enough capacity for this space at the selected time")
	}

	reservation := entity.Reservation{
		SpaceID:   spaceID,
		UserID:    userID,
		StartTime: startTime,
		EndTime:   endTime,
		Status:    entity.Pending,
		NumPersons: numPersons,
	}

	createdReservation, err := uc.ReservationRepo.Create(reservation)
	if err != nil {
		return entity.Reservation{}, err
	}

	space.Capacity -= numPersons

	if totalReserved+numPersons >= space.Capacity {
		space.IsAvailable = false
	}

    if _, err := uc.SpaceRepo.Update(space); err != nil {
        return entity.Reservation{}, err
    }

	return createdReservation, nil
}
