package taskservice

type Task struct {
	ID   int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Task string `json:"task"`
}
