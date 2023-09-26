package models

type Room struct {
	RoomID          string            `json:"room_id" gorm:"primaryKey"`
	DataReservation []ReservationTime `json:"data_reservationtime" gorm:"foreignKey:RoomRefer; references:RoomID; constraint:OnDelete:CASCADE;"`
}
