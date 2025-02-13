package services

import "controle/financeiro/domain/entities"

type UserService interface {
	CreateUser(username string, income float64) (*entities.User, error)
	GetUserByID(userID int) (*entities.User, error)
	UpdateUser(userID int, username string, income float64) error
	DeleteUser(userID int) error
}

type ExpenseService interface {
	AddExpense(userID int, name string, amount float64, expenseType entities.ExpenseType, recurrence uint) error
	GetUserExpenses(userID int) ([]entities.Expense, error)
	GetExpenseByID(expenseID int) (*entities.Expense, error)
	UpdateExpense(expenseID int, name string, amount float64, expenseType entities.ExpenseType, recurrence uint) error
	DeleteExpense(expenseID int) error
	GetTotalExpenses(userID int) (float64, error)
	GetTotalExpensesByType(userID int, expenseType entities.ExpenseType) (float64, error)
}
