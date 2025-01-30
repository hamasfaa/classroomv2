<script setup>
import Navigation from "@/components/main/Navigation.vue";
import { useGeneralStore } from "@/stores/generalStore";
import BannerAttendanceClass from "@/components/dosen/BannerAttendanceClass.vue";
import FormAttendance from "@/components/dosen/FormAttendance.vue";

const GENERAL_STORE = useGeneralStore();
</script>

<template>
  <div
    :class="[
      'flex flex-col h-screen p-6',
      GENERAL_STORE.isSidebarOpen ? 'mr-[82px]' : 'md:w-5/6 w-full',
    ]"
  >
    <BannerAttendanceClass :add="false" />
    <FormAttendance />
  </div>
</template>
