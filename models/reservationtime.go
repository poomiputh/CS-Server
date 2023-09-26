package models

type ReservationTime struct {
	ID                    uint   `json:"id"`
	UserRefer             uint   `json:"user_refer" gorm:"not null;default: null"`
	AdminRefer            uint   `json:"admin_refer" gorm:"default: null"`
	RoomRefer             string `json:"room_refer" gorm:"not null;default: null"`
	CourseID              uint   `json:"course_id" gorm:"default: null"`
	CourseSection         uint   `json:"course_section" gorm:"default: null"`
	CourseName            string `json:"course_name" gorm:"default: null"`
	CourseType            string `json:"course_type" gorm:"default: null"`
	CourseInstructor      string `json:"course_instructor" gorm:"default: null"`
	CourseInstructorEmail string `json:"course_instructor_email" gorm:"default: null"`
	DayOfWeek             string `json:"day_of_week" gorm:"default: null"`
	Description           string `json:"description" gorm:"default: null"`
	StartTime             string `json:"start_time" gorm:"not null;default: null"`
	EndTime               string `json:"end_time" gorm:"not null;default: null"`
	Date                  string `json:"date" gorm:"not null;default: null"`
	Type                  string `json:"type" gorm:"not null;default: null"`
	Status                string `json:"status" gorm:"not null;default: null"`
}
