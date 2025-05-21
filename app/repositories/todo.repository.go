package repositories

import (
	"maspulung/gotodo/app/entities"
	"maspulung/gotodo/config/database"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// CreateTodo create a todo entry in the todo's table
func (r *TodoRepository) CreateTodo(todo *entities.Todo) *gorm.DB {
	return database.DB.Create(todo)
}

// FindTodo finds a todo with given condition
func (r *TodoRepository) FindTodo(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&entities.Todo{}).Take(dest, conds...)
}

// FindTodoByUser finds a todo with given todo and user identifier
func (r *TodoRepository) FindTodoByUser(dest interface{}, todoIden interface{}, userIden interface{}) *gorm.DB {
	return r.FindTodo(dest, "id = ? AND user = ?", todoIden, userIden)
}

// FindTodosByUser finds the todos with user's identifier given
func (r *TodoRepository) FindTodosByUser(dest interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Model(&entities.Todo{}).Find(dest, "user = ?", userIden)
}

// DeleteTodo deletes a todo from todos' table with the given todo and user identifier
func (r *TodoRepository) DeleteTodo(todoIden interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Unscoped().Delete(&entities.Todo{}, "id = ? AND user = ?", todoIden, userIden)
}

// UpdateTodo allows to update the todo with the given todoID and userID
func (r *TodoRepository) UpdateTodo(todoIden interface{}, userIden interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&entities.Todo{}).Where("id = ? AND user = ?", todoIden, userIden).Updates(data)
}

// CountTodosByUser counts total todos for a user with optional search
func (r *TodoRepository) CountTodosByUser(dest *int64, userIden interface{}, search string) *gorm.DB {
	query := database.DB.Model(&entities.Todo{}).Where("user = ?", userIden)
	if search != "" {
		query = query.Where("task LIKE ?", "%"+search+"%")
	}
	return query.Count(dest)
}

// FindTodosByUserWithPagination finds paginated todos with optional search
func (r *TodoRepository) FindTodosByUserWithPagination(dest interface{}, userIden interface{}, search string, limit, offset int) *gorm.DB {
	query := database.DB.Model(&entities.Todo{}).Where("user = ?", userIden)
	if search != "" {
		query = query.Where("task LIKE ?", "%"+search+"%")
	}
	return query.Limit(limit).Offset(offset).Order("created_at DESC").Find(dest)
}
