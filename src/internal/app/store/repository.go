package store

import "github.com/behran/go-http-rest-api/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
