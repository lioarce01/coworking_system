package repository

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"

	"gorm.io/gorm"
)

type GormReservationRepository struct {
	DB *gorm.DB
}

func NewGormReservationRepository(db *gorm.DB) ports.ReservationRepository {
	return &GormReservationRepository{DB: db}
}

func (r *GormReservationRepository) Create(reservation entity.Reservation) (entity.Reservation, error) {
	result := r.DB.Create(&reservation)
	if result.Error != nil {
		return entity.Reservation{}, result.Error
	}
	return reservation, nil
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
	result := r.DB.Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

func (r *GormReservationRepository) GetByID(id string) (entity.Reservation, error) {
	var reservation entity.Reservation
	result := r.DB.Where("id = ?", id).First(&reservation)
	if result.Error != nil {
		return entity.Reservation{}, result.Error
	}
	return reservation, nil
}

func (r *GormReservationRepository) GetBySpace(id string) ([]entity.Reservation, error) {
	var reservations []entity.Reservation
	result := r.DB.Where("space_id = ?", id).Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

func (r *GormReservationRepository) GetByUser(id string) ([]entity.Reservation, error) {
	var reservations []entity.Reservation
	result := r.DB.Where("user_id = ?", id).Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}

