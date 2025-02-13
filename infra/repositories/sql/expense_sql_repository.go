package sql

import (
	"controle/financeiro/domain/entities"
	"controle/financeiro/infra/database/models"
	"errors"

	"gorm.io/gorm"
)

type ExpenseRepositorySQL struct {
	db *gorm.DB
}

func NewExpenseRepositorySQL(db *gorm.DB) *ExpenseRepositorySQL {
	return &ExpenseRepositorySQL{db: db}
}

func (r *ExpenseRepositorySQL) Create(expense *entities.Expense) error {
	model := models.Expense{
		UserID:             uint(expense.UserID),
		Name:               expense.Name,
		Amount:             expense.Amount,
		ExpenseType:        int(expense.ExpenseType),
		RecurrenceInMonths: expense.RecurrenceInMonths,
		CreatedAt:          expense.CreatedAt,
	}

	result := r.db.Create(&model)
	if result.Error != nil {
		return result.Error
	}

	expense.ID = int(model.ID)
	return nil
}

func (r *ExpenseRepositorySQL) GetByID(id int) (*entities.Expense, error) {
	var model models.Expense
	result := r.db.First(&model, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("gasto não encontrado")
	}

	return &entities.Expense{
		ID:                 int(model.ID),
		UserID:             int(model.UserID),
		Name:               model.Name,
		Amount:             model.Amount,
		ExpenseType:        entities.ExpenseType(model.ExpenseType),
		RecurrenceInMonths: model.RecurrenceInMonths,
		CreatedAt:          model.CreatedAt,
	}, result.Error
}

func (r *ExpenseRepositorySQL) GetByUserID(userID int) ([]entities.Expense, error) {
	var models []models.Expense
	result := r.db.Where("user_id = ?", userID).Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	expenses := make([]entities.Expense, len(models))
	for i, model := range models {
		expenses[i] = entities.Expense{
			ID:                 int(model.ID),
			UserID:             int(model.UserID),
			Name:               model.Name,
			Amount:             model.Amount,
			ExpenseType:        entities.ExpenseType(model.ExpenseType),
			RecurrenceInMonths: model.RecurrenceInMonths,
			CreatedAt:          model.CreatedAt,
		}
	}

	return expenses, nil
}

func (r *ExpenseRepositorySQL) GetTotalExpenseByUserID(userID int) (float64, error) {
	var total float64
	result := r.db.Model(&models.Expense{}).Where("user_id = ?", userID).Select("COALESCE(SUM(amount), 0)").Scan(&total)
	return total, result.Error
}

func (r *ExpenseRepositorySQL) GetTotalExpenseByUserIDAndType(userID int, expenseType entities.ExpenseType) (float64, error) {
	var total float64
	result := r.db.Model(&models.Expense{}).
		Where("user_id = ? AND expense_type = ?", userID, expenseType).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total)
	return total, result.Error
}

func (r *ExpenseRepositorySQL) Update(expense *entities.Expense) error {
	result := r.db.Model(&models.Expense{}).Where("id = ?", expense.ID).Updates(models.Expense{
		Name:               expense.Name,
		Amount:             expense.Amount,
		ExpenseType:        int(expense.ExpenseType),
		RecurrenceInMonths: expense.RecurrenceInMonths,
	})
	return result.Error
}

func (r *ExpenseRepositorySQL) Delete(id int) error {
	result := r.db.Delete(&models.Expense{}, id)
	return result.Error
}
