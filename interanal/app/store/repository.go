package store

import "github.com/dannycrief/go-http-rest-api/interanal/app/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
