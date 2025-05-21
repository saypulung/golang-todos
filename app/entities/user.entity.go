package entities

import (
	"gorm.io/gorm"
)

// User struct defines the user
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Todos    []Todo `gorm:"foreignKey:User"`
}
