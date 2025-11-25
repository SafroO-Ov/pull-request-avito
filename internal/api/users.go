package api

import (
	"log"
	"net/http"

	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/db"
	"github.com/labstack/echo/v4"
)

// Обновление статуса активности пользователя
func updateUserStatus(c echo.Context) error {
	userID := c.Param("user_id")
	isActive := c.Param("is_active")

	// Конвертируем в int
	userIDInt, err := ValidateID(userID)
	if err != nil {
		log.Printf("Error validating  userIDInt: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	isActiveBool := isActive == "true"

	user, err := db.UpdateUser(userIDInt, "", &isActiveBool, nil)
	if err != nil {
		log.Printf("Error updating user status: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// Обновление пользователя
func updateUser(c echo.Context) error {
	userID := c.Param("user_id")

	var updateData struct {
		Name     string `json:"name"`
		IsActive *bool  `json:"is_active"`
		TeamID   *int   `json:"team_id"`
	}

	// Читаем данные из тела запроса
	if err := c.Bind(&updateData); err != nil {
		log.Printf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Конвертируем в int
	userIDInt, err := ValidateID(userID)
	if err != nil {
		log.Printf("Error validating  userIDInt: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// Обновляем пользователя
	updatedUser, err := db.UpdateUser(userIDInt, updateData.Name, updateData.IsActive, updateData.TeamID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func getUser(c echo.Context) error {
	userID := c.Param("user_id")

	// Конвертируем в int
	userIDInt, err := ValidateID(userID)
	if err != nil {
		log.Printf("Error validating  userIDInt: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := db.GetUserByID(userIDInt)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
	var user db.User

	// Читаем данные из тела запроса
	if err := c.Bind(&user); err != nil {
		log.Printf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdUser, err := db.CreateUser(user.Name, user.TeamID, user.IsActive)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, createdUser)
}
