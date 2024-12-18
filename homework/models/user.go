package models

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Gender   int8   `db:"gender"`
	Email    string `db:"email"`
	Token    string
}
