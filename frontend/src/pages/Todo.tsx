import { useNavigate } from 'react-router-dom';
import TodoInput from '../components/TodoInput';
import TodoList from '../components/TodoList';
import { TodoProvider } from '../context/TodoProvider';

export default function Todo() {
  const navigate = useNavigate();
  const username = localStorage.getItem('username');

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('username');
    navigate('/login');
  };

  return (
    <TodoProvider>
      <div className="max-w-xl mx-auto mt-10 p-4 shadow-lg rounded-lg bg-white">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-2xl font-bold">My Todo List</h1>
          <div className="flex items-center gap-4">
            <span className="text-gray-600">Welcome, {username}</span>
            <button
              onClick={handleLogout}
              className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
            >
              Logout
            </button>
          </div>
        </div>
        <TodoInput />
        <TodoList />
      </div>
    </TodoProvider>
  );
} 