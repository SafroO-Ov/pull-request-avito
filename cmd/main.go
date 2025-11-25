package main

import (
	"log"

	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/api"
	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.InitDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.RegisterRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
