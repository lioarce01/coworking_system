package entity

import "time"

type Space struct {
	ID          string    `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Capacity    int       `json:"capacity"`
	IsAvailable bool      `json:"is_available" gorm:"default:true"`
	Price       float64   `json:"price"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Reservations []Reservation `json:"reservations,omitempty" gorm:"foreignKey:SpaceID;constraint:OnDelete:CASCADE;"` 
	// OwnerID uint `json:"owner_id"`
	// Owner *User `json:"owner,omitempty"`
}