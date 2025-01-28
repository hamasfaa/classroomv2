<script setup>
import Navigation from "@/components/main/Navigation.vue";
import { useGeneralStore } from "@/stores/generalStore";
import BannerAttendanceClass from "@/components/dosen/BannerAttendanceClass.vue";
import AttendanceClass from "@/components/dosen/AttendanceClass.vue";

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
    <BannerAttendanceClass />
    <AttendanceClass />
  </div>
</template>
