package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mehmeddjug/goba/users/internal/core"
	"github.com/mehmeddjug/goba/users/internal/ports"
)

type UserHandler struct {
	UserRepo        ports.UserRepository
	ActivityLogRepo ports.ActivityLogRepository
}

// CreateUser handles user creation.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user core.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.UserRepo.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log := core.ActivityLog{
		ID:        "gen-id",
		UserID:    user.ID,
		Action:    "Created user",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	h.ActivityLogRepo.LogActivity(log)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// UpdateUser handles user updates.
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user core.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.UserRepo.Update(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log := core.ActivityLog{
		ID:        "gen-id",
		UserID:    user.ID,
		Action:    "Update user",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	h.ActivityLogRepo.LogActivity(log)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser handles user deletion.
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	err := h.UserRepo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log := core.ActivityLog{
		ID:        "gen-id",
		UserID:    id,
		Action:    "Delete user",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	h.ActivityLogRepo.LogActivity(log)
	w.WriteHeader(http.StatusOK)
}

// GetUser handles fetching a single user.
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	user, err := h.UserRepo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log := core.ActivityLog{
		ID:        "gen-id",
		UserID:    user.ID,
		Action:    "Get user",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	h.ActivityLogRepo.LogActivity(log)
	json.NewEncoder(w).Encode(user)
}

// GetAll handles fetching all users.
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	users, err := h.UserRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
