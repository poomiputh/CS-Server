package models

type Course struct {
	ID                   uint              `json:"id" gorm:"primaryKey"`
	CourseID             uint              `json:"course_id" gorm:"unique;not null;default: null"`
	CourseName           string            `json:"coursename" gorm:"not null;default: null"`
	CourseDescription    string            `json:"coursedescription" gorm:"default: null"`
	ReservationRefer     uint              `json:"id_reservationtime" gorm:"unique;default: null"`
	Data_Reservationtime []ReservationTime `json:"data_reservationtime" gorm:"foreignKey:ID;references:ReservationRefer"`
}
