<script setup>
import BannerTaskClass from "@/components/dosen/BannerTaskClass.vue";
import FormTask from "@/components/dosen/FormTask.vue";
import Navigation from "@/components/main/Navigation.vue";
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
    <BannerTaskClass :add="false" />
    <FormTask />
  </div>
</template>
