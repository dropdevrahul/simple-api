package app

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/dropdevrahul/simple-api/internal/models"
)

func SignUpUser(a *App, u *models.User) (
	token *models.UserToken, err error) {

	tx, err := a.DB.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Print(err)
		return token, err
	}

	defer tx.Rollback()

	_, err = GetUserByEmail(a, u.Email)
	if err != nil && err != models.ErrNotFound {
		log.Print(err)
		return token, err
	}

	err = CreateUser(a, *u)
	if err != nil {
		log.Print(err)
		return token, err
	}

	user, err := GetUserByEmail(a, u.Email)
	if err != nil {
		log.Print(err)
		return token, err
	}

	ut, err := CreateUserToken(a, user.ID)
	if err != nil {
		log.Print(err)
		return token, err
	}

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return token, err
	}

	return ut, err
}

func CreateUser(a *App, u models.User) error {
	err := a.repos.User.Create(&a.DB, &u)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func GetUserByEmail(a *App, email string) (models.User, error) {
	var user models.User
	err := a.repos.User.GetByEmail(&a.DB, email, &user)

	return user, err
}

func CreateUserToken(a *App, userID string) (*models.UserToken, error) {
	token := uuid.New()
	ut := models.UserToken{
		UserID: userID,
		Token:  token.String(),
	}

	err := a.repos.UserToken.Create(&a.DB, &ut)
	if err != nil {
		return nil, err
	}

	return &ut, nil
}
