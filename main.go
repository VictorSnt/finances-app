package main

import (
	"controle/financeiro/domain/services"
	"controle/financeiro/infra/api"
	"controle/financeiro/infra/database"
	"controle/financeiro/infra/handlers"
	"controle/financeiro/infra/repositories/sql"
	"fmt"

	_ "controle/financeiro/docs"
)

// @title Controle Financeiro API
// @version 1.0
// @description API para gerenciamento financeiro pessoal
// @BasePath /
func main() {
	db := database.GetDB()
	expenseSqlRepository := sql.NewExpenseRepositorySQL(db)
	userSqlRepository := sql.NewUserRepositorySQL(db)

	expenseService := services.NewExpenseService(expenseSqlRepository, userSqlRepository)
	userService := services.NewUserService(userSqlRepository)

	userHandler := handlers.NewUserHandler(userService)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	router := api.SetupRouter(userHandler, expenseHandler)

	fmt.Println("🚀 Servidor rodando em http://localhost:8080")
	fmt.Println("📖 Documentação Swagger em http://localhost:8080/swagger/index.html")
	router.Run(":80")
}
