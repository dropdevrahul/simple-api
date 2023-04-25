package app

import (
	"github.com/dropdevrahul/simple-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(a *App, email, password string) (ut models.UserToken, err error) {
	var user models.User
	var userToken models.UserToken

	err = a.repos.User.GetByEmail(&a.DB, email, &user)
	if err != nil {
		return ut, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return ut, err
	}

	err = a.repos.UserToken.GetByUserID(&a.DB, user.ID, &userToken)
	if err == models.ErrNotFound {
		ut, err := CreateUserToken(a, user.ID)
		if err != nil {
			return userToken, err
		}

		userToken = *ut
	}

	return userToken, nil
}
