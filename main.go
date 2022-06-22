package main

import (
	"awesomeProject/db"
	"awesomeProject/handlers"
	"awesomeProject/repositpries"
	"awesomeProject/services"
)

func main() {
	db.Migrate()
	authRepo := repositpries.AuthRepo{}
	authService := services.NewAuthService(authRepo)
	handler := handlers.NewHandler(authService)
	handler.Start()
}
