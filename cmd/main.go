package main

import (
	"net/http"

	"github.com/mehmeddjug/goba/internal/adapters"
	"github.com/mehmeddjug/goba/internal/application"
	"github.com/mehmeddjug/goba/internal/domain"
	"github.com/mehmeddjug/goba/internal/infrastructure"
)

func main() {
	// Initialize the in-memory repository and service
	userRepo := infrastructure.NewUserRepositoryMemory()
	userService := adapters.NewUserService(userRepo)

	// Initialize handlers with the in-memory repositories
	mux := http.NewServeMux()
	userHandler := &application.UserHttpHandler{
		UserService: userService,
		Mux:         mux,
	}

	// In-memory storage for simplicity
	userRepo.Create(&domain.User{ID: "1", Username: "John", Password: "123"})
	userRepo.Create(&domain.User{ID: "2", Username: "Jane", Password: "456"})
	userRepo.Create(&domain.User{ID: "3", Username: "Jason", Password: "789"})
	userRepo.Create(&domain.User{ID: "4", Username: "James", Password: "123"})

	userHandler.RegisterRoutes()
	userHandler.StartServer()
}
