package main

import (
	"controle/financeiro/infra/database"
)

func main() {
	db := database.GetDB()
	print(db == nil)
}
