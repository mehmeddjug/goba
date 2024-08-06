package main

import (
	"log"
	"net/http"

	"github.com/mehmeddjug/goba/users/internal/adapters/database"
	"github.com/mehmeddjug/goba/users/internal/adapters/handler"
	"github.com/mehmeddjug/goba/users/internal/core"
)

func main() {
	// Initialize the in-memory repository and service
	userRepo := database.NewUserRepositoryMemory()
	activityLogRepo := database.NewInMemoryActivityLogRepository()

	// Initialize handlers with the in-memory repositories
	userHandler := &handler.UserHandler{
		UserRepo:        userRepo,
		ActivityLogRepo: activityLogRepo,
	}

	// In-memory storage for simplicity
	userRepo.Create(&core.User{ID: "1", Username: "John", Password: "123"})
	userRepo.Create(&core.User{ID: "2", Username: "Jane", Password: "456"})
	userRepo.Create(&core.User{ID: "3", Username: "Jason", Password: "789"})
	userRepo.Create(&core.User{ID: "4", Username: "James", Password: "123"})

	// Create a new ServeMux and define routes
	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.GetAll)
	mux.HandleFunc("/users/", userHandler.Get) // Notice the trailing slash
	mux.HandleFunc("/users/create", userHandler.Create)
	mux.HandleFunc("/users/update/", userHandler.Update) // Notice the trailing slash
	mux.HandleFunc("/users/delete/", userHandler.Delete) // Notice the trailing slash

	// Start the HTTP server
	log.Println("Start server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
