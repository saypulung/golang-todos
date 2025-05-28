import { useState } from 'react';
import type { ReactElement } from 'react';
import { useTodos } from '../hooks/useTodos';

export default function TodoInput(): ReactElement {
  const [text, setText] = useState('');
  const { addTodo, isLoading } = useTodos();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (text.trim()) {
      try {
        await addTodo(text.trim());
        setText('');
      } catch (err) {
        console.error('Error adding todo:', err);
      }
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex gap-2">
      <input
        className="flex-grow border p-2 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        value={text}
        onChange={e => setText(e.target.value)}
        placeholder="Apa yang perlu dilakukan?"
        disabled={isLoading}
      />
      <button 
        type="submit"
        disabled={isLoading || !text.trim()}
        className={`px-4 py-2 rounded text-white font-medium transition-colors
          ${isLoading || !text.trim()
            ? 'bg-blue-400 cursor-not-allowed'
            : 'bg-blue-500 hover:bg-blue-600'
          }`}
      >
        {isLoading ? 'Menambahkan...' : 'Tambah'}
      </button>
    </form>
  );
}
