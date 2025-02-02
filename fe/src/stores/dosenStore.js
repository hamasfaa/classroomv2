import api from "@/service/api";
import { errorMessages } from "vue/compiler-sfc";
import { defineStore } from "pinia"

export const useDosenStore = defineStore('dosen', {
    state: () => ({
        errorMessage: null,
    }),
    actions: {
        async addClass(namaKelas, mataKuliah) {
            try {
                const response = await api.post('dosen/addClass', { k_nama_kelas: namaKelas, k_mata_kuliah: mataKuliah });
            } catch (error) {
                if (error.response && error.response.data) {
                    this.errorMessage = error.response.data.error;
                } else {
                    this.errorMessage = "An error occurred during adding class.";
                }
            }
        }
    }
});