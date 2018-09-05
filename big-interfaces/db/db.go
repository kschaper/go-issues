package db

import (
	"errors"

	whatever "github.com/kschaper/go-issues/big-interfaces"
)

// DatabaseClient allows access to services and handles database connectivity.
type DatabaseClient struct {
}

// UserService returns a UserService.
func (client *DatabaseClient) UserService() whatever.UserService {
	return &UserService{}
}

// UserService manages users.
type UserService struct {
}

// GetUser returns the user with the given ID.
func (service *UserService) GetUser(id int) (*whatever.User, error) {
	// simulate db
	users := []*whatever.User{
		&whatever.User{ID: 123, Email: "me@example.com"},
		&whatever.User{ID: 456, Email: "you@example.com"},
	}

	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}
