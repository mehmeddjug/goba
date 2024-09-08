package application

import (
	"log"
	"net/http"

	"github.com/mehmeddjug/goba/internal/domain"
)

type UserHttpHandler struct {
	UserService *domain.UserService
	Mux         *http.ServeMux
}

func NewUserHttpHandler(mux *http.ServeMux, userService *domain.UserService) *UserHttpHandler {
	return &UserHttpHandler{
		UserService: userService,
		Mux:         mux,
	}
}

func (u *UserHttpHandler) RegisterRoutes() {
	u.Mux.HandleFunc("/users/", u.UserService.Read) // Notice the trailing slash
	u.Mux.HandleFunc("/users/create", u.UserService.Create)
	u.Mux.HandleFunc("/users/update/", u.UserService.Update) // Notice the trailing slash
	u.Mux.HandleFunc("/users/delete/", u.UserService.Delete) // Notice the trailing slash
}

func (u *UserHttpHandler) StartServer() {
	// Start the HTTP server
	log.Println("Start server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", u.Mux))
}
