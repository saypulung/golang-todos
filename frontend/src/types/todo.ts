export interface Todo {
  id: string;
  task: string;
  completed: boolean;
  created_at?: string;
  updated_at?: string;
}

export interface TodoListMeta {
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

export interface TodoListPagination {
  page: number;
  limit: number;
  total: number;
  total_page: number;
}

export interface TodoListResponse {
  todos: Todo[];
  pagination: TodoListPagination;
}

export interface TodoListParams {
  page?: number;
  search?: string;
  limit?: number;
}

export interface CreateTodoRequest {
  task: string;
}

export interface UpdateTodoRequest {
  task: string;
}

export interface ToggleTodoRequest {
  completed: boolean;
}

export interface UpdateTodoResponse {
  message: string;
}
