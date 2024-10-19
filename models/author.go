package models

import (
	"gorm.io/gorm"
)

// Author represents an author with a bio and other details
type Author struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:255;unique"`
	Bio       string `gorm:"type:text"` // Author's bio
	CreatedAt gorm.DeletedAt
	UpdatedAt gorm.DeletedAt
}
