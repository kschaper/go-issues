package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kschaper/go-issues/big-interfaces"
)

// UserGetter represents a service to get a user.
type UserGetter interface {
	GetUser(int) (*whatever.User, error)
}

// GetUserHandler handles GET /user/123
func GetUserHandler(userGetter UserGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get ID from URL path
		parts := strings.Split(r.URL.String(), "/")
		id, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			return
		}

		user, err := userGetter.GetUser(id)
		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			return
		}
		fmt.Fprintf(w, "User ID: %d, Email: %s", user.ID, user.Email)
	}
}
