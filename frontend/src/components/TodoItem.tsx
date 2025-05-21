import type { Todo } from '../types/todo';
import { useTodos } from '../context/TodoContext';

export default function TodoItem({ todo }: { todo: Todo }) {
  const { toggleTodo, deleteTodo } = useTodos();

  return (
    <div className="flex justify-between items-center p-2 border-b">
      <span
        onClick={() => toggleTodo(todo.id)}
        className={`cursor-pointer ${todo.completed ? 'line-through text-gray-400' : ''}`}
      >
        {todo.text}
      </span>
      <button onClick={() => deleteTodo(todo.id)} className="text-red-500">✕</button>
    </div>
  );
}
