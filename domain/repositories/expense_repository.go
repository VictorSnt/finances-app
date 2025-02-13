package repositories

import "controle/financeiro/domain/entities"

type ExpenseRepository interface {
	Create(expense *entities.Expense) error
	GetByID(id int) (*entities.Expense, error)
	GetByUserID(userID int) ([]entities.Expense, error)
	GetTotalExpenseByUserID(userID int) (float64, error)
	GetTotalExpenseByUserIDAndType(userID int, expenseType entities.ExpenseType) (float64, error)
	Update(expense *entities.Expense) error
	Delete(id int) error
}
