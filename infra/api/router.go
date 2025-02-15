package api

import (
	"controle/financeiro/infra/handlers"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configura e retorna o router do Gin
func SetupRouter(
	userHandler *handlers.UserHandler,
	expenseHandler *handlers.ExpenseHandler,

) *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	expenseRoutes := r.Group("/expenses")
	{
		expenseRoutes.POST("/", expenseHandler.AddExpense)
		expenseRoutes.GET("/user/:userID", expenseHandler.GetUserExpenses)
		expenseRoutes.GET("/:id", expenseHandler.GetExpenseByID)
		expenseRoutes.PUT("/:id", expenseHandler.UpdateExpense)
		expenseRoutes.DELETE("/:id", expenseHandler.DeleteExpense)
		expenseRoutes.GET("/total/:userID", expenseHandler.GetTotalExpenses)
		expenseRoutes.GET("/total/:userID/:type", expenseHandler.GetTotalExpensesByType)
	}

	return r
}
