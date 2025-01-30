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
        async register(name, dob, role, email, password) {
            const isoDate = new Date(dob).toISOString();

            try {
                const response = await api.post('register/', { u_nama: name, u_tanggal_lahir: isoDate, u_role: role, u_email: email, u_password: password });
                this.user = response.data;
                this.errorMessage = null;
            } catch (error) {
                if (error.response && error.response.data) {
                    this.errorMessage = error.response.data.error;
                } else {
                    this.errorMessage = "An error occurred during registration.";
                }
            }
        },
        logout() {
            localStorage.removeItem('token');
            this.user = null;
            this.token = null;
            this.role = null;
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