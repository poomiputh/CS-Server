package models

type Request_Res struct {
	RequestRefer         uint      `json:"id_request" gorm:"unique;default: null"`
	Data_Request         []Request `json:"data_request" gorm:"foreignKey:ID;references:RequestRefer"`
	ReservationTimeRefer uint      `json:"id_reservationtime" gorm:"unique;default: null"`
	Data_ReservationTime []Request `json:"data_reservationtime" gorm:"foreignKey:ID;references:ReservationTimeRefer"`
}
