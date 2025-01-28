<script setup lang="ts">
import BannerAttendance from "@/components/dosen/BannerAttendance.vue";
import Navigation from "@/components/main/Navigation.vue";
import List from "@/components/dosen/List.vue";
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
    <BannerAttendance />
    <List List="Attendance" />
  </div>
</template>
