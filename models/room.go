package models

type Room struct {
	RoomID          string            `json:"room_id" gorm:"primaryKey; not null; default: null"`
	DataReservation []ReservationTime `json:"data_reservationtime" gorm:"foreignKey:RoomRefer; references:RoomID; constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
