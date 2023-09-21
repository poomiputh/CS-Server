package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type CourseBody struct {
	CourseID          uint   `json:"course_id"`
	CourseName        string `json:"coursename"`
	CourseDescription string `json:"coursedescription"`
	ReservationRefer  uint   `json:"id_reservationtime"`
}

func (h handler) AddCourse(c *fiber.Ctx) error {
	body := CourseBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var course models.Course
	course.CourseID = body.CourseID
	course.CourseName = body.CourseName
	course.CourseDescription = body.CourseDescription
	course.ReservationRefer = body.ReservationRefer

	if result := h.DB.Create(&course); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&course)
}

func (h handler) GetCourses(c *fiber.Ctx) error {
	var Courses []models.Course

	if result := h.DB.Preload("Data_Reservationtime").Find(&Courses); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&Courses)
}
func (h handler) GetCourse(c *fiber.Ctx) error {
	course := c.Params("id")
	var courses models.Course

	if result := h.DB.Preload("Data_Reservationtime").Find(&courses, course); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&courses)
}

func (h handler) UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	body := CourseBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var course models.Course
	course.CourseID = body.CourseID
	course.CourseName = body.CourseName
	course.CourseDescription = body.CourseDescription
	course.ReservationRefer = body.ReservationRefer

	if result := h.DB.First(&course, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	h.DB.Save(&course)
	return c.Status(fiber.StatusOK).JSON(&course)
}
