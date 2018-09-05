package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	whatever "github.com/kschaper/go-issues/big-interfaces"
)

// GetUserHandler handles GET /user/123
func GetUserHandler(env whatever.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get ID from URL path
		parts := strings.Split(r.URL.String(), "/")
		id, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			return
		}

		user, err := env.DB.UserService().GetUser(id)
		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			return
		}
		fmt.Fprintf(w, "User ID: %d, Email: %s", user.ID, user.Email)
	}
}
