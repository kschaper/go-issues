package handler_test

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	whatever "github.com/kschaper/go-issues/big-interfaces"
	"github.com/kschaper/go-issues/big-interfaces/handler"
	"github.com/kschaper/go-issues/big-interfaces/mock"
)

func TestGetUserHandler(t *testing.T) {
	// set up
	var (
		user = &whatever.User{ID: 456, Email: "me@example.com"}

		userService = &mock.UserService{
			GetUserFunc: func(id int) (*whatever.User, error) {
				// ensure it was called with expected arg
				if id != user.ID {
					t.Fatalf("expected GetUser() to be called with id %d but was %d", user.ID, id)
				}
				return user, nil
			},
		}
	)

	// request
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", fmt.Sprintf("http://example.com/users/%d", user.ID), nil)
	handler.GetUserHandler(userService)(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	// ensure mocked method was called
	if !userService.GetUserCalled {
		t.Fatal("expected GetUser() to be called but wasn't")
	}

	// ensure output is as expected
	expected := fmt.Sprintf("User ID: %d, Email: %s", user.ID, user.Email)
	if string(body) != expected {
		t.Fatalf("expected to get %q but got %q", expected, body)
	}
}
