package core

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserRepository interface {
	Create(user *User) error
	Get(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
	GetAll() ([]*User, error)
}
