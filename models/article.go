package models

import (
	"gorm.io/gorm"
)

type Article struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Body        string `gorm:"type:text;not null"`
	AuthorID    uint   // Foreign key linking to Author
	Author      Author `gorm:"foreignKey:AuthorID"` // GORM will automatically join these
	CreatedAt   gorm.DeletedAt
	UpdatedAt   gorm.DeletedAt
}
