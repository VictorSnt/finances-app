package services

import (
	"controle/financeiro/domain/entities"
	"controle/financeiro/domain/repositories"
	"errors"
	"fmt"
	"time"
)

type ExpenseServiceImp struct {
	expenseRepo repositories.ExpenseRepository
	userRepo    repositories.UserRepository
}

func NewExpenseService(
	expenseRepo repositories.ExpenseRepository,
	userRepo repositories.UserRepository,
) *ExpenseServiceImp {
	return &ExpenseServiceImp{
		expenseRepo: expenseRepo,
		userRepo:    userRepo,
	}
}

func (s *ExpenseServiceImp) AddExpense(
	userID int,
	name string,
	amount float64,
	expenseType entities.ExpenseType,
	recurrence uint,
) error {

	expense := entities.Expense{
		UserID:             userID,
		Name:               name,
		Amount:             amount,
		ExpenseType:        expenseType,
		RecurrenceInMonths: recurrence,
		CreatedAt:          time.Now(),
	}

	expenseErr := expense.Validate()
	if expenseErr != nil {
		return expenseErr
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	expensesTotal, err := s.expenseRepo.GetTotalExpenseByUserID(userID)
	if err != nil {
		return errors.New("erro ao buscar gastos do usuário")
	}

	newTotal := expensesTotal + amount

	if newTotal > user.Income {
		return fmt.Errorf(
			"esse gasto excede sua renda mensal, sua renda é de %.2f e você tem %.2f disponível",
			user.Income, user.Income-expensesTotal,
		)
	}

	return s.expenseRepo.Create(&expense)
}

func (s *ExpenseServiceImp) GetUserExpenses(userID int) ([]entities.Expense, error) {
	return s.expenseRepo.GetByUserID(userID)
}

func (s *ExpenseServiceImp) GetExpenseByID(expenseID int) (*entities.Expense, error) {
	return s.expenseRepo.GetByID(expenseID)
}

func (s *ExpenseServiceImp) UpdateExpense(
	expenseID int,
	name string,
	amount float64,
	expenseType entities.ExpenseType,
	recurrence uint,
) error {

	expense, err := s.expenseRepo.GetByID(expenseID)
	if err != nil {
		return errors.New("gasto não encontrado")
	}

	expense.Name = name
	expense.Amount = amount
	expense.ExpenseType = expenseType
	expense.RecurrenceInMonths = recurrence

	expenseErr := expense.Validate()
	if expenseErr != nil {
		return expenseErr
	}

	return s.expenseRepo.Update(expense)
}

func (s *ExpenseServiceImp) DeleteExpense(expenseID int) error {
	return s.expenseRepo.Delete(expenseID)
}

func (s *ExpenseServiceImp) GetTotalExpenses(userID int) (float64, error) {
	return s.expenseRepo.GetTotalExpenseByUserID(userID)
}

func (s *ExpenseServiceImp) GetTotalExpensesByType(
	userID int,
	expenseType entities.ExpenseType,
) (float64, error) {
	return s.expenseRepo.GetTotalExpenseByUserIDAndType(userID, expenseType)
}
