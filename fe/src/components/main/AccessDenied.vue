<script setup lang="ts">
import router from "@/router";
</script>

<template>
  <div class="flex items-center justify-center h-screen">
    <div
      class="bg-white shadow-lg rounded-lg p-8 flex flex-col items-center text-center"
    >
      <h1 class="text-4xl font-bold text-dark-teal uppercase mb-4">
        Not Allowed
      </h1>
      <p class="text-xl text-gray-700 italic mb-6">
        Anda tidak memiliki hak untuk mengakses halaman ini.
      </p>
      <router-link
        to="/login"
        class="bg-dark-teal text-white text-lg px-6 py-3 rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
      >
        Kembali ke Halaman Awal
      </router-link>
    </div>
  </div>
</template>
