package db

import (
	"github.com/dropdevrahul/simple-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	// create user
	Create(d *models.DB, u *models.User) error

	// set/update password
	//	SetPassword(d models.DB, u *models.User) error
}

// postgres implementation for UserRepo
type UserPG struct {
	TableName string
}

func (ud *UserPG) Create(d *models.DB, u *models.User) error {
	p, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	_, err = d.DB.Exec("Insert into $1 (id, email, name, password) VALUES ($2 $3 $4 $5)",
		ud.TableName,
		u.ID, u.Email, p, u.Name)
	if err != nil {
		return err
	}

	return nil
}
