package models

type RequestReservation struct {
	RequestRefer         uint `json:"request_refer" gorm:"primaryKey"`
	ReservationTimeRefer uint `json:"reservationTime_refer" gorm:"primaryKey"`
}
