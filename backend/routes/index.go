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

func MainRoute(w http.ResponseWriter, r *http.Request) {
	lib.SendSuccess(w, map[string]string{
		"health": "good",
	}, "Healt check succeeded")
}
