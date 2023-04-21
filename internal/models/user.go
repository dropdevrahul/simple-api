package models

type User struct {
	ID       string `db:id`
	Email    string `db:"email"`
	Password string `db:"password"` // salted hashed password using bcrypt
	Name     string `db:"name"`
}
