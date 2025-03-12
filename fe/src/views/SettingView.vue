<script setup lang="ts">
import Contact from "@/components/main/Contact.vue";
import EditPhoto from "@/components/main/EditPhoto.vue";
import Navigation from "@/components/main/Navigation.vue";
import Other from "@/components/main/Other.vue";
import PersonalInformation from "@/components/main/PersonalInformation.vue";
import { useAuthStore } from "@/stores/authStore";
import { useGeneralStore } from "@/stores/generalStore";

const GENERAL_STORE = useGeneralStore();
</script>

<template>
  <div
    :class="[
      'flex flex-col h-screen p-6',
      GENERAL_STORE.isSidebarOpen ? 'mr-[82px]' : 'md:w-5/6 w-full',
    ]"
  >
    <div class="p-6 rounded-lg flex flex-col">
      <div class="mb-4">
        <h1 class="px-4 text-3xl font-bold text-dark-teal uppercase mb-2">
          Pengaturan Akun
        </h1>
        <EditPhoto />
        <PersonalInformation />
        <Contact />
        <Other />
      </div>
    </div>
  </div>
</template>
