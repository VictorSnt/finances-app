package services

import (
	"controle/financeiro/domain/entities"
	"controle/financeiro/domain/repositories"
)

type UserServiceImp struct {
	userRepo repositories.UserRepository
}

func NewUserService(
	userRepo repositories.UserRepository,
) *UserServiceImp {
	return &UserServiceImp{
		userRepo: userRepo,
	}
}

func (s *UserServiceImp) CreateUser(username string, income float64) (*entities.User, error) {
	user := entities.User{
		Username: username,
		Income:   income,
	}

	ValidateErr := user.Validate()
	if ValidateErr != nil {
		return nil, ValidateErr
	}

	err := s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserServiceImp) GetUserByID(userID int) (*entities.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImp) UpdateUser(userID int, username string, income float64) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	user.Username = username
	user.Income = income

	ValidateErr := user.Validate()
	if ValidateErr != nil {
		return ValidateErr
	}

	err = s.userRepo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImp) DeleteUser(userID int) error {
	err := s.userRepo.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}
