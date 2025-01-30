<template>
  <div
    class="w-full md:w-2/3 bg-white flex flex-col justify-center items-center p-8"
  >
    <div
      class="border border-light-teal py-8 px-6 space-y-4 rounded-3xl shadow-lg md:py-16 md:px-12 md:space-y-8"
    >
      <h1
        class="text-2xl mb-4 text-light-teal font-extrabold md:text-4xl md:mb-6"
      >
        Siap belajar? Daftar sekarang!
      </h1>
      <form @submit.prevent="handleRegister">
        <div class="mb-4 w-full md:mb-4">
          <h2 class="text-lg mb-1 md:text-xl md:mb-2">Nama</h2>
          <div class="flex items-center border md:p-4 p-2 w-full rounded">
            <span class="material-symbols-outlined mr-2 text-light-teal">
              id_card
            </span>
            <input
              type="text"
              v-model="name"
              placeholder="Masukkan Nama Anda"
              class="flex-1 outline-none"
            />
          </div>
        </div>
        <div class="mb-4 w-full md:mb-4">
          <h2 class="text-lg mb-1 md:text-xl md:mb-2">Tanggal Lahir</h2>
          <div class="flex items-center border md:p-4 p-2 w-full rounded">
            <span class="material-symbols-outlined mr-2 text-light-teal">
              calendar_today
            </span>
            <input type="date" v-model="dob" class="flex-1 outline-none" />
          </div>
        </div>
        <div class="mb-4 w-full md:mb-4">
          <h2 class="text-lg mb-1 md:text-xl md:mb-2">Peran</h2>
          <div class="flex items-center space-x-4">
            <input
              type="radio"
              v-model="role"
              value="dosen"
              class="h-5 w-5 text-light-teal border-light-teal rounded-md focus:ring-2 focus:ring-light-teal"
            />
            <label for="Dosen" class="text-lg">Dosen</label>
            <input
              type="radio"
              v-model="role"
              value="mahasiswa"
              class="h-5 w-5 text-light-teal border-light-teal rounded-md focus:ring-2 focus:ring-light-teal"
            />
            <label for="mahasiswa" class="text-lg">Mahasiswa</label>
          </div>
        </div>
        <div class="mb-4 w-full md:mb-4">
          <h2 class="text-lg mb-1 md:text-xl md:mb-2">Email</h2>
          <div class="flex items-center border md:p-4 p-2 w-full rounded">
            <span class="material-symbols-outlined mr-2 text-light-teal">
              mail
            </span>
            <input
              type="email"
              v-model="email"
              placeholder="Masukkan Email Anda"
              class="flex-1 outline-none"
            />
          </div>
        </div>
        <div class="mb-4 w-full md:mb-4">
          <h2 class="text-lg mb-1 md:text-xl md:mb-2">Password</h2>
          <div class="flex items-center border md:p-4 p-2 w-full rounded">
            <span class="material-symbols-outlined mr-2 text-light-teal">
              lock
            </span>
            <input
              :type="isVisible ? 'text' : 'password'"
              v-model="password"
              placeholder="Masukkan Password Anda"
              class="flex-1 outline-none"
            />
            <span
              class="material-symbols-outlined mr-2 text-light-teal"
              @click="toggleVisible"
            >
              {{ isVisible ? "visibility" : "visibility_off" }}
            </span>
          </div>
        </div>
        <div class="mb-4 w-full md:mb-4">
          <h2 class="text-lg mb-1 md:text-xl md:mb-2">Konfirmasi Password</h2>
          <div class="flex items-center border md:p-4 p-2 w-full rounded">
            <span class="material-symbols-outlined mr-2 text-light-teal">
              lock
            </span>
            <input
              :type="isVisible ? 'text' : 'password'"
              v-model="confirmPassword"
              placeholder="Konfirmasi Password Anda"
              class="flex-1 outline-none"
            />
            <span
              class="material-symbols-outlined mr-2 text-light-teal"
              @click="toggleVisibleConfirm"
            >
              {{ isVisibleConfirm ? "visibility" : "visibility_off" }}
            </span>
          </div>
        </div>
        <button
          class="bg-light-teal text-white text-lg px-4 py-2 rounded border border-transparent hover:bg-white hover:border-light-teal hover:text-light-teal w-full"
          type="submit"
        >
          Daftar
        </button>
      </form>
      <div v-if="errorMessage" class="text-red-700">{{ errorMessage }}</div>
    </div>
  </div>
</template>

<script>
import { useAuthStore } from "@/stores/authStore";
import { useRouter } from "vue-router";

export default {
  setup() {
    const AUTH_STORE = useAuthStore();
    const ROUTER = useRouter();
    return { AUTH_STORE, ROUTER };
  },
  data() {
    return {
      name: "",
      dob: "",
      role: "",
      email: "",
      password: "",
      confirmPassword: "",
      errorMessage: "",
      isVisible: false,
      isVisibleConfirm: false,
    };
  },
  methods: {
    toggleVisible() {
      this.isVisible = !this.isVisible;
    },
    toggleVisibleConfirm() {
      this.isVisibleConfirm = !this.isVisibleConfirm;
    },
    async handleRegister() {
      if (this.password !== this.confirmPassword) {
        this.errorMessage = "Password tidak sama";
        return;
      }

      await this.AUTH_STORE.register(
        this.name,
        this.dob,
        this.role,
        this.email,
        this.password
      );
      if (this.AUTH_STORE.errorMessage) {
        this.errorMessage = this.AUTH_STORE.errorMessage;
      } else {
        this.ROUTER.push({ name: "login" });
      }
    },
  },
};
</script>
