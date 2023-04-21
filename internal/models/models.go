package models

import "github.com/jmoiron/sqlx"

type DB struct {
	DB *sqlx.DB
}
