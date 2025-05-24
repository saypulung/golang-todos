import { useEffect, useState } from 'react';
import { useTodos } from '../context/TodoContext';
import TodoItem from './TodoItem';

export default function TodoList() {
  const { todos, meta, isLoading, error, fetchTodos } = useTodos();
  const [search, setSearch] = useState('');
  const [currentPage, setCurrentPage] = useState(1);

  useEffect(() => {
    fetchTodos({ page: currentPage, search, limit: 10 });
  }, [fetchTodos, currentPage, search]);

  const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(e.target.value);
    setCurrentPage(1); // Reset ke halaman pertama saat mencari
  };

  if (isLoading && todos.length === 0) {
    return <div className="text-center py-4">Loading...</div>;
  }

  if (error) {
    return <div className="text-center text-red-500 py-4">{error}</div>;
  }

  if (todos.length === 0) {
    return <p className="text-center text-gray-500 mt-4">Tidak ada todo.</p>;
  }

  return (
    <div className="mt-4">
      <div className="mb-4">
        <input
          type="text"
          placeholder="Cari todo..."
          value={search}
          onChange={handleSearch}
          className="w-full p-2 border rounded"
        />
      </div>

      <div className="space-y-2">
        {todos.map(todo => (
          <TodoItem key={todo.id} todo={todo} />
        ))}
      </div>

      {/* Pagination */}
      {meta.total_page > 1 && (
        <div className="mt-4 flex justify-center gap-2">
          <button
            onClick={() => setCurrentPage(prev => Math.max(prev - 1, 1))}
            disabled={currentPage === 1}
            className="px-3 py-1 border rounded disabled:opacity-50"
          >
            Previous
          </button>
          <span className="px-3 py-1">
            Page {currentPage} of {meta.total_page}
          </span>
          <button
            onClick={() => setCurrentPage(prev => Math.min(prev + 1, meta.total_page))}
            disabled={currentPage === meta.total_page}
            className="px-3 py-1 border rounded disabled:opacity-50"
          >
            Next
          </button>
        </div>
      )}
    </div>
  );
}
