package api

import (
	"log"
	"net/http"

	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/db"
	"github.com/labstack/echo/v4"
)

// Создание команды
func createTeam(c echo.Context) error {
	var team db.Team
	if err := c.Bind(&team); err != nil {
		log.Printf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdTeam, err := db.CreateTeam(team.Name)
	if err != nil {
		log.Printf("Error creating team: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, createdTeam)
}

// Добавление пользователя в команду
func addUserToTeam(c echo.Context) error {
	teamID := c.Param("team_id")
	userID := c.Param("user_id")

	// Конвертируем в int
	teamIDInt, err := ValidateID(teamID)
	if err != nil {
		log.Printf("Error validating  teamIDInt: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userIDInt, err := ValidateID(userID)
	if err != nil {
		log.Printf("Error validating  userIDInt: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = db.AddUserToTeam(teamIDInt, userIDInt)
	if err != nil {
		log.Printf("Error adding user to team: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User added to team successfully")
}
func updateTeamName(c echo.Context) error {
	teamID := c.Param("team_id")
	newName := c.QueryParam("name")

	// Конвертируем в int
	teamIDInt, err := ValidateID(teamID)
	if err != nil {
		log.Printf("Error validating  teamIDInt: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	team, err := db.UpdateTeamName(teamIDInt, newName)
	if err != nil {
		log.Printf("Error updating team name: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, team)
}
func getTeam(c echo.Context) error {
	teamID := c.Param("team_id")

	// Конвертируем в int
	teamIDInt, err := ValidateID(teamID)
	if err != nil {
		log.Printf("Error validating  teamIDInt: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	team, err := db.GetTeamByID(teamIDInt)
	if err != nil {
		log.Printf("Error fetching team: %v", err)
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, team)
}
