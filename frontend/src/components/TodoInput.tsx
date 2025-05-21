import { useState } from 'react';
import { useTodos } from '../context/TodoContext';

export default function TodoInput() {
  const [text, setText] = useState('');
  const { addTodo } = useTodos();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (text.trim()) {
      addTodo(text);
      setText('');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex gap-2">
      <input
        className="flex-grow border p-2 rounded"
        value={text}
        onChange={e => setText(e.target.value)}
        placeholder="What do you need to do?"
      />
      <button className="bg-blue-500 text-white px-4 py-2 rounded">Add</button>
    </form>
  );
}
