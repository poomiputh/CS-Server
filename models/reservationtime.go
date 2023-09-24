package models

type ReservationTime struct {
	ID                 uint               `json:"id" gorm:"primaryKey"`
	RoomRefer          string               `json:"room_refer" gorm:"not null;default: null"`
	CourseRefer        uint               `json:"course_refer" gorm:"default: null"`
	StartTime          string             `json:"start_time" gorm:"not null;default: null"`
	EndTime            string             `json:"end_time" gorm:"not null;default: null"`
	StartDate          string             `json:"start_date" gorm:"not null;default: null"`
	EndDate            string             `json:"end_date" gorm:"not null;default: null"`
	Type               string             `json:"type" gorm:"not null;default: null"`
	RequestReservation RequestReservation `gorm:"foreignKey:ReservationTimeRefer"`
}
