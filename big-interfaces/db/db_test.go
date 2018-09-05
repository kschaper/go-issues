package db_test

import (
	"testing"

	"github.com/kschaper/go-issues/big-interfaces/db"
)

func TestUserService_GetUser(t *testing.T) {
	service := &db.UserService{}
	user, err := service.GetUser(123)
	if err != nil {
		t.Fatalf("expected no error but got %q", err)
	}

	expected := "me@example.com"
	if user.Email != expected {
		t.Fatalf("expected user with email %q but got %q", expected, user.Email)
	}
}
