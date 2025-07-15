package userservice

import (
	"myproject/internal/taskservice"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string         `gorm:"type:varchar(255);not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`

	Tasks []taskservice.Task `gorm:"foreignKey:UserID"`
}
