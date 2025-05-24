import axios from 'axios';
import type { InternalAxiosRequestConfig, AxiosResponse, AxiosError } from 'axios';

// Konfigurasi default untuk Axios instance
const baseURL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

// Membuat instance Axios dengan konfigurasi default
const api = axios.create({
  baseURL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Interceptor untuk request
api.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // Di sini Anda bisa menambahkan token autentikasi jika diperlukan
    const token = localStorage.getItem('token');
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);

// Interceptor untuk response
api.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error: AxiosError) => {
    // Handle error secara global di sini
    if (error.response) {
      // Server merespons dengan status code di luar range 2xx
      switch (error.response.status) {
        case 401:
          // Handle unauthorized
          localStorage.removeItem('token');
          window.location.href = '/login';
          break;
        case 403:
          // Handle forbidden
          console.error('Akses ditolak');
          break;
        case 404:
          // Handle not found
          console.error('Resource tidak ditemukan');
          break;
        case 500:
          // Handle server error
          console.error('Terjadi kesalahan pada server');
          break;
        default:
          console.error('Terjadi kesalahan:', error.response.data);
      }
    } else if (error.request) {
      // Request dibuat tapi tidak ada response
      console.error('Tidak ada response dari server');
    } else {
      // Ada error saat setup request
      console.error('Error:', error.message);
    }
    return Promise.reject(error);
  }
);

// Export fungsi-fungsi HTTP method yang sudah dikonfigurasi
export const http = {
  get: <T>(url: string, config?: InternalAxiosRequestConfig) => 
    api.get<T>(url, config).then((response: AxiosResponse<T>) => response.data),
  
  post: <T>(url: string, data?: any, config?: InternalAxiosRequestConfig) =>
    api.post<T>(url, data, config).then((response: AxiosResponse<T>) => response.data),
  
  put: <T>(url: string, data?: any, config?: InternalAxiosRequestConfig) =>
    api.put<T>(url, data, config).then((response: AxiosResponse<T>) => response.data),
  
  delete: <T>(url: string, config?: InternalAxiosRequestConfig) =>
    api.delete<T>(url, config).then((response: AxiosResponse<T>) => response.data),
  
  patch: <T>(url: string, data?: any, config?: InternalAxiosRequestConfig) =>
    api.patch<T>(url, data, config).then((response: AxiosResponse<T>) => response.data),
};

export default api; 