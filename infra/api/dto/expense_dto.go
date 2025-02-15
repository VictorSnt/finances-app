package dto

import "controle/financeiro/domain/entities"

type CreateExpenseDTO struct {
	UserID      int                  `json:"user_id"`
	Name        string               `json:"name"`
	Amount      float64              `json:"amount"`
	ExpenseType entities.ExpenseType `json:"expense_type"`
	Recurrence  uint                 `json:"recurrence"`
}

type UpdateExpenseDTO struct {
	Name        string               `json:"name" binding:"omitempty"`
	Amount      float64              `json:"amount" binding:"omitempty,gt=0"`
	ExpenseType entities.ExpenseType `json:"expense_type" binding:"omitempty,gt=0"`
	Recurrence  uint                 `json:"recurrence" binding:"omitempty,gt=0"`
}

type ExpenseDTO struct {
	ID                 int                  `json:"expense_id"`
	UserID             int                  `json:"user_id"`
	ExpenseType        entities.ExpenseType `json:"expense_type"`
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	Amount             float64              `json:"amount"`
	RecurrenceInMonths uint                 `json:"recurrence"`
}
