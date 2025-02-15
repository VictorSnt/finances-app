package handlers

import (
	"controle/financeiro/domain/entities"
	"controle/financeiro/domain/services"
	"controle/financeiro/infra/api/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	service services.ExpenseService
}

func NewExpenseHandler(service services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

// @Summary Adicionar gasto
// @Description Adiciona um novo gasto para um usuário
// @Tags expenses
// @Accept json
// @Produce json
// @Param request body dto.CreateExpenseDTO true "Dados do gasto"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /expenses [post]
func (h *ExpenseHandler) AddExpense(c *gin.Context) {
	var req dto.CreateExpenseDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.service.AddExpense(req.UserID, req.Name, req.Amount, req.ExpenseType, req.Recurrence)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Expense added successfully"})
}

// @Summary Buscar gastos de um usuário
// @Description Retorna todos os gastos de um usuário
// @Tags expenses
// @Produce json
// @Param userID path int true "ID do usuário"
// @Success 200 {object} []dto.ExpenseDTO
// @Failure 404 {object} map[string]string
// @Router /expenses/user/{userID} [get]
func (h *ExpenseHandler) GetUserExpenses(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	expenses, err := h.service.GetUserExpenses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expensesDTO := make([]dto.ExpenseDTO, 0, len(expenses))

	for _, expense := range expenses {
		expenseDTO := dto.ExpenseDTO{
			ID:                 expense.ID,
			UserID:             expense.UserID,
			ExpenseType:        expense.ExpenseType,
			Name:               expense.Name,
			Description:        expense.Description,
			Amount:             expense.Amount,
			RecurrenceInMonths: expense.RecurrenceInMonths,
		}

		expensesDTO = append(expensesDTO, expenseDTO)
	}
	c.JSON(http.StatusOK, expensesDTO)

}

// @Summary Buscar gasto por ID
// @Description Retorna os detalhes de um gasto pelo ID
// @Tags expenses
// @Produce json
// @Param id path int true "ID do gasto"
// @Success 200 {object} dto.ExpenseDTO
// @Failure 404 {object} map[string]string
// @Router /expenses/{id} [get]
func (h *ExpenseHandler) GetExpenseByID(c *gin.Context) {
	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	expense, err := h.service.GetExpenseByID(expenseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	expenseDTO := dto.ExpenseDTO{
		ID:                 expense.ID,
		UserID:             expense.UserID,
		ExpenseType:        expense.ExpenseType,
		Name:               expense.Name,
		Description:        expense.Description,
		Amount:             expense.Amount,
		RecurrenceInMonths: expense.RecurrenceInMonths,
	}

	c.JSON(http.StatusOK, expenseDTO)
}

// @Summary Atualizar gasto
// @Description Atualiza os detalhes de um gasto pelo ID
// @Tags expenses
// @Accept json
// @Produce json
// @Param id path int true "ID do gasto"
// @Param request body dto.UpdateExpenseDTO true "Novos dados do gasto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /expenses/{id} [put]
func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	var req dto.UpdateExpenseDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.service.UpdateExpense(
		expenseID,
		req.Name,
		req.Amount,
		req.ExpenseType,
		req.Recurrence,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})
}

// @Summary Deletar gasto
// @Description Remove um gasto pelo ID
// @Tags expenses
// @Param id path int true "ID do gasto"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /expenses/{id} [delete]
func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	err = h.service.DeleteExpense(expenseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}

// @Summary Busca todos os gastos de um usuário
// @Description Retorna todos os gastos de um usuário
// @Tags expenses
// @Param id path int true "ID do usuário"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /expenses/total/{id} [get]
func (h *ExpenseHandler) GetTotalExpenses(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	total, err := h.service.GetTotalExpenses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_expense": total})
}

// @Summary Busca todos os gastos de um usuário por tipo
// @Description Retorna todos os gastos de um usuário de um determinado tipo
// @Tags expenses
// @Param id path int true "ID do usuário"
// @Param type path string true "Tipo do gasto"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /expenses/total/{id}/{type} [get]
func (h *ExpenseHandler) GetTotalExpensesByType(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req struct {
		ExpenseType entities.ExpenseType `json:"expense_type"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	total, err := h.service.GetTotalExpensesByType(userID, req.ExpenseType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_expense": total})
}
