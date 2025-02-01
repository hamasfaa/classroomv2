<template>
  <Teleport to="body">
    <div
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-100"
      v-if="GENERAL_STORE.isLogout"
    >
      <div
        class="bg-white shadow-lg rounded-lg p-8 flex flex-col items-center text-center"
      >
        <h1 class="text-4xl font-bold text-dark-teal uppercase mb-4">
          Konfirmasi
        </h1>
        <p class="text-xl text-gray-700 italic mb-6">
          Apakah Anda yakin ingin keluar?
        </p>
        <div class="flex space-x-4">
          <button
            class="bg-dark-teal text-white text-lg px-6 py-3 rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
            @click="AUTH_STORE.logout() && GENERAL_STORE.toggleLogout"
          >
            Ya, Keluar
          </button>
          <button
            class="bg-gray-300 text-dark-teal text-lg px-6 py-3 rounded-xl transition duration-300"
            @click="GENERAL_STORE.toggleLogout"
          >
            Batal
          </button>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script>
import { useAuthStore } from "@/stores/authStore";
import { useGeneralStore } from "@/stores/generalStore";
import { Teleport } from "vue";

export default {
  components: { Teleport },
  setup() {
    const GENERAL_STORE = useGeneralStore();
    const AUTH_STORE = useAuthStore();
    return { AUTH_STORE, GENERAL_STORE };
  },
};
</script>
