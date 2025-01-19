package ports

import (
	"cowork_system/internal/domain/entity"
	"time"
)

type ReservationRepository interface {
	GetAll() ([]entity.Reservation, error)
	GetByID(id string) (entity.Reservation, error)
	GetByUser(id string) ([]entity.Reservation, error)
	GetBySpace(id string) ([]entity.Reservation, error)
	Create(reservation entity.Reservation) (entity.Reservation, error)
	Update(reservation entity.Reservation) (entity.Reservation, error)
	Delete(id string) error
	CountActiveBySpace(spaceID string) (int, error)
	GetBySpaceAndTime(spaceID string, startTime time.Time, endTime time.Time) ([]entity.Reservation, error)
}
