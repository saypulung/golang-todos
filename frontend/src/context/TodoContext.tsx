import { createContext } from 'react';
import type { Todo, TodoListResponse, TodoListParams, UpdateTodoRequest } from '../types/todo';

interface TodoContextType {
  todos: Todo[];
  meta: TodoListResponse['pagination'];
  isLoading: boolean;
  error: string | null;
  fetchTodos: (params?: TodoListParams) => Promise<void>;
  addTodo: (task: string) => Promise<void>;
  updateTodo: (id: string, data: UpdateTodoRequest) => Promise<void>;
  toggleTodo: (id: string, completed: boolean) => Promise<void>;
  deleteTodo: (id: string) => Promise<void>;
}

export const TodoContext = createContext<TodoContextType | undefined>(undefined);
