package db

import (
	"database/sql"
	"log"

	"github.com/dropdevrahul/simple-api/internal/models"
)

type DBRepo struct {
	User      UserRepo
	UserToken UserTokenRepo
}

func NewDBRepo() *DBRepo {
	return &DBRepo{
		User: &UserPG{
			TableName: "users",
		},
		UserToken: &UserTokenPG{
			TableName: "user_tokens",
		},
	}
}

func handleError(err error) error {
	log.Print(err)

	if err == sql.ErrNoRows {
		return models.ErrNotFound
	}

	return err
}
