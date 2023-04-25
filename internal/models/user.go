package models

type User struct {
	ID       string `db:"id" json:"-"`
	Email    string `db:"email"`
	Password string `db:"password"` // salted hashed password using bcrypt
	Name     string `db:"name"`
}

type UserToken struct {
	ID     string `db:"id" json:"-"`
	Token  string `db:"token" json:"token"`
	UserID string `db:"user_id" json:"user_id"`
}
