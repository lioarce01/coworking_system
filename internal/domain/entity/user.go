package entity

import "time"

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
	Role      Role    `json:"role" gorm:"default:'normal'"`

	Reservations []Reservation `json:"reservations,omitempty"` 
}

type Role string

const (
	Admin Role = "admin"
	Normal Role = "normal"
)

func (s Role) IsValid() bool {
	switch s {
		case Admin, Normal:
			return true
	}
	return false
}
