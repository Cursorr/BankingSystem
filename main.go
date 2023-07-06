package main

import (
	"Cursorr/BankingSystem/models"
)

func main() {
	router := models.NewRouter()

	router.SetupRoutes()

	err := router.Run("localhost:6969")
	if err != nil {
		return
	}
}
