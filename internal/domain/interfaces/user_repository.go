package interfaces

import "github.com/anthoz69/go-temfy/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id uint) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
	GetAll(limit, offset int) ([]*entities.User, error)
}
