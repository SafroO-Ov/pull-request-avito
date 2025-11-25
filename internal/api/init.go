package api

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InitAPI(e *echo.Echo) {
	RegisterRoutes(e)
}

// Функция для регистрации всех маршрутов
func RegisterRoutes(e *echo.Echo) {
	// Маршруты для PR
	e.POST("/pr", createPR)
	e.GET("/pr/:pr_id", getPR)    // Получение PR по ID
	e.PUT("/pr/:pr_id", updatePR) // Обновление PR
	e.PUT("/pr/:pr_id/status/:status", updatePRStatus)

	// Маршруты для пользователей
	e.POST("/users", createUser)                                 // Создание пользователя
	e.GET("/users/:user_id", getUser)                            // Получение пользователя по ID
	e.PUT("/users/:user_id", updateUser)                         // Обновление пользователя
	e.PUT("/users/:user_id/status/:is_active", updateUserStatus) // Обновление статуса пользователя

	// Маршруты для команд
	e.POST("/teams", createTeam)
	e.GET("/teams/:team_id", getTeam)        // Получение команды по ID
	e.PUT("/teams/:team_id", updateTeamName) // Обновление названия команды

	e.PUT("/teams/:team_id/users/:user_id", addUserToTeam)
	e.PUT("/pr/:pr_id/reviewers/:old_reviewer_id/:new_reviewer_id", reassignReviewer) // Переназначение ревьюера
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
