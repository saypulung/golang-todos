import { createContext, useContext, useState, useCallback } from 'react';
import type { ReactNode } from 'react';
import type { Todo, TodoListResponse, TodoListParams, CreateTodoRequest, UpdateTodoRequest, ToggleTodoRequest, UpdateTodoResponse } from '../types/todo';
import { http } from '../services/api';

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

const TodoContext = createContext<TodoContextType | undefined>(undefined);

export const TodoProvider = ({ children }: { children: ReactNode }) => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [meta, setMeta] = useState<TodoListResponse['pagination']>({
    total: 0,
    page: 1,
    limit: 10,
    total_page: 1
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchTodos = useCallback(async (params: TodoListParams = {}) => {
    try {
      setIsLoading(true);
      setError(null);
      
      const queryParams = new URLSearchParams();
      if (params.page) queryParams.append('page', params.page.toString());
      if (params.search) queryParams.append('search', params.search);
      if (params.limit) queryParams.append('limit', params.limit.toString());

      const response = await http.get<TodoListResponse>(`/todo/list?${queryParams.toString()}`);
      setTodos(response.todos);
      setMeta(response.pagination);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Gagal mengambil data todo');
      console.error('Error fetching todos:', err);
    } finally {
      setIsLoading(false);
    }
  }, []);

  const addTodo = async (task: string) => {
    try {
      setIsLoading(true);
      setError(null);
      
      await http.post<{ todo: Todo }>('/todo/create', { task });
      // setTodos(prev => [response.todo, ...prev]);
      fetchTodos({ page: 1, limit: meta.limit});
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Gagal menambahkan todo');
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  const updateTodo = async (id: string, data: UpdateTodoRequest) => {
    try {
      setIsLoading(true);
      setError(null);
      
      await http.patch<UpdateTodoResponse>(`/todo/${id}`, data);
      // Refresh todo list setelah update
      await fetchTodos({ page: meta.page, limit: meta.limit });
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Gagal mengupdate todo');
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  const toggleTodo = async (id: string, completed: boolean) => {
    try {
      setIsLoading(true);
      setError(null);
      
      await http.patch<UpdateTodoResponse>(`/todo/${id}/check`, { completed });
      // Refresh todo list setelah update
      await fetchTodos({ page: meta.page, limit: meta.limit });
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Gagal mengubah status todo');
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  const deleteTodo = async (id: string) => {
    try {
      setIsLoading(true);
      setError(null);
      
      await http.delete(`/todo/${id}`);
      // setTodos(prev => prev.filter(todo => todo.id !== id));
      fetchTodos({ page: meta.page, limit: meta.limit });
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Gagal menghapus todo');
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <TodoContext.Provider 
      value={{ 
        todos, 
        meta,
        isLoading, 
        error, 
        fetchTodos, 
        addTodo, 
        updateTodo, 
        toggleTodo, 
        deleteTodo 
      }}
    >
      {children}
    </TodoContext.Provider>
  );
};

export const useTodos = () => {
  const context = useContext(TodoContext);
  if (!context) throw new Error("useTodos must be used within a TodoProvider");
  return context;
};
