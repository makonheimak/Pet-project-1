package taskservice

type Task struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
}
