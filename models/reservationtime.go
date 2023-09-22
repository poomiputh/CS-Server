package models

type ReservationTime struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	RoomRefer   uint     `json:"id_room" gorm:"not null;default: null"`
	StartTime   string   `json:"starttime" gorm:"not null;default: null"`
	EndTime     string   `json:"endtime" gorm:"not null;default: null"`
	StartDate   string   `json:"startdate" gorm:"not null;default: null"`
	EndDate     string   `json:"enddate" gorm:"not null;default: null"`
	Type        string   `json:"type" gorm:"not null;default: null"`
	CourseRefer uint     `json:"id_course" gorm:"default: null"`
	Data_Course []Course `json:"data_course" gorm:"foreignKey:ID;references:CourseRefer"`
}
