import { useState } from 'react';
import type { ReactElement } from 'react';
import type { Todo } from '../types/todo';
import { useTodos } from '../context/TodoContext';

interface TodoItemProps {
  todo: Todo;
}

export default function TodoItem({ todo }: TodoItemProps): ReactElement {
  const { toggleTodo, updateTodo, deleteTodo } = useTodos();
  const [isEditing, setIsEditing] = useState(false);
  const [editText, setEditText] = useState(todo.task);

  const handleToggle = async (): Promise<void> => {
    try {
      await toggleTodo(todo.id, !todo.completed);
    } catch (err) {
      console.error('Error toggling todo:', err);
    }
  };

  const handleUpdate = async (): Promise<void> => {
    if (editText.trim() === todo.task) {
      setIsEditing(false);
      return;
    }

    try {
      await updateTodo(todo.id, { task: editText.trim() });
      setIsEditing(false);
    } catch (err) {
      console.error('Error updating todo:', err);
    }
  };

  const handleDelete = async (): Promise<void> => {
    if (window.confirm('Apakah Anda yakin ingin menghapus todo ini?')) {
      try {
        await deleteTodo(todo.id);
      } catch (err) {
        console.error('Error deleting todo:', err);
      }
    }
  };

  return (
    <div className="flex items-center gap-2 p-2 border rounded hover:bg-gray-50">
      <input
        type="checkbox"
        checked={todo.completed}
        onChange={handleToggle}
        className="h-5 w-5 rounded border-gray-300"
      />
      
      {isEditing ? (
        <input
          type="text"
          value={editText}
          onChange={(e) => setEditText(e.target.value)}
          onBlur={handleUpdate}
          onKeyDown={(e) => e.key === 'Enter' && handleUpdate()}
          className="flex-grow p-1 border rounded"
          autoFocus
        />
      ) : (
        <span
          onClick={() => setIsEditing(true)}
          className={`flex-grow cursor-pointer ${todo.completed ? 'line-through text-gray-400' : ''}`}
        >
          {todo.task}
        </span>
      )}

      <div className="flex gap-2">
        <button
          onClick={() => setIsEditing(!isEditing)}
          className="text-blue-500 hover:text-blue-700"
        >
          {isEditing ? '✕' : '✎'}
        </button>
        <button
          onClick={handleDelete}
          className="text-red-500 hover:text-red-700"
        >
          🗑
        </button>
      </div>
    </div>
  );
}
