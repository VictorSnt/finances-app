package repositories

import "controle/financeiro/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id int) (*entities.User, error)
	GetByUsername(username string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id int) error
}
