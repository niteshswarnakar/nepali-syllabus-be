package route

import (
	"github.com/labstack/echo/v4"
	"github.com/niteshswarnakar/nepali-syllabus-be/application"
	"github.com/niteshswarnakar/nepali-syllabus-be/internal/handler"
	"github.com/niteshswarnakar/nepali-syllabus-be/internal/repository"
)

func MakeRoute(server *echo.Echo, app *application.App) {
	apiGroup := server.Group("/api")

	// LIST OF HANDLERS
	collegeHandler := handler.NewCollegeHandler(repository.NewCollegeRepository(app.DataSource, app.Env))

	CollegeRoute(apiGroup, collegeHandler)
}
