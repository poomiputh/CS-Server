package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type CourseBody struct {
	CourseID          uint   `json:"course_id"`
	CourseName        string `json:"course_name"`
	CourseDescription string `json:"course_description"`
	Type              string `json:"type"`
	DayofWeek         string `json:"dayofweek"`
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
	course.Type = body.Type
	course.DayofWeek = body.DayofWeek

	if result := h.DB.Create(&course); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&course)
}

func (h handler) GetCourses(c *fiber.Ctx) error {
	var Courses []models.Course

	if result := h.DB.Preload("ReservationTime").Find(&Courses); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&Courses)
}
func (h handler) GetCourse(c *fiber.Ctx) error {
	course := c.Params("id")
	var courses models.Course

	if result := h.DB.Preload("ReservationTime").Find(&courses, course); result.Error != nil {
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

	if result := h.DB.First(&course, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	course.CourseID = body.CourseID
	course.CourseName = body.CourseName
	course.CourseDescription = body.CourseDescription
	course.Type = body.Type
	course.DayofWeek = body.DayofWeek

	h.DB.Save(&course)

	return c.Status(fiber.StatusOK).JSON(&course)
}

func (h handler) DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")

	var Courses models.Course

	if result := h.DB.First(&Courses, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&id)

	return c.SendStatus(fiber.StatusOK)
}
