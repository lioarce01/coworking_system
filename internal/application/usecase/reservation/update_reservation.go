package reservation

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
	"time"
)

type UpdateReservationUseCase struct {
	ReservationRepo ports.ReservationRepository
	SpaceRepo ports.SpaceRepository
	UserRepo ports.UserRepository
}

func NewUpdateReservationUseCase(userRepo ports.UserRepository, spaceRepo ports.SpaceRepository, reservationRepo ports.ReservationRepository) *UpdateReservationUseCase {
	return &UpdateReservationUseCase{
		ReservationRepo: reservationRepo,
		SpaceRepo:       spaceRepo,
		UserRepo:       userRepo,
	}
}

func (uc *UpdateReservationUseCase) Execute(id string, updatedFields entity.Reservation) (entity.Reservation, error) {
    existingReservation, err := uc.ReservationRepo.GetByID(id)
    if err != nil {
        return entity.Reservation{}, errors.New("reservation not found")
    }

    if updatedFields.SpaceID != "" && updatedFields.SpaceID != existingReservation.SpaceID {
        space, err := uc.SpaceRepo.GetByID(updatedFields.SpaceID)
        if err != nil {
            return entity.Reservation{}, errors.New("space not found")
        }
        
        existingReservation.SpaceID = updatedFields.SpaceID
        
        if !space.IsAvailable {
            return entity.Reservation{}, errors.New("space is not available")
        }

        activeReservationsCount, err := uc.ReservationRepo.CountActiveBySpace(space.ID)
        if err != nil {
            return entity.Reservation{}, err
        }

        if activeReservationsCount >= space.Capacity {
            return entity.Reservation{}, errors.New("space capacity exceeded")
        }

        updatedSpace, err := uc.SpaceRepo.Update(space)
        if err != nil {
            return entity.Reservation{}, err
        }
        space = updatedSpace
    }

    if updatedFields.UserID != "" && updatedFields.UserID != existingReservation.UserID {
        _, err := uc.UserRepo.GetByID(updatedFields.UserID)
        if err != nil {
            return entity.Reservation{}, errors.New("user not found")
        }
        existingReservation.UserID = updatedFields.UserID
    }

    if !updatedFields.StartTime.IsZero() || !updatedFields.EndTime.IsZero() {
        if updatedFields.StartTime.Before(updatedFields.EndTime) {
            existingReservations, err := uc.ReservationRepo.GetBySpace(existingReservation.SpaceID)
            if err != nil {
                return entity.Reservation{}, errors.New("error checking existing reservations")
            }
            for _, res := range existingReservations {
                if res.ID != id && updatedFields.StartTime.Before(res.EndTime) && updatedFields.EndTime.After(res.StartTime) {
                    return entity.Reservation{}, errors.New("time conflict with an existing reservation")
                }
            }
            existingReservation.StartTime = updatedFields.StartTime
            existingReservation.EndTime = updatedFields.EndTime
        } else {
            return entity.Reservation{}, errors.New("invalid time range")
        }
    }

    if updatedFields.Status != "" {
        existingReservation.Status = updatedFields.Status
    }

    existingReservation.UpdatedAt = time.Now()

    updatedReservation, err := uc.ReservationRepo.Update(*existingReservation)
    if err != nil {
        return entity.Reservation{}, err
    }

    return updatedReservation, nil
}
