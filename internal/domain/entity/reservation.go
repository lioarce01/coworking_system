package entity

import "time"

type Reservation struct {
	ID         string      `json:"id" gorm:"type:uuid;primaryKey"`
	SpaceID    string      `json:"space_id"`
	Space      Space     `json:"space,omitempty"` 
	UserID     string      `json:"user_id"`
	User       User      `json:"user,omitempty"` 
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status Status `json:"status" gorm:"default:'pending'"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Status string

const (
	Pending   Status = "pending"
	Confirmed Status = "confirmed"
	Cancelled Status = "cancelled"
)

func (s Status) IsValid() bool {
	switch s {
	case Pending, Confirmed, Cancelled:
		return true
	}
	return false
}

func (s Status) String() string {
	return string(s)
}