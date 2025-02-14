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
        },
        async addTask(namaTugas, deskripsi, deadline, files, idClass) {
            const isoDate = new Date(deadline).toISOString();

            const filePromises = files.map(file => {
                return new Promise((resolve) => {
                    const reader = new FileReader();
                    reader.onloadend = () => {
                        resolve({
                            tf_nama: file.name,
                            tf_content: reader.result.split(',')[1],
                            tf_type: file.type
                        });
                    };
                    reader.readAsDataURL(file);
                })
            });

            const filesData = await Promise.all(filePromises);

            try {
                await api.post(`dosen/addTask/${idClass}`, { td_judul: namaTugas, td_deskripsi: deskripsi, td_deadline: isoDate, files: filesData });
            } catch (error) {
                if (error.response && error.response.data) {
                    this.errorMessage = error.response.data.error;
                    const AUTH_STORE = useAuthStore();
                    await AUTH_STORE.refresh();
                    await api.post(`dosen/addTask/${idClass}`, { td_judul: namaTugas, td_deskripsi: deskripsi, td_deadline: deadline, files: filesData });
                    this.errorMessage = error.response.data.error;
                } else {
                    this.errorMessage = "An error occurred during adding task.";
                }
            }
        }
    }
});