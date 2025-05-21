import { useTodos } from '../context/TodoContext';
import TodoItem from './TodoItem';

export default function TodoList() {
  const { todos } = useTodos();

  if (todos.length === 0) {
    return <p className="text-center text-gray-500 mt-4">No todos yet.</p>;
  }

  return (
    <div className="mt-4">
      {todos.map(todo => (
        <TodoItem key={todo.id} todo={todo} />
      ))}
    </div>
  );
}
