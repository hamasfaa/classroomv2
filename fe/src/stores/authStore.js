import { defineStore } from "pinia"
import axios from "axios"
import router from "@/router/index"
import { jwtDecode } from "jwt-decode"
import { errorMessages } from "vue/compiler-sfc"
import api from "@/service/api"

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: null,
        token: null,
        role: null,
        errorMessage: null
    }),
    actions: {
        async login(email, password) {
            try {
                const response = await api.post('login/', { email, password });
                this.token = response.data.token;
                localStorage.setItem('token', this.token);
                const decoded = jwtDecode(this.token);
                this.user = { uid: decoded.uid, email: decoded.email };
                this.role = decoded.role;
                errorMessages = null;
                // console.log(this.token);

            } catch (error) {
                if (error.response && error.response.data) {
                    this.errorMessage = error.response.data.error;
                } else {
                    this.errorMessage = "An error occurred during login.";
                }
            }
        },
        logout() {
            this.user = null;
            this.token = null;
            this.role = null;
            localStorage.removeItem('token');
        },
        loadToken() {
            const token = localStorage.getItem('token');
            if (token) {
                this.token = token;
                const decoded = jwtDecode(this.token);
                this.user = { uid: decoded.uid, email: decoded.email };
                this.role = decoded.role;
            }
        }
    }
})