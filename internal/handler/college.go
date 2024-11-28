package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/niteshswarnakar/nepali-syllabus-be/internal/repository"
)

type College struct {
	repository repository.ICollegeRepository
}

func NewCollegeHandler(repository repository.ICollegeRepository) *College {
	return &College{
		repository: repository,
	}
}

func (me *College) CollegesHandler(c echo.Context) error {

	colleges, err := me.repository.CollegesRepository()
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, colleges)
}
