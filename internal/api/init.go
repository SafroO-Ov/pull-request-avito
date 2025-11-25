package api

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InitAPI(e *echo.Echo) {
	RegisterRoutes(e)
}

func RegisterRoutes(e *echo.Echo) {
	// Маршруты для PR
	// Маршруты для PR
	e.POST("/pr", createPR)                            // Создание нового PR
	e.GET("/pr/:pr_id", getPR)                         // Получение PR по ID
	e.PUT("/pr/:pr_id", updatePR)                      // Обновление данных PR
	e.PUT("/pr/:pr_id/status/:status", updatePRStatus) // Обновление статуса PR (например, на MERGED)

	// Маршруты для пользователей
	e.POST("/users", createUser)                                 // Создание нового пользователя
	e.GET("/users/:user_id", getUser)                            // Получение пользователя по ID
	e.PUT("/users/:user_id", updateUser)                         // Обновление данных пользователя
	e.PUT("/users/:user_id/status/:is_active", updateUserStatus) // Обновление статуса пользователя (активен/неактивен)

	// Маршруты для команд
	e.POST("/teams", createTeam)             // Создание новой команды
	e.GET("/teams/:team_id", getTeam)        // Получение информации о команде по ID
	e.PUT("/teams/:team_id", updateTeamName) // Обновление названия команды

	e.PUT("/teams/:team_id/users/:user_id", addUserToTeam)                            // Добавление пользователя в команду
	e.PUT("/pr/:pr_id/reviewers/:old_reviewer_id/:new_reviewer_id", reassignReviewer) // Переназначение ревьюера для PR
}

func ValidateID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: %v", err)
	}

	if id <= 0 {
		return 0, fmt.Errorf("ID must be a positive integer")
	}

	return id, nil
}
