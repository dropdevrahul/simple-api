package app

import (
	"github.com/dropdevrahul/simple-api/internal/db"
	"github.com/dropdevrahul/simple-api/internal/models"
)

type App struct {
	DB    models.DB
	repos *db.DBRepo
}

func NewApp(db models.DB, r *db.DBRepo) *App {
	return &App{
		DB:    db,
		repos: r,
	}
}
