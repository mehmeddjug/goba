package ports

import "github.com/mehmeddjug/goba/content/internal/core"

type ContentRepository interface {
	Create(user *core.Content) error
	Get(id string) (*core.Content, error)
	Update(user *core.Content) error
	Delete(id string) error
	GetAll() ([]*core.Content, error)
}
