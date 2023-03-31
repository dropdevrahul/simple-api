package userrepo

import (
	"github.com/dropdevrahul/simple-api/internal/apiserver"
)

type User struct {
	id       string `db:id`
	Email    string `db:"email"`
	Password string `db:"password"` // salted hashed password using bcrypt
	Name     string `db:"name"`
}

type UserRepo interface {
	// create user
	Create(d apiserver.Db, u *User) error

	// set/update password
	SetPassword(d apiserver.Db, u *User) error
}

// postgres implementation for UserRepo
type UserPG struct {
}

func Create(d apiserver.Db, u *User) error {
	d.DB.Exec()

}
