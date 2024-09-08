package domain

type UserRepository interface {
	Create(user *User) error
	Read(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}
