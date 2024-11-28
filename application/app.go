package application

import "github.com/niteshswarnakar/nepali-syllabus-be/infrastructure"

type App struct {
	Env        *infrastructure.Env
	DataSource *infrastructure.DatabaseConnection
}

func MyApp(envFile string) *App {
	envData := infrastructure.MyEnv(envFile)
	datasource := infrastructure.NewDatabaseConnection(envData)

	return &App{
		Env:        envData,
		DataSource: datasource,
	}
}
