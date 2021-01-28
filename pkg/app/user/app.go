package user

import (
	"github.com/a-ovch/sa-course2020.09/pkg/common/infrastructure/database"
	"github.com/a-ovch/sa-course2020.09/pkg/user/app"
	"github.com/a-ovch/sa-course2020.09/pkg/user/infrastructure"
)

type App struct {
	appService *app.Service
}

func NewApplication(c database.Client) *App {
	ur := infrastructure.NewUserRepository(c)
	s := app.NewService(ur)

	return &App{appService: s}
}
