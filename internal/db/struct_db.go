package db

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	TeamID   int    `gorm:"not null;index" json:"team_id"`
	Team     Team   `gorm:"foreignKey:TeamID" json:"team"`
}

type Team struct {
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Users []User `gorm:"foreignKey:TeamID" json:"users"`
}

type PRequest struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`                         // primaryKey, auto-increment
	AuthorID  int    `gorm:"not null" json:"author_id" `                                 // Обязательно для заполнения
	TeamID    int    `gorm:"not null" json:"team_id" `                                   // Обязательно для заполнения
	Title     string `gorm:"not null" json:"title"`                                      // Название PR
	Status    string `gorm:"not null;check:status in ('OPEN', 'MERGED')" json:"status" ` // Статус PR, с типом enum
	Reviewers []User `gorm:"many2many:assignments;" json:"reviewers"`                    // Это не поле для GORM, так как оно используется только для передачи данных
}

type Assignment struct {
	PRID   int `gorm:"primaryKey" json:"pr_id"`
	UserID int `gorm:"primaryKey" json:"user_id"`
}
