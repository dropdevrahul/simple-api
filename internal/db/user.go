package db

import (
	"fmt"
	"log"

	"github.com/dropdevrahul/simple-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name UserRepo
type UserRepo interface {
	// create user
	Create(d *models.DB, u *models.User) error
	GetByEmail(d *models.DB, email string, u *models.User) error

	// set/update password
	//	SetPassword(d models.DB, u *models.User) error
}

// UserTokenRepo interface for `user_tokens` table
//
//go:generate mockery --name UserTokenRepo
type UserTokenRepo interface {
	Create(d *models.DB, ut *models.UserToken) error
	GetByUserID(d *models.DB, userID string, ut *models.UserToken) error
}

// UserPD postgres implementation for UserRepo

type UserPG struct {
	TableName string
}

func (ud *UserPG) GetPassword(u *models.User) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return "", err
	}

	return string(p), nil
}

func (ud *UserPG) Create(d *models.DB, u *models.User) error {
	p, err := ud.GetPassword(u)
	if err != nil {
		log.Print(err)
		return err
	}

	q := fmt.Sprintf(
		"Insert into %s (email, name, password) VALUES ($1, $2, $3)",
		ud.TableName)
	_, err = d.DB.Exec(q, u.Email, u.Name, p)

	return handleError(err)
}

// Fetches a User form DB into given u User pointer argument
func (ud *UserPG) GetByEmail(d *models.DB, email string, u *models.User) error {
	err := d.DB.Get(u, fmt.Sprintf("SELECT * from %s where email=$1", ud.TableName), email)

	return handleError(err)
}

type UserTokenPG struct {
	TableName string
}

func (ut *UserTokenPG) Create(d *models.DB, u *models.UserToken) error {
	q := fmt.Sprintf("Insert into %s (token, user_id) VALUES ($1, $2)", ut.TableName)
	_, err := d.DB.Exec(q, u.Token, u.UserID)

	return handleError(err)
}

func (ut *UserTokenPG) GetByUserID(d *models.DB,
	userID string, t *models.UserToken) error {
	err := d.DB.Get(t, fmt.Sprintf("Select * from %s where user_id=$1", ut.TableName), userID)

	return handleError(err)
}
