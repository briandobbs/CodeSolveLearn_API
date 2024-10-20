package models

import (
	"time"
)

type Article struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"size:255;not null"`
	Description string `json:"description" gorm:"type:text"`
	Body        string `json:"body" gorm:"type:text;not null"`
	AuthorID    uint   `json:"author_id" gorm:"not null"`
	Author      Author `gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
