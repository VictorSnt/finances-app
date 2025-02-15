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
	SeedDB(db)
}

func SeedDB(db *gorm.DB) {

	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		users := []models.User{
			{Username: "victor", Income: 5000},
			{Username: "ana", Income: 7000},
		}

		if err := db.Create(&users).Error; err != nil {
			log.Fatalf("Erro ao inserir usuários: %v", err)
		}
	}
}
