package memory

import (
	"controle/financeiro/domain/entities"
	"errors"
	"sync"
)

type UserRepositoryMemory struct {
	users  map[int]*entities.User
	mutex  sync.RWMutex
	nextID int
}

func NewInMemoryUserRepository() *UserRepositoryMemory {
	return &UserRepositoryMemory{
		users:  make(map[int]*entities.User),
		nextID: 1,
	}
}

func (r *UserRepositoryMemory) Create(user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	user.ID = r.nextID

	if _, exists := r.users[user.ID]; exists {
		return errors.New("usuário já existe")
	}

	r.nextID++

	r.users[user.ID] = user
	return nil
}

func (r *UserRepositoryMemory) GetByID(id int) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}

func (r *UserRepositoryMemory) GetByUsername(username string) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("usuário não encontrado")
}

func (r *UserRepositoryMemory) Update(user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("usuário não encontrado")
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserRepositoryMemory) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("usuário não encontrado")
	}

	delete(r.users, id)
	return nil
}
