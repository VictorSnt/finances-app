package database

import (
	"controle/financeiro/infra"
	"controle/financeiro/infra/database/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	config := infra.LoadConfig()
	db, err := gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco: %v", err)
	}
	log.Println("Banco de dados conectado com sucesso!")
	migrateDB(db)
	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Expense{})
}
