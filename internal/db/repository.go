package db

import (
	"fmt"
)

// Получение команды по ID
func GetTeamByID(teamID int) (Team, error) {
	var team Team
	if err := DB.Preload("Users").First(&team, teamID).Error; err != nil {
		return team, fmt.Errorf("error fetching team: %w", err)
	}
	return team, nil
}

// Создание команды
func CreateTeam(name string) (Team, error) {
	team := Team{Name: name}
	if err := DB.Create(&team).Error; err != nil {
		return team, fmt.Errorf("failed to create team: %w", err)
	}
	return team, nil
}

// Добавление пользователя в команду
func AddUserToTeam(teamID, userID int) error {
	var team Team
	if err := DB.First(&team, teamID).Error; err != nil {
		return fmt.Errorf("team not found: %w", err)
	}

	var user User
	if err := DB.First(&user, userID).Error; err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Добавляем пользователя в команду
	if err := DB.Model(&team).Association("Users").Append(&user); err != nil {
		return fmt.Errorf("failed to add user to team: %w", err)
	}

	return nil
}
func UpdateTeamName(teamID int, newName string) (Team, error) {
	var team Team
	if err := DB.First(&team, teamID).Error; err != nil {
		return team, fmt.Errorf("team not found: %w", err)
	}

	team.Name = newName
	if err := DB.Save(&team).Error; err != nil {
		return team, fmt.Errorf("failed to update team name: %w", err)
	}

	return team, nil
}

func UpdatePR(prID int, newTitle string, newStatus string) (PRequest, error) {
	var pr PRequest
	if err := DB.First(&pr, prID).Error; err != nil {
		return pr, fmt.Errorf("PR not found: %w", err)
	}

	// Обновляем заголовок, если он передан
	if newTitle != "" {
		pr.Title = newTitle
	}

	// Обновляем статус, если он передан
	if newStatus != "" {
		pr.Status = newStatus
	}

	// Сохраняем обновленный PR
	if err := DB.Save(&pr).Error; err != nil {
		return pr, fmt.Errorf("failed to update PR: %w", err)
	}

	return pr, nil
}

// Назначение ревьюеров для PR
func AssignReviewersToPR(prID int, reviewers []int) error {
	for _, reviewerID := range reviewers {
		assignment := Assignment{PRID: prID, UserID: reviewerID}
		if err := DB.Create(&assignment).Error; err != nil {
			return fmt.Errorf("error assigning reviewer: %w", err)
		}
	}
	return nil
}

// Получение PR по ID
func GetPRByID(prID int) (PRequest, error) {
	var pr PRequest
	if err := DB.First(&pr, prID).Error; err != nil {
		return pr, fmt.Errorf("error fetching PR: %w", err)
	}
	return pr, nil
}

// Переназначение ревьюера
func ReassignReviewer(prID int, oldReviewerID int, newReviewerID int) error {
	// Удаляем старого ревьюера
	if err := DB.Where("pr_id = ? AND user_id = ?", prID, oldReviewerID).Delete(&Assignment{}).Error; err != nil {
		return fmt.Errorf("error deleting old reviewer: %w", err)
	}

	// Назначаем нового ревьюера
	assignment := Assignment{PRID: prID, UserID: newReviewerID}
	if err := DB.Create(&assignment).Error; err != nil {
		return fmt.Errorf("error assigning new reviewer: %w", err)
	}

	return nil
}

func GetReviewersForPR(prID int) ([]User, error) {
	var reviewers []User
	err := DB.Table("assignments").
		Select("users.id, users.name").
		Joins("JOIN users ON users.id = assignments.user_id").
		Where("assignments.pr_id = ?", prID).
		Find(&reviewers).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reviewers: %w", err)
	}
	return reviewers, nil
}

func GetUserByID(userID int) (User, error) {
	var user User
	if err := DB.Preload("Team").First(&user, userID).Error; err != nil {
		return user, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}
func CreateUser(name string, teamID int, isActive bool) (User, error) {
	user := User{Name: name, TeamID: teamID, IsActive: isActive}
	if err := DB.Create(&user).Error; err != nil {
		return user, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

// Обновление информации о пользователе
func UpdateUser(userID int, newName string, newStatus *bool, newTeamID *int) (User, error) {
	var user User
	if err := DB.First(&user, userID).Error; err != nil {
		return user, fmt.Errorf("user not found: %w", err)
	}

	// Обновляем имя, если оно передано
	if newName != "" {
		user.Name = newName
	}

	// Обновляем статус активности, если он передан
	if newStatus != nil {
		user.IsActive = *newStatus
	}

	// Смена команды, если передан новый TeamID
	if newTeamID != nil {
		// Удаляем пользователя из старой команды
		var oldTeam Team
		if err := DB.First(&oldTeam, user.TeamID).Error; err == nil {
			DB.Model(&oldTeam).Association("Users").Delete(&user)
		}

		// Добавляем пользователя в новую команду
		user.TeamID = *newTeamID
		var newTeam Team
		if err := DB.First(&newTeam, *newTeamID).Error; err != nil {
			return user, fmt.Errorf("new team not found: %w", err)
		}
		DB.Model(&newTeam).Association("Users").Append(&user)
	}

	if err := DB.Save(&user).Error; err != nil {
		return user, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}
