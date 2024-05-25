package domain

import "context"

type User struct{
	ID int64 `db:"id"`
	FullName string `db:"full_name"`
	Phone string `db:"phone"`
	Username string `db: `
	Password string `db:"password"`
}

type UserInterface interface {
	FindByID(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
}

type UserService interface{
	Authenticatr()
}