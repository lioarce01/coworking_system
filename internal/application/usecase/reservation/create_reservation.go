package reservation

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"cowork_system/internal/utils"
	"errors"
)

type CreateReservationUseCase struct {
	ReservationRepo ports.ReservationRepository
	SpaceRepo ports.SpaceRepository
	UserRepo ports.UserRepository
}

func NewCreateReservationUseCase(reservationRepo ports.ReservationRepository, spaceRepo ports.SpaceRepository, userRepo ports.UserRepository) *CreateReservationUseCase {
	return &CreateReservationUseCase{
		ReservationRepo: reservationRepo,
		SpaceRepo:       spaceRepo,
		UserRepo:        userRepo,
	}
}

func (uc *CreateReservationUseCase) Execute(reservation entity.Reservation) (entity.Reservation, error) {
	space, err := uc.SpaceRepo.GetByID(reservation.SpaceID)
	if err != nil {
		return entity.Reservation{}, errors.New("space not found")
	}

	if !space.IsAvailable {
		return entity.Reservation{}, errors.New("space is not available")
	}

	activeReservations, err := uc.ReservationRepo.CountActiveBySpace(reservation.SpaceID)
	if err != nil {
		return entity.Reservation{}, errors.New("error checking existing reservations")
	}

	if activeReservations >= space.Capacity {
		return entity.Reservation{}, errors.New("space has reached its maximum capacity")
	}

	_, err = uc.UserRepo.GetByID(reservation.UserID)
	if err != nil {
		return entity.Reservation{}, errors.New("user not found")
	}

	if !reservation.StartTime.Before(reservation.EndTime) {
		return entity.Reservation{}, errors.New("start time must be before end time")
	}

	existingReservations, err := uc.ReservationRepo.GetBySpace(reservation.SpaceID)
	if err != nil {
		return entity.Reservation{}, errors.New("error checking existing reservations")
	}

	for _, res := range existingReservations {
		if reservation.StartTime.Before(res.EndTime) && reservation.EndTime.After(res.StartTime) {
			return entity.Reservation{}, errors.New("reservation conflicts with an existing reservation")
		}
	}

	reservation.ID = utils.GenerateUUID()
	reservation.Status = entity.Pending

	createdReservation, err := uc.ReservationRepo.Create(reservation)
	if err != nil {
		return entity.Reservation{}, err
	}

	space.Capacity -= 1
	_, err = uc.SpaceRepo.Update(space)
	if err != nil {
		return entity.Reservation{}, errors.New("error updating space capacity")
	}

	return createdReservation, nil
}
