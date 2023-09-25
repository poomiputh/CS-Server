package handler

import (
	"fmt"
	"go-fiber-api-docker/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ReservationTimeBody struct {
	UserRefer   uint   `json:"user_refer"`
	AdminRefer  uint   `json:"admin_refer"` // Can Null
	RoomRefer   string `json:"room_refer"`
	CourseID    uint   `json:"course_id"`           // Can Null
	CourseName  string `json:"course_name"`         // Can Null
	CourseType  string `json:"course_type"`         // Can Null
	Instructor  string `json:"instructor"`          // Can Null
	DayOfWeek   string `json:"dayofweek"`           // Can Null
	Description string `json:"request_description"` // Can Null
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Date        string `json:"date"`
	Type        string `json:"type"`
	Status      string `json:"status"`
}

func (h handler) AddReservationTime(c *fiber.Ctx) error {
	body := ReservationTimeBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var res_time models.ReservationTime
	res_time.UserRefer = body.UserRefer
	res_time.AdminRefer = body.AdminRefer
	res_time.RoomRefer = body.RoomRefer
	res_time.CourseID = body.CourseID
	res_time.CourseName = body.CourseName
	res_time.CourseType = body.CourseType
	res_time.Instructor = body.Instructor
	res_time.DayOfWeek = body.DayOfWeek
	res_time.Description = body.Description
	res_time.StartTime = body.StartTime
	res_time.EndTime = body.EndTime
	res_time.Date = body.Date
	res_time.Type = body.Type
	res_time.Status = body.Status

	if result := h.DB.Create(&res_time); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&res_time)
}

func (h handler) GetReservationTimes(c *fiber.Ctx) error {
	var ReservationTimes []models.ReservationTime

	if result := h.DB.Find(&ReservationTimes); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&ReservationTimes)
}

func (h handler) GetReservationTime(c *fiber.Ctx) error {
	reservationtime := c.Params("id")
	var reservationtimes models.ReservationTime

	if result := h.DB.Find(&reservationtimes, reservationtime); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&reservationtimes)
}

func (h handler) UpdateReservationTime(c *fiber.Ctx) error {
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
	res_time.CourseName = body.CourseName
	res_time.CourseType = body.CourseType
	res_time.Instructor = body.Instructor
	res_time.DayOfWeek = body.DayOfWeek
	res_time.Description = body.Description
	res_time.StartTime = body.StartTime
	res_time.EndTime = body.EndTime
	res_time.Date = body.Date
	res_time.Type = body.Type
	res_time.Status = body.Status

	if result := h.DB.First(&res_time, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	h.DB.Save(&res_time)
	return c.Status(fiber.StatusOK).JSON(&res_time)
}

func (h handler) DeleteReservationTime(c *fiber.Ctx) error {
	id := c.Params("id")

	var ReservationTimes models.ReservationTime

	if result := h.DB.First(&ReservationTimes, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&ReservationTimes)

	return c.SendStatus(fiber.StatusOK)
}

func (h handler) AddReservationTimeSeries(c *fiber.Ctx) error {

	startDate := time.Now()
	stopDate := time.Now().AddDate(0, 0, 30)

	body := ReservationTimeBody{}

	var res_time models.ReservationTime
	res_time.UserRefer = body.UserRefer
	res_time.AdminRefer = body.AdminRefer
	res_time.RoomRefer = body.RoomRefer
	res_time.CourseID = body.CourseID
	res_time.CourseName = body.CourseName
	res_time.CourseType = body.CourseType
	res_time.Instructor = body.Instructor
	res_time.DayOfWeek = body.DayOfWeek
	res_time.Description = body.Description
	res_time.StartTime = body.StartTime
	res_time.EndTime = body.EndTime
	res_time.Type = body.Type
	res_time.Status = body.Status

	for date := startDate; date.Before(stopDate); date = date.AddDate(0, 0, 1) {
		if date.Weekday() == time.Monday {

			res_time.Date = date.Format("2006-01-02")

			if result := h.DB.Create(&res_time); result.Error != nil {
				return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
			}

			fmt.Printf("Inserted record for %s\n", date.Format("2006-01-02"))
		}
	}

	return nil
}
