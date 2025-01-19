package entity

import "time"

type Reservation struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	SpaceID    uint      `json:"space_id"`
	Space      Space     `json:"space,omitempty"` 
	UserID     uint      `json:"user_id"`
	User       User      `json:"user,omitempty"` 
	ReservedAt time.Time `json:"reserved_at"`
}
