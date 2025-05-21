package dal

import (
	"numtostr/gotodo/config/database"

	"gorm.io/gorm"
)

// Todo struct defines the Todo Model
type Todo struct {
	gorm.Model
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	User      *uint  `gorm:"index,not null"`
	// this is a pointer because int == 0,
}

// CreateTodo create a todo entry in the todo's table
func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(todo)
}

// FindTodo finds a todo with given condition
func FindTodo(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Take(dest, conds...)
}

// FindTodoByUser finds a todo with given todo and user identifier
func FindTodoByUser(dest interface{}, todoIden interface{}, userIden interface{}) *gorm.DB {
	return FindTodo(dest, "id = ? AND user = ?", todoIden, userIden)
}

// FindTodosByUser finds the todos with user's identifier given
func FindTodosByUser(dest interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Find(dest, "user = ?", userIden)
}

// DeleteTodo deletes a todo from todos' table with the given todo and user identifier
func DeleteTodo(todoIden interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Unscoped().Delete(&Todo{}, "id = ? AND user = ?", todoIden, userIden)
}

// UpdateTodo allows to update the todo with the given todoID and userID
func UpdateTodo(todoIden interface{}, userIden interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ? AND user = ?", todoIden, userIden).Updates(data)
}

// CountTodosByUser counts total todos for a user with optional search
func CountTodosByUser(dest *int64, userIden interface{}, search string) *gorm.DB {
	query := database.DB.Model(&Todo{}).Where("user = ?", userIden)
	if search != "" {
		query = query.Where("task LIKE ?", "%"+search+"%")
	}
	return query.Count(dest)
}

// FindTodosByUserWithPagination finds paginated todos with optional search
func FindTodosByUserWithPagination(dest interface{}, userIden interface{}, search string, limit, offset int) *gorm.DB {
	query := database.DB.Model(&Todo{}).Where("user = ?", userIden)
	if search != "" {
		query = query.Where("task LIKE ?", "%"+search+"%")
	}
	return query.Limit(limit).Offset(offset).Order("created_at DESC").Find(dest)
}
