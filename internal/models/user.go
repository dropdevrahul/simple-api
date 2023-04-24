package models

type User struct {
	ID       string `db:"id" json:"-"`
	Email    string `db:"email"`
	Password string `db:"password"` // salted hashed password using bcrypt
	Name     string `db:"name"`
}

type UserToken struct {
	ID     string `db:"id"`
	Token  string `db:"token"`
	UserID string `db:"user_id"`
}
