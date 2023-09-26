package handler

import (
	"go-fiber-api-docker/models"
	"slices"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ReservationTimeBody struct {
	UserRefer             uint   `json:"user_refer"`
	AdminRefer            uint   `json:"admin_refer"`
	RoomRefer             string `json:"room_refer"`
	CourseID              uint   `json:"course_id"`
	CourseSection         uint   `json:"course_section"`
	CourseName            string `json:"course_name"`
	CourseType            string `json:"course_type"`
	CourseInstructor      string `json:"course_instructor"`
	CourseInstructorEmail string `json:"course_instructor_email"`
	DayOfWeek             string `json:"day_of_week"`
	LeadReservation       uint   `json:"lead_reservation" gorm:"default: null"`
	Description           string `json:"description"`
	StartTime             string `json:"start_time"`
	EndTime               string `json:"end_time"`
	StartDate             string `json:"start_date"`
	EndDate               string `json:"end_date"` // ใช้เฉพาะตอนเพิ่ม Course เป็นชุด, ไม่ใช้ตั้งเป็น null
	Type                  string `json:"type"`
	Status                string `json:"status"`
}

// สำหรับเพิ่ม Reservation ทั้งแบบเดี่ยวและเป็นชุด
func (h handler) AddReservation(c *fiber.Ctx) error {

	body := ReservationTimeBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var lead_res_time = models.ReservationTime{
		UserRefer:             body.UserRefer,
		AdminRefer:            body.AdminRefer,
		RoomRefer:             body.RoomRefer,
		CourseID:              body.CourseID,
		CourseSection:         body.CourseSection,
		CourseName:            body.CourseName,
		CourseType:            body.CourseType,
		CourseInstructor:      body.CourseInstructor,
		CourseInstructorEmail: body.CourseInstructorEmail,
		DayOfWeek:             body.DayOfWeek,
		LeadReservation:       body.LeadReservation,
		Description:           body.Description,
		StartTime:             body.StartTime,
		EndTime:               body.EndTime,
		StartDate:             body.StartDate,
		EndDate:               body.EndDate,
		Type:                  body.Type,
		Status:                body.Status,
	}

	if result := h.DB.Create(&lead_res_time); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	date_format := "02-01-2006"

	if body.EndDate != "" {
		start_date_str := body.StartDate
		start_date, err := time.Parse(date_format, start_date_str)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		stop_date_str := body.EndDate
		stop_date, err := time.Parse(date_format, stop_date_str)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		batch_trail_reservations := []models.ReservationTime{}

		day_split := strings.Split(body.DayOfWeek, ",")

		for date := start_date; date.Before(stop_date); date = date.AddDate(0, 0, 1) {

			var trail_res_time = models.ReservationTime{
				UserRefer:             body.UserRefer,
				AdminRefer:            body.AdminRefer,
				RoomRefer:             body.RoomRefer,
				CourseID:              body.CourseID,
				CourseSection:         body.CourseSection,
				CourseName:            body.CourseName,
				CourseType:            body.CourseType,
				CourseInstructor:      body.CourseInstructor,
				CourseInstructorEmail: body.CourseInstructorEmail,
				DayOfWeek:             body.DayOfWeek,
				LeadReservation:       lead_res_time.ID,
				Description:           body.Description,
				StartTime:             body.StartTime,
				EndTime:               body.EndTime,
				EndDate:               body.EndDate,
				Type:                  body.Type,
				Status:                body.Status,
			}

			if date.Weekday() == time.Monday && slices.Contains(day_split, "1") {
				trail_res_time.StartDate = date.Format(date_format)
				batch_trail_reservations = append(batch_trail_reservations, trail_res_time)
			}
			if date.Weekday() == time.Tuesday && slices.Contains(day_split, "2") {
				trail_res_time.StartDate = date.Format(date_format)
				batch_trail_reservations = append(batch_trail_reservations, trail_res_time)
			}
			if date.Weekday() == time.Wednesday && slices.Contains(day_split, "3") {
				trail_res_time.StartDate = date.Format(date_format)
				batch_trail_reservations = append(batch_trail_reservations, trail_res_time)
			}
			if date.Weekday() == time.Thursday && slices.Contains(day_split, "4") {
				trail_res_time.StartDate = date.Format(date_format)
				batch_trail_reservations = append(batch_trail_reservations, trail_res_time)
			}
			if date.Weekday() == time.Friday && slices.Contains(day_split, "5") {
				trail_res_time.StartDate = date.Format(date_format)
				batch_trail_reservations = append(batch_trail_reservations, trail_res_time)
			}
		}

		batch_trail_reservations = batch_trail_reservations[1:]
		if result := h.DB.Create(&batch_trail_reservations); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}

	}

	var result_reservation models.ReservationTime

	if result := h.DB.Preload("TrailReservations").First(&result_reservation, lead_res_time.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&result_reservation)
}

// สำหรับลบ Course ทั้งแบบเดี่ยวและเป็นชุด
func (h handler) DeleteCourseReservations(c *fiber.Ctx) error {
	del_course_id := c.Params("course_id")
	del_course_section := c.Params("course_section")
	del_course_type := c.Params("course_type")

	var reservation_times []models.ReservationTime

	if result := h.DB.Where("course_id = ? AND course_type = ? AND course_section = ? AND lead_reservation IS NULL", del_course_id, del_course_type, del_course_section).First(&reservation_times); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&reservation_times)

	return c.SendStatus(fiber.StatusOK)
}

// สำหรับลบ Reservation แบบเดี่ยว
func (h handler) DeleteReservation(c *fiber.Ctx) error {
	id := c.Params("id")

	var ReservationTimes models.ReservationTime

	if result := h.DB.First(&ReservationTimes, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&ReservationTimes)

	return c.SendStatus(fiber.StatusOK)
}

// สำหรับดึงค่า Reservation แบบเดี่ยว
func (h handler) GetReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation_times models.ReservationTime

	if result := h.DB.First(&reservation_times, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&reservation_times)
}

// สำหรับดึงค่า Reservation ทั้งหมด
func (h handler) GetAllReservations(c *fiber.Ctx) error {
	var ReservationTimes []models.ReservationTime

	if result := h.DB.Find(&ReservationTimes); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&ReservationTimes)
}

// สำหรับดึงค่า Reservation ทั้งหมดที่มี Type และ Status ที่ต้องการ
// Ex. http://localhost:3000/api/reservations/list
// Output: ค่า Reservation ทั้งหมด
// Ex. http://localhost:3000/api/reservations/list/all/approved
// Output: ค่า Reservation ทั้งหมดที่มี Status = approved
// Ex. http://localhost:3000/api/reservations/list/request/waiting
// Output: ค่า Reservation ทั้งหมดที่มี Type = request และ Status = waiting
func (h handler) GetAllReservationsByFilter(c *fiber.Ctx) error {
	reservation_type := c.Params("type")
	reservation_status := c.Params("status")

	var filtered_reservation_times []models.ReservationTime

	if reservation_type == "" {
		if result := h.DB.Find(&filtered_reservation_times); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
	} else if reservation_type == "all" {
		if result := h.DB.Where("status = ?", reservation_status).Find(&filtered_reservation_times); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
	} else if reservation_status == "" {
		if result := h.DB.Where("type = ?", reservation_type).Find(&filtered_reservation_times); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
	} else {
		if result := h.DB.Where("type = ? AND status = ?", reservation_type, reservation_status).Find(&filtered_reservation_times); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(&filtered_reservation_times)
}

// สำหรับดึงค่า Course ทั้งชุด
func (h handler) GetCourseReservations(c *fiber.Ctx) error {
	course_id := c.Params("course_id")
	course_section := c.Params("course_section")
	course_type := c.Params("course_type")

	var reservation_times []models.ReservationTime

	if result := h.DB.Preload("TrailReservations").Where("course_id = ? AND course_type = ? AND course_section = ? AND lead_reservation IS NULL", course_id, course_type, course_section).First(&reservation_times); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&reservation_times)
}

// Needed update
func (h handler) UpdateReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	body := ReservationTimeBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var res_time models.ReservationTime
	res_time.UserRefer = body.UserRefer
	res_time.AdminRefer = body.AdminRefer
	res_time.RoomRefer = body.RoomRefer
	res_time.CourseID = body.CourseID
	res_time.CourseSection = body.CourseSection
	res_time.CourseName = body.CourseName
	res_time.CourseType = body.CourseType
	res_time.CourseInstructor = body.CourseInstructor
	res_time.CourseInstructorEmail = body.CourseInstructorEmail
	res_time.DayOfWeek = body.DayOfWeek
	res_time.Description = body.Description
	res_time.StartTime = body.StartTime
	res_time.EndTime = body.EndTime
	res_time.StartDate = body.StartDate
	res_time.Type = body.Type
	res_time.Status = body.Status

	if result := h.DB.First(&res_time, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	h.DB.Save(&res_time)
	return c.Status(fiber.StatusOK).JSON(&res_time)
}
