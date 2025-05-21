package services

import (
	"errors"
	"math"

	"maspulung/gotodo/app/entities"
	"maspulung/gotodo/app/repositories"
	"maspulung/gotodo/app/types"
	"maspulung/gotodo/config/database"
	"maspulung/gotodo/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var todoRepository = repositories.NewTodoRepository(database.DB)

// CreateTodo is responsible for create todo
func CreateTodo(c *fiber.Ctx) error {
	body := new(types.CreateDTO)

	if err := utils.ParseBodyAndValidate(c, body); err != nil {
		return err
	}

	todo := &entities.Todo{
		Task: body.Task,
		User: utils.GetUser(c),
	}

	if err := todoRepository.CreateTodo(todo).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TodoCreateResponse{
		Todo: &types.TodoResponse{
			ID:        todo.ID,
			Task:      todo.Task,
			Completed: todo.Completed,
		},
	})
}

// GetTodos returns the todos list with search and pagination
func GetTodos(c *fiber.Ctx) error {
	// Get query parameters
	search := c.Query("search", "")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Calculate offset
	offset := (page - 1) * limit

	todoResponse := &[]types.TodoResponse{}
	var total int64
	var totalPage int64

	// Get total count for pagination
	if err := todoRepository.CountTodosByUser(&total, utils.GetUser(c), search).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Calculate total pages
	totalPage = int64(math.Ceil(float64(total) / float64(limit)))

	// Get paginated and searched todos
	err := todoRepository.FindTodosByUserWithPagination(todoResponse, utils.GetUser(c), search, limit, offset).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(&types.TodosResponse{
		Todos: todoResponse,
		Pagination: &types.PaginationResponse{
			Page:      page,
			Limit:     limit,
			Total:     total,
			TotalPage: totalPage,
		},
	})
}

// GetTodo return a single todo
func GetTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	todoResponse := &types.TodoResponse{}

	err := todoRepository.FindTodoByUser(todoResponse, todoID, utils.GetUser(c)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	return c.JSON(&types.TodoCreateResponse{
		Todo: todoResponse,
	})
}

// DeleteTodo deletes a single todo
func DeleteTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	todoResponse := &types.TodoResponse{}

	var err = todoRepository.FindTodoByUser(todoResponse, todoID, utils.GetUser(c)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	res := todoRepository.DeleteTodo(todoID, utils.GetUser(c))
	if res.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusConflict, "Unable to delete todo")
	}

	err = res.Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully deleted",
	})
}

// CheckTodo TODO
func CheckTodo(c *fiber.Ctx) error {
	b := new(types.CheckTodoDTO)
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}
	todoResponse := &types.TodoResponse{}
	var err = todoRepository.FindTodoByUser(todoResponse, todoID, utils.GetUser(c)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	err = todoRepository.UpdateTodo(todoID, utils.GetUser(c), map[string]interface{}{"completed": b.Completed}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully updated",
	})
}

// UpdateTodoTitle TODO
func UpdateTodoTitle(c *fiber.Ctx) error {
	b := new(types.CreateDTO)
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	todoResponse := &types.TodoResponse{}
	var err = todoRepository.FindTodoByUser(todoResponse, todoID, utils.GetUser(c)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	err = todoRepository.UpdateTodo(todoID, utils.GetUser(c), &entities.Todo{Task: b.Task}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully updated",
	})
}
