package entity

import "time"

type Reservation struct {
	ID         string      `json:"id" gorm:"type:uuid;primaryKey"`
	SpaceID    string      `json:"space_id"`
	Space      Space     `json:"space" gorm:"foreignKey:SpaceID"`
	UserID     string      `json:"user_id"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status Status `json:"status" gorm:"default:'confirmed'"`
	NumPersons int `json:"num_persons"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Status string

const (
	Confirmed Status = "confirmed"
	Cancelled Status = "cancelled"
)

func (s Status) IsValid() bool {
	switch s {
	case Confirmed, Cancelled:
		return true
	}
	return false
}

func (s Status) String() string {
	return string(s)
}