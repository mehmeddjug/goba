package ports

import "github.com/mehmeddjug/goba/users/internal/core"

type UserRepository interface {
	Create(user *core.User) error
	Get(id string) (*core.User, error)
	Update(user *core.User) error
	Delete(id string) error
	GetAll() ([]*core.User, error)
}
