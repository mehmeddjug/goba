package main

import (
	"log"
	"net/http"

	"github.com/mehmeddjug/goba/content/internal/adapters/database"
	"github.com/mehmeddjug/goba/content/internal/adapters/handler"
	"github.com/mehmeddjug/goba/content/internal/core"
)

func main() {
	// Initialize the in-memory repository and service
	contentRepo := database.NewContentRepositoryMemory()

	// Initialize handlers with the in-memory repositories
	contentHandler := &handler.ContentHandler{
		ContentRepo: contentRepo,
	}

	// In-memory storage for simplicity
	contentRepo.Create(&core.Content{ID: "1", Username: "John", Title: "Title name 1"})
	contentRepo.Create(&core.Content{ID: "2", Username: "Jane", Title: "Title name 2"})
	contentRepo.Create(&core.Content{ID: "3", Username: "Jason", Title: "Title name 3"})
	contentRepo.Create(&core.Content{ID: "4", Username: "James", Title: "Title name 4"})

	// Create a new ServeMux and define routes
	mux := http.NewServeMux()
	mux.HandleFunc("/content", contentHandler.GetAll)
	mux.HandleFunc("/content/", contentHandler.Get) // Notice the trailing slash
	mux.HandleFunc("/content/create", contentHandler.Create)
	mux.HandleFunc("/content/update/", contentHandler.Update) // Notice the trailing slash
	mux.HandleFunc("/content/delete/", contentHandler.Delete) // Notice the trailing slash

	// Start the HTTP server
	log.Println("Start server on port 8090...")
	log.Fatal(http.ListenAndServe(":8090", mux))
}
