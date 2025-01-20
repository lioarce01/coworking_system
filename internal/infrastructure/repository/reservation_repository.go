package repository

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"cowork_system/internal/utils"
	"time"

	"gorm.io/gorm"
)

type GormReservationRepository struct {
	DB *gorm.DB
}

func NewGormReservationRepository(db *gorm.DB) ports.ReservationRepository {
	return &GormReservationRepository{DB: db}
}


func (r *GormReservationRepository) Create(reservation entity.Reservation) (entity.Reservation, error) {
    if reservation.ID == "" {
        reservation.ID = utils.GenerateUUID()
    }

    result := r.DB.Create(&reservation)
    if result.Error != nil {
        return entity.Reservation{}, result.Error
    }

    var reservationWithDetails entity.Reservation
    result = r.DB.Preload("Space").Preload("User").First(&reservationWithDetails, "id = ?", reservation.ID)
    if result.Error != nil {
        return entity.Reservation{}, result.Error
    }

    return reservationWithDetails, nil
}

func (r *GormReservationRepository) Update(reservation entity.Reservation) (entity.Reservation, error) {
	result := r.DB.Updates(&reservation)
	if result.Error != nil {
		return entity.Reservation{}, result.Error
	}
	return reservation, nil
}

func (r *GormReservationRepository) Delete(id string) error {
	var reservation entity.Reservation
	result := r.DB.Where("id = ?", id).Delete(&reservation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormReservationRepository) GetAll() ([]entity.Reservation, error) {
	var reservations []entity.Reservation
	result := r.DB.Preload("Space").Preload("User").Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

func (r *GormReservationRepository) GetByID(id string) (*entity.Reservation, error) {
    var reservation entity.Reservation
    result := r.DB.Preload("Space").Preload("User").Where("id = ?", id).First(&reservation)

    if result.Error != nil {
        return nil, result.Error
    }

    return &reservation, nil
}

func (r *GormReservationRepository) GetBySpace(id string) ([]entity.Reservation, error) {
	var reservations []entity.Reservation
	result := r.DB.Where("space_id = ?", id).Preload("Space").Preload("User").Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

func (r *GormReservationRepository) GetByUser(id string) ([]entity.Reservation, error) {
	var reservations []entity.Reservation
	result := r.DB.Where("user_id = ?", id).Preload("Space").Preload("User").Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}

func (r *GormReservationRepository) CountActiveBySpace(spaceID string) (int, error) {
    var count int64
    result := r.DB.Model(&entity.Reservation{}).
        Where("space_id = ? AND status = ?", spaceID, entity.Confirmed).Count(&count)
    if result.Error != nil {
        return 0, result.Error
    }
    return int(count), nil
}

func (r *GormReservationRepository) GetBySpaceAndTime(spaceID string, startTime, endTime time.Time) ([]entity.Reservation, error) {
    var reservations []entity.Reservation
    result := r.DB.Where("space_id = ? AND ((start_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?))", spaceID, startTime, endTime, startTime, endTime).Find(&reservations)
    if result.Error != nil {
        return nil, result.Error
    }
    return reservations, nil
}
