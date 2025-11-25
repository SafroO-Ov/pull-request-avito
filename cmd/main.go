package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/api"
	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	e := echo.New()

	// Инициализация API и зависимостей
	api.InitAPI(e)

	// Настройка порта
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // По умолчанию 8080
	}

	// Запуск сервера
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
