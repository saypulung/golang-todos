package repositories

import (
	"maspulung/gotodo/app/entities"
	"maspulung/gotodo/config/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser create a user entry in the user's table
func (r *UserRepository) CreateUser(user *entities.User) *gorm.DB {
	return database.DB.Create(user)
}

// FindUser searches the user's table with the condition given
func (r *UserRepository) FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&entities.User{}).Take(dest, conds...)
}

// FindUserByEmail searches the user's table with the email given
func (r *UserRepository) FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return r.FindUser(dest, "email = ?", email)
}
