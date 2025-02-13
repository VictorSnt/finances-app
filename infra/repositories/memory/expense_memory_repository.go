package memory

import (
	"errors"
	"sync"

	"controle/financeiro/domain/entities"
	"controle/financeiro/domain/repositories"
)

type ExpenseRepositoryMemory struct {
	expenses map[int]*entities.Expense
	mutex    sync.RWMutex
	nextID   int
}

func NewInMemoryExpenseRepository() repositories.ExpenseRepository {
	return &ExpenseRepositoryMemory{
		expenses: make(map[int]*entities.Expense),
		nextID:   1,
	}
}

func (r *ExpenseRepositoryMemory) Create(expense *entities.Expense) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	expense.ID = r.nextID
	r.nextID++
	r.expenses[expense.ID] = expense
	return nil
}

func (r *ExpenseRepositoryMemory) GetByID(id int) (*entities.Expense, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	expense, exists := r.expenses[id]
	if !exists {
		return nil, errors.New("gasto não encontrado")
	}
	return expense, nil
}

func (r *ExpenseRepositoryMemory) GetByUserID(userID int) ([]entities.Expense, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var userExpenses []entities.Expense
	for _, expense := range r.expenses {
		if expense.UserID == userID {
			userExpenses = append(userExpenses, *expense)
		}
	}
	return userExpenses, nil
}

func (r *ExpenseRepositoryMemory) GetTotalExpenseByUserID(userID int) (float64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	total := 0.0
	for _, expense := range r.expenses {
		if expense.UserID == userID {
			total += expense.Amount
		}
	}
	return total, nil
}

func (r *ExpenseRepositoryMemory) GetTotalExpenseByUserIDAndType(userID int, expenseType entities.ExpenseType) (float64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	total := 0.0
	for _, expense := range r.expenses {
		if expense.UserID == userID && expense.ExpenseType == expenseType {
			total += expense.Amount
		}
	}
	return total, nil
}

func (r *ExpenseRepositoryMemory) Update(expense *entities.Expense) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.expenses[expense.ID]; !exists {
		return errors.New("gasto não encontrado")
	}

	r.expenses[expense.ID] = expense
	return nil
}

func (r *ExpenseRepositoryMemory) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.expenses[id]; !exists {
		return errors.New("gasto não encontrado")
	}

	delete(r.expenses, id)
	return nil
}
