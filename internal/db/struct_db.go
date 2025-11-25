package db

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}

type Team struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Users []User `gorm:"many2many:team_users;" json:"users"`
}

type PR struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"not null" json:"title"`
	AuthorID uint   `gorm:"not null" json:"author_id"`
	Status   string `gorm:"not null;check:status in ('OPEN', 'MERGED')" json:"status"`
	Reviews  []User `gorm:"many2many:assignments;" json:"reviewers"`
}

type Assignment struct {
	PRID   uint `gorm:"primaryKey" json:"pr_id"`
	UserID uint `gorm:"primaryKey" json:"user_id"`
}
