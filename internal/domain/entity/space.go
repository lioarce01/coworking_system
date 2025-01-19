package entity

type Space struct {
	ID          string `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}