package entity

type Account struct {
	Id       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
