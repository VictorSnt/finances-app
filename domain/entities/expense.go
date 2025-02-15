package entities

import (
	"errors"
	"time"
)

type Expense struct {
	ID                 int
	UserID             int
	ExpenseType        ExpenseType
	Name               string
	Description        string
	Amount             float64
	RecurrenceInMonths uint
	CreatedAt          time.Time
}

func (e *Expense) Validate() error {
	if e.Amount <= 0 {
		return errors.New("o valor do gasto deve ser positivo")
	}

	if e.Name == "" {
		return errors.New("o gasto deve ter um nome")
	}

	if e.ExpenseType.String() == "N/A" {
		return errors.New("tipo de gasto inválido")
	}

	if e.RecurrenceInMonths == 0 {
		return errors.New("recorrência inválida")
	}

	return nil
}
