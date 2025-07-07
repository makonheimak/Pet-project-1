package taskservice

type Task struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Task string `json:"task"`
}
