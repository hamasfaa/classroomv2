import api from "@/service/api";
import { errorMessages } from "vue/compiler-sfc";
import { defineStore } from "pinia"
import { useAuthStore } from "./authStore";

export const useDosenStore = defineStore('dosen', {
    state: () => ({
        errorMessage: null,
        classList: []
    }),
    actions: {
        async addClass(namaKelas, mataKuliah) {
            try {
                const response = await api.post('dosen/addClass', { k_nama_kelas: namaKelas, k_mata_kuliah: mataKuliah });
            } catch (error) {
                if (error.response && error.response.data) {
                    this.errorMessage = error.response.data.error;
                    const AUTH_STORE = useAuthStore();
                    await AUTH_STORE.refresh();

                    await api.post('dosen/addClass', { k_nama_kelas: namaKelas, k_mata_kuliah: mataKuliah });
                    this.errorMessage = error.response.data.error;
                } else {
                    this.errorMessage = "An error occurred during adding class.";
                }
            }
        },
        async getAllClass() {
            try {
                const response = await api.get('dosen/class');
                this.classList = response.data.data;
            } catch (error) {
                if (error.response && error.response.data) {
                    const AUTH_STORE = useAuthStore();
                    await AUTH_STORE.refresh();
                    await this.getAllClass();
                    this.errorMessage = error.response.data.error;
                } else {
                    this.errorMessage = "An error occurred during fetching class.";
                }
            }
        },
        async deleteClass(id) {
            try {
                const response = await api.delete(`dosen/deleteClass/${id}`);
            } catch (error) {
                if (error.response && error.response.data) {
                    this.errorMessage = error.response.data.error;
                    const AUTH_STORE = useAuthStore();
                    await AUTH_STORE.refresh();
                    await api.delete(`dosen/deleteClass/${id}`);
                    this.errorMessage = error.response.data.error;
                } else {
                    this.errorMessage = "An error occurred during deleting class.";
                }
            }
        }
    }
});