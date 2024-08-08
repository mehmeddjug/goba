package core

type Content struct {
	ID       string `json:"id"`
	Title    string `json:"username"`
	Username string `json:"password"`
}

type ContentRepository interface {
	Create(user *Content) error
	Get(id string) (*Content, error)
	Update(user *Content) error
	Delete(id string) error
	GetAll() ([]*Content, error)
}
