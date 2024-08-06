package database

import (
	"errors"
	"sync"

	"github.com/mehmeddjug/goba/users/internal/core"
)

type UserRepositoryMemory struct {
	users map[string]*core.User
	mu    sync.Mutex
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{
		users: make(map[string]*core.User),
	}
}

func (r *UserRepositoryMemory) Create(user *core.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}

func (r *UserRepositoryMemory) Get(id string) (*core.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepositoryMemory) Update(user *core.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *UserRepositoryMemory) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

func (r *UserRepositoryMemory) GetAll() ([]*core.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userList []*core.User
	for _, user := range r.users {
		userList = append(userList, user)
	}
	return userList, nil
}
