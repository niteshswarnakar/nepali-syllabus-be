package route

import (
	"github.com/labstack/echo/v4"
	"github.com/niteshswarnakar/nepali-syllabus-be/internal/handler"
)

func CollegeRoute(apiRoute *echo.Group, collegeHandler *handler.College) {
	apiRoute.GET("/colleges", collegeHandler.CollegesHandler)
}
