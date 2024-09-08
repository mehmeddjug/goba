package infrastructure

import (
	"errors"
	"sync"

	"github.com/mehmeddjug/goba/internal/domain"
)

type UserRepositoryMemory struct {
	users map[string]*domain.User
	mu    sync.Mutex
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{
		users: make(map[string]*domain.User),
	}
}

func (r *UserRepositoryMemory) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.users[user.ID]; err {
		return errors.New("user already exists")
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserRepositoryMemory) Read(id string) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, err := r.users[id]
	if !err {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepositoryMemory) Update(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.users[user.ID]; !err {
		return errors.New("user not found")
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserRepositoryMemory) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.users[id]; !err {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}
