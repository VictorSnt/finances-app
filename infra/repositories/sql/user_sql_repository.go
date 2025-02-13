package sql

import (
	"controle/financeiro/domain/entities"
	"controle/financeiro/infra/database/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepositorySQL struct {
	db *gorm.DB
}

func NewUserRepositorySQL(db *gorm.DB) *UserRepositorySQL {
	return &UserRepositorySQL{db: db}
}

func (r *UserRepositorySQL) Create(user *entities.User) error {
	model := models.User{
		Username: user.Username,
		Income:   user.Income,
	}

	result := r.db.Create(&model)
	if result.Error != nil {
		return result.Error
	}

	user.ID = int(model.ID)
	return nil
}

func (r *UserRepositorySQL) GetByID(id int) (*entities.User, error) {
	var model models.User
	result := r.db.First(&model, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("usuário não encontrado")
	}

	return &entities.User{
		ID:       int(model.ID),
		Username: model.Username,
		Income:   model.Income,
	}, result.Error
}

func (r *UserRepositorySQL) GetByUsername(username string) (*entities.User, error) {
	var model models.User
	result := r.db.Where("username = ?", username).First(&model)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("usuário não encontrado")
	}

	return &entities.User{
		ID:       int(model.ID),
		Username: model.Username,
		Income:   model.Income,
	}, result.Error
}

func (r *UserRepositorySQL) Update(user *entities.User) error {
	result := r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{
		Username: user.Username,
		Income:   user.Income,
	})
	return result.Error
}

func (r *UserRepositorySQL) Delete(id int) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
