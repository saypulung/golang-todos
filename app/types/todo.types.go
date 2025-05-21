package types

// TodoResponse struct contains the todo field which should be returned in a response
type TodoResponse struct {
	ID        uint   `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// CreateDTO struct defines the /todo/create payload
type CreateDTO struct {
	Task string `json:"task" validate:"required,min=3,max=150"`
}

// TodoCreateResponse struct defines the /todo/create response
type TodoCreateResponse struct {
	Todo *TodoResponse `json:"todo"`
}

// PaginationResponse defines the pagination metadata
type PaginationResponse struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

// TodosResponse defines the todos list with pagination
type TodosResponse struct {
	Todos      *[]TodoResponse     `json:"todos"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// CheckTodoDTO defined the payload for the check todo
type CheckTodoDTO struct {
	Completed bool `json:"completed"`
}
