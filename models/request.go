package models

type Request struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	RoomID       string      `json:"room_id" gorm:"not null;default: null"`
	UserRefer     uint        `json:"user_refer" gorm:"not null;default: null"`
	AdminRefer    uint        `json:"admin_refer" gorm:"default: null"`
	Instructor    string      `json:"instructor" gorm:"default: null"`
	ReDescription string      `json:"request_description" gorm:"default: null"`
	StartTime     string      `json:"start_time" gorm:"not null;default: null"`
	EndTime       string      `json:"end_time" gorm:"not null;default: null"`
	Date          string      `json:"date" gorm:"not null;default: null"`
	Status        string      `json:"status" gorm:"not null;default: null"`
	RequestReservation   RequestReservation `gorm:"foreignKey:RequestRefer"`
}
