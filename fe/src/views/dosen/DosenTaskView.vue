<script setup lang="ts">
import BannerTask from "@/components/dosen/BannerTask.vue";
import Navigation from "@/components/main/Navigation.vue";
import List from "@/components/dosen/List.vue";
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
    <BannerTask />
    <List List="Class" />
  </div>
</template>
