package models

type Room struct {
	Room_ID uint `json:"room_id" gorm:"not null;default: null"`
}
