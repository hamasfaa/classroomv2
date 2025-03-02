<script setup lang="ts">
import BannerClass from "@/components/dosen/BannerClass.vue";
import Navigation from "@/components/main/Navigation.vue";
import Class from "@/components/dosen/Class.vue";
import { useGeneralStore } from "@/stores/generalStore";
import { useDosenStore } from "@/stores/dosenStore";
import { onMounted } from "vue";

const GENERAL_STORE = useGeneralStore();
const DOSEN_STORE = useDosenStore();

onMounted(() => {
  DOSEN_STORE.getAllClass();
});
</script>

<template>
  <div
    :class="[
      'flex flex-col h-screen p-6',
      GENERAL_STORE.isSidebarOpen ? 'mr-[82px]' : 'md:w-5/6 w-full',
    ]"
  >
    <BannerClass />
    <Class />
  </div>
</template>
