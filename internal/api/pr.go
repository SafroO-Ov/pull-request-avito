package api

import (
	"log"
	"net/http"

	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/db"
	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/usecase"
	"github.com/labstack/echo/v4"
)

// Создание PR и автоматическое назначение ревьюеров
func createPR(c echo.Context) error {
	var pr db.PRequest
	if err := c.Bind(&pr); err != nil {
		log.Printf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Получаем ID автора
	authorID := pr.AuthorID
	teamID := pr.TeamID

	// Назначаем ревьюеров
	reviewers, err := usecase.AssignReviewers(pr.ID, authorID, teamID)
	if err != nil {
		log.Printf("Error assigning reviewers for PR %d: %v", pr.ID, err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Возвращаем успешный ответ с ревьюерами
	return c.JSON(http.StatusOK, reviewers)
}

// Переназначение ревьюера
func reassignReviewer(c echo.Context) error {
	prID := c.Param("pr_id")
	oldReviewerID := c.Param("old_reviewer_id")
	newReviewerID := c.Param("new_reviewer_id")

	// Конвертируем в int
	prIDint, err := ValidateID(prID)
	if err != nil {
		log.Printf("Error validating  prID: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	oldReviewerIDint, err := ValidateID(oldReviewerID)
	if err != nil {
		log.Printf("Error validating  oldReviewerIDint: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newReviewerIDint, err := ValidateID(newReviewerID)
	if err != nil {
		log.Printf("Error validating  newReviewerIDint: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Выполняем переназначение
	err = usecase.ReassignReviewer(prIDint, oldReviewerIDint, newReviewerIDint)
	if err != nil {
		log.Printf("Error reassigning reviewer for PR %d: %v", prIDint, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Reviewer successfully reassigned")
}
func updatePR(c echo.Context) error {
	prID := c.Param("pr_id")

	var updateData struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}

	// Читаем данные из тела запроса
	if err := c.Bind(&updateData); err != nil {
		log.Printf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Конвертируем в int
	prIDint, err := ValidateID(prID)
	if err != nil {
		log.Printf("Error validating  prID: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// Обновляем PR
	updatedPR, err := db.UpdatePR(prIDint, updateData.Title, updateData.Status)
	if err != nil {
		log.Printf("Error updating PR: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedPR)
}
func getPR(c echo.Context) error {
	prID := c.Param("pr_id")

	// Конвертируем в int
	prIDint, err := ValidateID(prID)
	if err != nil {
		log.Printf("Error validating  prID: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	pr, err := db.GetPRByID(prIDint)
	if err != nil {
		log.Printf("Error fetching PR: %v", err)
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, pr)
}

func updatePRStatus(c echo.Context) error {
	prID := c.Param("pr_id")
	status := c.Param("status")

	// Конвертируем в int
	prIDint, err := ValidateID(prID)
	if err != nil {
		log.Printf("Error converting prID to int: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid PR ID format")
	}
	// Обновляем статус PR
	updatedPR, err := db.UpdatePR(prIDint, "", status)
	if err != nil {
		log.Printf("Error updating PR status: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedPR)
}
