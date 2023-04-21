package application

import (
	"github.com/dropdevrahul/simple-api/internal/db"
	"github.com/dropdevrahul/simple-api/internal/models"
)

type App struct {
	DB    models.DB
	repos db.DBRepo
}

func NewApp(db models.DB, r db.DBRepo) *App {
	return &App{
		DB:    db,
		repos: r,
	}
}

func CreateUser(a *App, u models.User) error {
	err := a.repos.User.Create(&a.DB, &u)
	if err != nil {
		return err
	}

	return nil
}
