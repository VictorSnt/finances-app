package services_test

import (
	"controle/financeiro/domain/entities"
	"controle/financeiro/domain/services"
	"controle/financeiro/infra/repositories/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpenseService(t *testing.T) {
	userRepo := memory.NewInMemoryUserRepository()
	expenseRepo := memory.NewInMemoryExpenseRepository()
	expenseService := services.NewExpenseService(expenseRepo, userRepo)

	userRepo.Create(&entities.User{Username: "Victor", Income: 5000})
	user, err := userRepo.GetByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	t.Run("Adicionar gasto com sucesso", func(t *testing.T) {
		err := expenseService.AddExpense(user.ID, "Aluguel", 1500, entities.FixedCost, 1)
		assert.NoError(t, err)
	})

	t.Run("Não permitir gasto que ultrapassa a renda", func(t *testing.T) {
		err := expenseService.AddExpense(user.ID, "Viagem", 6000, entities.Leisure, 1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "excede sua renda mensal")
	})

	t.Run("Buscar gastos de um usuário", func(t *testing.T) {
		expenses, err := expenseService.GetUserExpenses(user.ID)
		assert.NoError(t, err)
		assert.Len(t, expenses, 1)
	})

	t.Run("Buscar gasto específico", func(t *testing.T) {
		expenses, _ := expenseService.GetUserExpenses(user.ID)
		expenseID := expenses[0].ID

		expense, err := expenseService.GetExpenseByID(expenseID)
		assert.NoError(t, err)
		assert.Equal(t, "Aluguel", expense.Name)
	})

	t.Run("Atualizar um gasto", func(t *testing.T) {
		expenses, _ := expenseService.GetUserExpenses(user.ID)
		expenseID := expenses[0].ID

		err := expenseService.UpdateExpense(expenseID, "Aluguel Atualizado", 1400, entities.FixedCost, 1)
		assert.NoError(t, err)

		updatedExpense, _ := expenseService.GetExpenseByID(expenseID)
		assert.Equal(t, "Aluguel Atualizado", updatedExpense.Name)
		assert.Equal(t, 1400.0, updatedExpense.Amount)
	})

	t.Run("Excluir um gasto", func(t *testing.T) {
		expenses, _ := expenseService.GetUserExpenses(user.ID)
		expenseID := expenses[0].ID

		err := expenseService.DeleteExpense(expenseID)
		assert.NoError(t, err)

		expense, err := expenseService.GetExpenseByID(expenseID)
		assert.Error(t, err)
		assert.Nil(t, expense)
	})

	t.Run("Calcular total de gastos", func(t *testing.T) {
		_ = expenseService.AddExpense(user.ID, "Mercado", 500, entities.Food, 1)
		_ = expenseService.AddExpense(user.ID, "Cinema", 100, entities.Leisure, 1)

		total, err := expenseService.GetTotalExpenses(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, 600.0, total)
	})

	t.Run("Calcular total de gastos por tipo", func(t *testing.T) {
		total, err := expenseService.GetTotalExpensesByType(user.ID, entities.Food)
		assert.NoError(t, err)
		assert.Equal(t, 500.0, total)
	})
}
