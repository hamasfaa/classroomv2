import { useAuthStore } from '@/stores/authStore';
import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:3000/api',
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
    },
});

api.interceptors.request.use((config) => {
    const AUTH_STORE = useAuthStore();
    if (AUTH_STORE.token) {
        config.headers.Authorization = `Bearer ${AUTH_STORE.token}`;
    }
    return config;
});

export default api;