package main

import (
	"backend/middleware"
	"backend/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	indexRouter := routes.SetupMainRoutes()

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", indexRouter))

	handler := middleware.LoggingMiddleware(middleware.CORSMiddleware(mux))

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Printf("Server starting on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
