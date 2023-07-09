package main

import (
	"Cursorr/BankingSystem/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	router := models.NewRouter()

	router.SetupRoutes()

	err = router.Run("localhost:6969")
	if err != nil {
		return
	}
}
