package routes

import (
	"backend/lib"
	"backend/router"
	"net/http"
)

func SetupMainRoutes() *router.Router {
	r := router.New("")

	r.Get("/", MainRoute)

	return r
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func MainRoute(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:       "1",
		Username: "john_doe",
		Email:    "john@example.com",
	}
	lib.SendJSONResponse(w, user, 201)
}
