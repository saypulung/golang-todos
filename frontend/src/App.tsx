import { useState } from 'react';
import TodoInput from './components/TodoInput';
import TodoList from './components/TodoList';
import { TodoProvider } from './context/TodoContext';
import Login from './pages/Login';

function App() {

  const [username, setUsername] = useState<string | null>(
    localStorage.getItem('username')
  );

  const handleLogin = (user: string) => {
    localStorage.setItem('username', user);
    setUsername(user);
  };

  if (!username) {
    return <Login onLogin={handleLogin} />;
  }

  return (
    <TodoProvider>
      <div className="max-w-xl mx-auto mt-10 p-4 shadow-lg rounded-lg bg-white">
        <h1 className="text-2xl font-bold mb-4 text-center">My Todo List</h1>
        <TodoInput />
        <TodoList />
      </div>
    </TodoProvider>
  );
}

export default App;
