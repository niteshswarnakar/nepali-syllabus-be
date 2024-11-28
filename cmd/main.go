package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/niteshswarnakar/nepali-syllabus-be/application"
	"github.com/niteshswarnakar/nepali-syllabus-be/internal/route"
)

func main() {
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})

	app := application.MyApp(".env")

	route.MakeRoute(server, app)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%v", app.Env.ServerPort)))

}
