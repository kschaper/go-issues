package mock

import (
	whatever "github.com/kschaper/go-issues/big-interfaces"
)

// UserService manages users.
type UserService struct {
	GetUserFunc   func(int) (*whatever.User, error)
	GetUserCalled bool
	// ... many more fields, two for each method of whatever.UserService
}

// GetUser is a mocked method.
func (service *UserService) GetUser(id int) (*whatever.User, error) {
	service.GetUserCalled = true
	return service.GetUserFunc(id)
}
