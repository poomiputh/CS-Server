package models

type ReservationTime struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	RoomRefer uint   `json:"id_room" gorm:"default: null"`
	StartTime string `json:"starttime" gorm:"not null;default: null"`
	EndTime   string `json:"endtime" gorm:"not null;default: null"`
	StartDate string `json:"startdate" gorm:"default: null"`
	EndDate   string `json:"enddate" gorm:"default: null"`
	Type      string `json:"type" gorm:"default: null"`
}
