package models

type Reservation struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	RoomID      string `json:"room_id" gorm:"not null;default: null"`
	Instructor  string `json:"instructor"`
	Phone       string `json:"phone" gorm:"not null;default: null"`
	Description string `json:"description"`
	Date        string `json:"date" gorm:"not null;default: null"`
	TimeS       string `json:"time_start" gorm:"not null;default: null"`
	TimeE       string `json:"time_end" gorm:"not null;default: null"`
	Status      string `json:"status" gorm:"not null;default: null"`
}
