<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <form @submit.prevent="handleAddClass">
      <div class="mb-6">
        <label
          for="namaKelas"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Nama Kelas:</label
        >
        <input
          type="text"
          v-model="namaKelas"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
          placeholder="Masukkan nama kelas"
        />
      </div>
      <div class="mb-6">
        <label
          for="mataKuliah"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Mata Kuliah:</label
        >
        <input
          type="text"
          v-model="mataKuliah"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
          placeholder="Masukkan mata kuliah"
        />
      </div>
      <div class="flex items-center justify-between">
        <button
          type="submit"
          class="bg-dark-teal text-white text-lg px-4 py-2 h-fit rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
        >
          Tambah Kelas
        </button>
      </div>
    </form>
    <div v-if="errorMessage" class="text-red-700">{{ errorMessage }}</div>
  </div>
</template>

<script>
import { useDosenStore } from "@/stores/dosenStore";

export default {
  setup() {
    const DOSEN_STORE = useDosenStore();
    return { DOSEN_STORE };
  },
  data() {
    return {
      namaKelas: "",
      mataKuliah: "",
      errorMessage: "",
    };
  },
  methods: {
    async handleAddClass() {
      await this.DOSEN_STORE.addClass(this.namaKelas, this.mataKuliah);
      if (this.DOSEN_STORE.errorMessage) {
        this.errorMessage = this.DOSEN_STORE.errorMessage;
      } else {
        this.$router.push("/dosen/class");
      }
    },
  },
};
</script>
