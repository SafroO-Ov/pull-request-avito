package usecase

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SafroO-Ov/-Pull-Request-Avito-/internal/db"
)

// Инициализация генератора случайных чисел
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Функция для назначения ревьюеров для PR
func AssignReviewers(prID int, authorID int, teamID int) ([]int, error) {
	// Получаем команду по ID через репозиторий
	team, err := db.GetTeamByID(teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch team with ID %d: %w", teamID, err)
	}

	// Получаем список текущих ревьюеров для PR
	currentReviewers, err := db.GetReviewersForPR(prID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch current reviewers for PR %d: %w", prID, err)
	}

	// Создаём карту текущих ревьюеров для быстрого поиска
	existingReviewers := make(map[int]bool)
	for _, reviewer := range currentReviewers {
		existingReviewers[reviewer.ID] = true
	}

	// Фильтруем только активных пользователей, исключая автора и текущих ревьюеров
	var availableUsers []int
	for _, user := range team.Users {
		if user.ID != authorID && user.IsActive && !existingReviewers[user.ID] {
			availableUsers = append(availableUsers, user.ID)
		}
	}

	// Если нет доступных, можно не назначать, добавляем после этого функцию по поиску PR без reviewer!
	if len(availableUsers) == 0 {
		return nil, fmt.Errorf("no active users available for reviewer assignment")
	}

	// Назначаем ревьюеров (не более двух)
	var reviewers []int
	for i := 0; i < 2 && i < len(availableUsers); i++ {
		index := rand.Intn(len(availableUsers))
		reviewers = append(reviewers, availableUsers[index])
		// Удаляем выбранного пользователя из списка
		availableUsers = append(availableUsers[:index], availableUsers[index+1:]...)
	}

	// Назначаем ревьюеров в PR через репозиторий
	err = db.AssignReviewersToPR(prID, reviewers)
	if err != nil {
		return nil, fmt.Errorf("failed to assign reviewers to PR %d: %w", prID, err)
	}

	return reviewers, nil
}

// Переназначение ревьюера
func ReassignReviewer(prID int, oldReviewerID int, newReviewerID int) error {
	// Проверяем, что PR еще не закрыт через репозиторий
	pr, err := db.GetPRByID(prID)
	if err != nil {
		return fmt.Errorf("failed to fetch PR with ID %d: %w", prID, err)
	}

	// Если PR уже в статусе MERGED, возвращаем ошибку
	if pr.Status == "MERGED" {
		return fmt.Errorf("cannot modify reviewers for merged PR %d", prID)
	}

	// Переназначаем ревьюера через репозиторий
	err = db.ReassignReviewer(prID, oldReviewerID, newReviewerID)
	if err != nil {
		return fmt.Errorf("failed to reassign reviewer: %w", err)
	}

	return nil
}
