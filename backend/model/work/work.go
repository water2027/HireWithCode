package work

import (
	// "gorm.io/gorm"
	"time"
)

type Work struct {
	ID        uint   `gorm:"primaryKey"`
	GithubId  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	GithubUrl string `gorm:"type:varchar(100);"`
	OnlineUrl string `gorm:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
