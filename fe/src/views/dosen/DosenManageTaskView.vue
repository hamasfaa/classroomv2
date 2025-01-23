<script setup>
import BannerTaskClass from "@/components/dosen/BannerTaskClass.vue";
import Navigation from "@/components/main/Navigation.vue";
import TaskClass from "@/components/dosen/TaskClass.vue";
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
    <BannerTaskClass />
    <TaskClass />
  </div>
</template>
