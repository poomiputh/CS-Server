package models

type Course struct {
	ID                uint            `json:"id" gorm:"primaryKey"`
	CourseID          uint            `json:"course_id" gorm:"not null;default: null"`
	CourseName        string          `json:"course_name" gorm:"not null;default: null"`
	CourseDescription string          `json:"course_description" gorm:"default: null"`
	Type              string          `json:"type" gorm:"not null;default: null"`
	DayofWeek         string          `json:"dayofweek" gorm:"not null;default: null"`
	ReservationTime   ReservationTime `json:"data_reservationtime" gorm:"foreignKey:CourseRefer"`
}
