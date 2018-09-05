package db

import (
	"errors"

	whatever "github.com/kschaper/go-issues/big-interfaces"
)

// DatabaseClient manages database connectivity.
type DatabaseClient struct {
}

// Open opens a database connection.
func (c *DatabaseClient) Open() error {
	return nil
}

// UserService manages users.
type UserService struct {
	DatabaseClient *DatabaseClient
}

// GetUser returns the user with the given ID.
func (service *UserService) GetUser(id int) (*whatever.User, error) {
	// simulate db (in reality use service.DatabaseClient to access the real database)
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
