import { useState } from 'react';
import { http } from '../services/api';

interface LoginCredentials {
  email: string;
  password: string;
}

interface AuthResponse {
  auth: {
    token: string;
  };
}

interface UseAuthReturn {
  login: (credentials: LoginCredentials) => Promise<void>;
  isLoading: boolean;
  error: string | null;
}

export const useAuth = (): UseAuthReturn => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const login = async (credentials: LoginCredentials) => {
    try {
      setIsLoading(true);
      setError(null);
      
      const response = await http.post<AuthResponse>('/auth/login', credentials);
      
      if (response.auth.token) {
        localStorage.setItem('token', response.auth.token);
        localStorage.setItem('username', credentials.email);
      }
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message || 'Terjadi kesalahan saat login');
      } else {
        setError('Terjadi kesalahan yang tidak diketahui');
      }
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  return { login, isLoading, error };
}; 