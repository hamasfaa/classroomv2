<script setup lang="ts">
import BannerTask from "@/components/dosen/BannerTask.vue";
import Navigation from "@/components/main/Navigation.vue";
import Task from "@/components/dosen/Task.vue";
import { useGeneralStore } from "@/stores/generalStore";

const GENERAL_STORE = useGeneralStore();
</script>

<template>
  <Navigation />
  <div
    :class="[
      'flex flex-col h-screen p-6',
      GENERAL_STORE.isSidebarOpen ? 'mr-[82px]' : 'md:w-5/6 w-full',
    ]"
  >
    <BannerTask />
    <Task />
  </div>
</template>
